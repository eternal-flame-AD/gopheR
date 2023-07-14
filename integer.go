package gopher

// #define R_NO_REMAP
// #include <R.h>
// #include <Rinternals.h>
import "C"

func Int64FromSEXP(sexp SEXP) int64 {
	return int64(C.INTEGER_ELT(sexp, C.R_xlen_t(0)))
}

func Int64ToSEXP(i int64) SEXP {
	return (SEXP)(C.Rf_ScalarInteger(C.int(i)))
}
