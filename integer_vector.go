package gopher

// #define R_NO_REMAP
// #include <R.h>
// #include <Rinternals.h>
import "C"

type IntegerVector struct {
	sexp SEXP
}

func NewIntegerVector(length int) IntegerVector {
	return IntegerVector{(SEXP)(C.Rf_allocVector(C.INTSXP, C.R_xlen_t(length)))}
}

func IntegerVectorFromSEXP(sexp SEXP) IntegerVector {
	return IntegerVector{sexp}
}

func (v IntegerVector) SEXP() SEXP {
	return v.sexp
}

func (v IntegerVector) Length() int {
	return int(C.Rf_length(v.sexp))
}

func (v IntegerVector) Get(i int) int {
	return int(C.INTEGER_ELT(v.sexp, C.R_xlen_t(i)))
}

func (v IntegerVector) Set(i int, n int) {
	C.SET_INTEGER_ELT(v.sexp, C.R_xlen_t(i), C.int(n))
}
