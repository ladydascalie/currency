package accounting

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"

	"github.com/ladydascalie/currency/v2"
)

const (
	// QuadruplePrecision describes 128 bits of precision for IEEE 754 decimals
	// see: https://en.wikipedia.org/wiki/Quadruple-precision_floating-point_format
	QuadruplePrecision uint = 128

	// OctuplePrecision describes 256 bits of precision for for IEEE 754 decimals
	// see: https://en.wikipedia.org/wiki/Octuple-precision_floating-point_format
	OctuplePrecision uint = 256
)

var (
	// Minimum rate of tax must be 0.
	// As far as we know, there are as of writing, no markets with a negative sales tax.
	min = itof(0)

	// Baseline for creating a rate divisor.
	base = itof(1)
)

// Amount defines an amount in a given currency, in it's minor unit form.
type Amount struct {
	Currency   currency.Currency
	MinorValue int64
}

// MakeAmount returns an Amount from the provided currency and minor unit value.
func MakeAmount(c currency.Currency, minorValue int64) Amount {
	return Amount{
		Currency:   c,
		MinorValue: minorValue,
	}
}

// String implements a default stringer for an Amount
// Note that the string will be in "human readable" format, rather than
// using the minor currency unit, this conversion is done using the
// AmountToFloat64 function, also available within this package.
// ISO_4217 does not regulate spacing or prefixing vs. suffixing.
// Strings produced using this method always follow this pattern:
//
//	   ┏━━ always decimal dot separated.
//	   ┃
//	123.45 GBP
//	    ┃  ┗━━ ISO currency code.
//	    ┗━━ maximum 2 digits precision.
func (a Amount) String() string {
	f64 := AmountToFloat64(a)
	format := fmt.Sprintf("%%.%df %%s", a.Currency.MinorUnits())
	return fmt.Sprintf(format, f64, a.Currency.Code())
}

// ValidateManyFloatsArePrecise tests that the given float64 arguments
// have the desired precision, this is a convenience wrapper
// around ValidateFloatIsPrecise.
// NOTE: 2 digits past the dot is a business rule.
func ValidateManyFloatsArePrecise(args ...float64) error {
	for _, f := range args {
		if err := ValidateFloatIsPrecise(f); err != nil {
			return err
		}
	}
	return nil
}

// ValidateFloatIsPrecise ensures that a float64 value does not exceed
// a precision of 2 digits past the decimal period. This ensures we do not
// store incorrect currency data.
// NOTE: 2 digits past the decimal period is a business rule.
func ValidateFloatIsPrecise(f float64) error {
	// parse the float, with the smallest number of digits necessary
	parsed := strconv.FormatFloat(f, 'f', -1, 64)

	// split the parsed number on the decimal period
	parts := strings.Split(parsed, ".")

	// in case of exact number...
	if len(parts) == 1 {
		return nil
	}

	// check the float's precision
	if prec := len([]rune(parts[1])); prec > 2 {
		return ErrFloatPrecision{
			Value:     parsed,
			Precision: prec,
		}
	}
	return nil
}

// Float64ToAmount returns an amount from the provided currency and value.
func Float64ToAmount(c currency.Currency, value float64) Amount {
	minor := int64(math.Round(value * float64(c.Factor())))
	return Amount{
		Currency:   c,
		MinorValue: minor,
	}
}

// AmountToFloat64 returns the currency data as a floating point from it's
// minor currency unit format.
func AmountToFloat64(amount Amount) float64 {
	// fast path for currencies like JPY with a factor of 1.
	if amount.Currency.Factor() == 1 {
		return float64(amount.MinorValue)
	}
	var (
		v = itof(amount.MinorValue)
		f = ftof(amount.Currency.FactorF64())
	)
	f64, _ := newf().Quo(v, f).Float64()
	return f64
}

// Exchange - Apply currency exchange rates to an amount.
//
// rate - should always be given from the approved finance list.
// NOTE: A rate of zero will return the amount you put in, unchanged.
//
// Rounding to the nearest even is a defined business rule.
// Tills may round up to the nearest penny, but for reporting, the rule is
// always to use banker's rounding.
//
// If unclear, see: // http://wiki.c2.com/?BankersRounding.
func Exchange(amount Amount, c currency.Currency, rate float64) (Amount, error) {
	switch {
	case rate < 0:
		return Amount{}, ErrSubZeroRate
	case rate == 0:
		// Decomissioned currencies might trigger that case,
		// but that should really not be the general rule.
		return MakeAmount(c, 0), nil
	}

	from := ftof(AmountToFloat64(amount))
	bigRate := ftof(rate)

	// Here we divide the value, by it's minor currency
	// unit factor, then divide it once more by the
	// exchange rate.
	// -> v / e
	to, _ := from.Quo(from, bigRate).Float64()

	// http://wiki.c2.com/?BankersRounding
	// This part guarantees that we will not have more than 2 decimals after the dot.
	to = math.RoundToEven(to*100) / 100

	return Float64ToAmount(c, to), nil
}

// RatNetAmount applies a VAT rate to a big.Rat value. This method returns a big.Float
// so it's accuracy can be checked, and it's value applied with .Rat(some.field)
func RatNetAmount(gross, rate *big.Rat) (*big.Float, error) {
	// Here we go for octuple precision as we are dealing with rational numbers.
	bf := func(rat *big.Rat) *big.Float {
		return big.NewFloat(0).SetRat(rat).SetPrec(OctuplePrecision).SetMode(big.ToNearestEven)
	}

	v := bf(gross)
	r := bf(rate)

	// Guard against impossible (negative) tax rates.
	switch r.Cmp(min) {
	case -1:
		return min, ErrSubZeroRate
	case 0:
		return v, nil
	}
	// Turn the rate into a divisor by making it superior to 1.
	divisor := newf().Add(base, r)

	// Here we divide the gross by it's vat:
	// -> val / vat
	// where vat is a gross superior to 1.
	return newf().Quo(v, divisor), nil
}

// NetAmount derives the net amount before tax is applied using the given rate.
func NetAmount(gross int64, rate float64) (int64, error) {
	grossRat, _ := itof(gross).Rat(nil)
	rateRat, _ := ftof(rate).Rat(nil)

	bf, err := RatNetAmount(grossRat, rateRat)
	if err != nil {
		return 0, err
	}

	netStr := bf.Text('f', 0)
	net, _ := new(big.Int).SetString(netStr, 10)
	return net.Int64(), nil
}

// TaxAmount returns the difference between the gross and the net amounts.
func TaxAmount(gross, net int64) (int64, error) {
	// Guard against values that are not allowed in this context.
	if gross < 0 {
		return 0, ErrSubZeroGross
	}
	if net < 0 {
		return 0, ErrSubZeroNet
	}
	if net > gross {
		return 0, ErrNetOverGrossAmount
	}
	// list amount - net amount = tax amount
	return gross - net, nil
}
