package gopher

// #define R_NO_REMAP
// #include <R.h>
// #include <Rinternals.h>
import "C"

type DataFrame struct {
	sexp SEXP
}

func DataFrameFromSEXP(sexp SEXP) DataFrame {
	return DataFrame{sexp: sexp}
}

func (df DataFrame) NCol() int {
	return int(C.Rf_length(df.sexp))
}

func (df DataFrame) NRow() int {
	return int(C.Rf_length(C.VECTOR_ELT(df.sexp, C.R_xlen_t(0))))
}

func (df DataFrame) Col(i int) SEXP {
	return (SEXP)(C.VECTOR_ELT(df.sexp, C.R_xlen_t(i)))
}
