package gopher

// #define R_NO_REMAP
// #include <R.h>
// #include <Rinternals.h>
import "C"

func Float64FromSEXP(sexp SEXP) float64 {
	return float64(C.REAL_ELT(sexp, C.R_xlen_t(0)))
}

func Float64ToSEXP(f float64) SEXP {
	return (SEXP)(C.Rf_ScalarReal(C.double(f)))
}
