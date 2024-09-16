package accounting

import "math/big"

func newf() *big.Float {
	f := new(big.Float)
	f.SetPrec(QuadruplePrecision)
	f.SetMode(big.ToNearestEven)
	return f
}

func itof(v int64) *big.Float {
	f := newf()
	f.SetInt64(v)
	return f
}

func ftof(v float64) *big.Float {
	f := newf()
	f.SetFloat64(v)
	return f
}
