//go:build cgo
// Code generated by mkrcall; DO NOT EDIT.

package main

// #cgo CFLAGS: -I/usr/include/R/
// #cgo LDFLAGS: -L/usr/lib64/R/lib -lR -lpcre2-8 -llzma -lbz2 -lz -ltirpc -lrt -ldl -lm -licuuc -licui18n
// #define R_NO_REMAP
// #include <R.h>
// #include <Rinternals.h>
// SEXP ReverseString(SEXP);
// SEXP ReverseSingleString(SEXP);
// SEXP Add(SEXP, SEXP);
// SEXP Sum(SEXP, SEXP);
// SEXP ColumnSum(SEXP, SEXP);
import "C"

import (
	"unsafe"
	gopher "github.com/eternal-flame-AD/gopheR"
)

//export ReverseString
func ReverseString(input0 C.SEXP) C.SEXP {
	input0_sexp := (*gopher.SEXP)(unsafe.Pointer(&input0))
	input0_val := gopher.CharacterVectorFromSEXP(*input0_sexp)
	output := reverseString(input0_val)
	output_sexp := output.SEXP()
	return *(*C.SEXP)(unsafe.Pointer(&output_sexp))
}

//export ReverseSingleString
func ReverseSingleString(input0 C.SEXP) C.SEXP {
	input0_sexp := (*gopher.SEXP)(unsafe.Pointer(&input0))
	input0_val := gopher.StringFromSEXP(*input0_sexp)
	output := reverseSingleString(input0_val)
	output_sexp := gopher.StringToSEXP(output)
	return *(*C.SEXP)(unsafe.Pointer(&output_sexp))
}

//export Add
func Add(input0 C.SEXP, input1 C.SEXP) C.SEXP {
	input0_sexp := (*gopher.SEXP)(unsafe.Pointer(&input0))
	input0_val := gopher.Float64FromSEXP(*input0_sexp)
	input1_sexp := (*gopher.SEXP)(unsafe.Pointer(&input1))
	input1_val := gopher.Float64FromSEXP(*input1_sexp)
	output := add(input0_val, input1_val)
	output_sexp := gopher.Float64ToSEXP(output)
	return *(*C.SEXP)(unsafe.Pointer(&output_sexp))
}

//export Sum
func Sum(input0 C.SEXP, input1 C.SEXP) C.SEXP {
	input0_sexp := (*gopher.SEXP)(unsafe.Pointer(&input0))
	input0_val := gopher.NumericVectorFromSEXP(*input0_sexp)
	input1_sexp := (*gopher.SEXP)(unsafe.Pointer(&input1))
	input1_val := gopher.BoolFromSEXP(*input1_sexp)
	output := sum(input0_val, input1_val)
	output_sexp := gopher.Float64ToSEXP(output)
	return *(*C.SEXP)(unsafe.Pointer(&output_sexp))
}

//export ColumnSum
func ColumnSum(input0 C.SEXP, input1 C.SEXP) C.SEXP {
	input0_sexp := (*gopher.SEXP)(unsafe.Pointer(&input0))
	input0_val := gopher.DataFrameFromSEXP(*input0_sexp)
	input1_sexp := (*gopher.SEXP)(unsafe.Pointer(&input1))
	input1_val := gopher.Int64FromSEXP(*input1_sexp)
	output := columnSum(input0_val, input1_val)
	output_sexp := gopher.Float64ToSEXP(output)
	return *(*C.SEXP)(unsafe.Pointer(&output_sexp))
}

func R_init_example(DllInfo *C.DllInfo) {
	callMethodDef := []C.R_CallMethodDef{
		{C.CString("ReverseString"), (C.DL_FUNC)(C.ReverseString), 1, [4]byte{0,0,0,0}},
		{C.CString("ReverseSingleString"), (C.DL_FUNC)(C.ReverseSingleString), 1, [4]byte{0,0,0,0}},
		{C.CString("Add"), (C.DL_FUNC)(C.Add), 2, [4]byte{0,0,0,0}},
		{C.CString("Sum"), (C.DL_FUNC)(C.Sum), 2, [4]byte{0,0,0,0}},
		{C.CString("ColumnSum"), (C.DL_FUNC)(C.ColumnSum), 2, [4]byte{0,0,0,0}},
		{nil, nil, 0, [4]byte{0,0,0,0}},
	}
	C.R_registerRoutines(DllInfo, nil, &callMethodDef[0], nil, nil)
	C.R_useDynamicSymbols(DllInfo, 0)
}

