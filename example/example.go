package main

import (
	"log"

	gopher "github.com/eternal-flame-AD/gopheR"
)

//go:generate mkrcall -n example example.go

// rcall ReverseString
func reverseString(input gopher.CharacterVector) gopher.CharacterVector {
	length := input.Length()
	output := gopher.NewCharacterVector(length)
	for i := 0; i < length; i++ {
		s := input.Get(i)
		o := ""
		for _, c := range s {
			o = string(c) + o
		}
		output.Set(i, o)
	}
	return output
}

// rcall ReverseSingleString
func reverseSingleString(input string) string {
	o := ""
	for _, c := range input {
		o = string(c) + o
	}
	return o
}

// rcall Add
func add(a float64, b float64) float64 {
	return a + b
}

// rcall Sum
func sum(input gopher.NumericVector, quiet bool) float64 {
	length := input.Length()
	sum := 0.0
	for i := 0; i < length; i++ {
		if !quiet {
			log.Printf("Adding %f\n", input.Get(i))
		}
		sum += input.Get(i)
	}
	return sum
}

// rcall ColumnSum
func columnSum(input gopher.DataFrame, i int64) float64 {
	v := gopher.NumericVectorFromSEXP(input.Col(int(i)))
	s := 0.0
	for i := 0; i < v.Length(); i++ {
		s += v.Get(i)
	}
	return s
}

func main() {

}
