package gopher

// #define R_NO_REMAP
// #include <R.h>
// #include <Rinternals.h>
import "C"

type LogicalVector struct {
	sexp SEXP
}

func NewLogicalVector(n int) LogicalVector {
	return LogicalVector{(SEXP)(C.Rf_allocVector(C.LGLSXP, C.R_xlen_t(n)))}
}

func LogicalVectorFromSEXP(sexp SEXP) LogicalVector {
	return LogicalVector{sexp}
}

func (v LogicalVector) SEXP() SEXP {
	return v.sexp
}

func (v LogicalVector) Length() int {
	return int(C.Rf_length(v.sexp))
}

func (v LogicalVector) Get(i int) bool {
	return int(C.LOGICAL_ELT(v.sexp, C.R_xlen_t(i))) != 0
}

func (v LogicalVector) Set(i int, b bool) {
	if b {
		C.SET_LOGICAL_ELT(v.sexp, C.R_xlen_t(i), C.int(1))
	} else {
		C.SET_LOGICAL_ELT(v.sexp, C.R_xlen_t(i), C.int(0))
	}
}
