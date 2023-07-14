//go:build cgo

package gopher

// #define R_NO_REMAP
// #include <R.h>
// #include <Rinternals.h>
import "C"

type SEXP C.SEXP

type Robj interface {
	SEXP() SEXP
}

func Protect(obj Robj) {
	C.Rf_protect(obj.SEXP())
}

func Unprotect(obj Robj) {
	C.Rf_unprotect_ptr(obj.SEXP())
}

func UnprotectN(n int) {
	C.Rf_unprotect(C.int(n))
}
