package gopher

// #define R_NO_REMAP
// #include <R.h>
// #include <Rinternals.h>
import "C"

func StringFromSEXP(sexp SEXP) string {
	str := C.Rf_translateCharUTF8(C.STRING_ELT(sexp, C.R_xlen_t(0)))
	return C.GoString(str)
}

func StringToSEXP(s string) SEXP {
	return (SEXP)(C.Rf_mkString(C.CString(s)))
}
