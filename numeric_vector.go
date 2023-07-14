package gopher

// #define R_NO_REMAP
// #include <R.h>
// #include <Rinternals.h>
import "C"

type NumericVector struct {
	sexp SEXP
}

func NewNumericVector(length int) NumericVector {
	return NumericVector{(SEXP)(C.Rf_allocVector(C.REALSXP, C.R_xlen_t(length)))}
}

func NumericVectorFromSEXP(sexp SEXP) NumericVector {
	return NumericVector{sexp}
}

func (v NumericVector) SEXP() SEXP {
	return v.sexp
}

func (v NumericVector) Length() int {
	return int(C.Rf_length(v.sexp))
}

func (v NumericVector) Get(i int) float64 {
	return float64(C.REAL_ELT(v.sexp, C.R_xlen_t(i)))
}

func (v NumericVector) Set(i int, f float64) {
	C.SET_REAL_ELT(v.sexp, C.R_xlen_t(i), C.double(f))
}
