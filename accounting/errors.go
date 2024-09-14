package accounting

import (
	"errors"
	"fmt"
)

var (
	// ErrSubZeroRate happens when a rate is lower than zero.
	ErrSubZeroRate = errors.New("rate must not be less than zero")
	// ErrSubZeroGross happens when the gross amount is less than zero.
	ErrSubZeroGross = errors.New("gross amount must not be less than zero")
	// ErrSubZeroNet happens when the net amount is less than zero.
	ErrSubZeroNet = errors.New("net amount must not be less than zero")
	// ErrNetOverGrossAmount happens when the net amount is higher than the gross amount.
	ErrNetOverGrossAmount = errors.New("net amount must be equal to or lower than the gross amount")
)

// ErrFloatPrecision happens when a floating point number is not following business precision rules.
type ErrFloatPrecision struct {
	Value     string
	Precision int
}

func (e ErrFloatPrecision) Error() string {
	return fmt.Sprintf("incorrect value %s with precision %d", e.Value, e.Precision)
}
