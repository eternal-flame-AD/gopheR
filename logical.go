package gopher

// #define R_NO_REMAP
// #include <R.h>
// #include <Rinternals.h>
import "C"

func BoolFromSEXP(sexp SEXP) bool {
	return int(C.LOGICAL_ELT(sexp, C.R_xlen_t(0))) != 0
}

func BoolToSEXP(b bool) SEXP {
	if b {
		return (SEXP)(C.Rf_ScalarLogical(C.int(1)))
	} else {
		return (SEXP)(C.Rf_ScalarLogical(C.int(0)))
	}
}
