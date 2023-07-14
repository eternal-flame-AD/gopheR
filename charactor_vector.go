package gopher

// #define R_NO_REMAP
// #include <R.h>
// #include <Rinternals.h>
import "C"

type CharacterVector struct {
	sexp SEXP
}

func NewCharacterVector(length int) CharacterVector {
	return CharacterVector{(SEXP)(C.Rf_allocVector(C.STRSXP, C.R_xlen_t(length)))}
}

func CharacterVectorFromSEXP(sexp SEXP) CharacterVector {
	return CharacterVector{sexp}
}

func (v CharacterVector) SEXP() SEXP {
	return v.sexp
}

func (v CharacterVector) Length() int {
	return int(C.Rf_length(v.sexp))
}

func (v CharacterVector) Get(i int) string {
	return C.GoString(C.Rf_translateCharUTF8(C.STRING_ELT(v.sexp, C.R_xlen_t(i))))
}

func (v CharacterVector) Set(i int, s string) {
	C.SET_STRING_ELT(v.sexp, C.R_xlen_t(i), C.Rf_mkChar(C.CString(s)))
}
