# GopheR

Tool for generating R bindings to Go Functions.

Supported types:
- Logical (vector) -> bool
- Integer (vector) -> int64
- Real (vector) -> float64
- String (vector) -> string
- Data Frame

## Usage

First install the `mkrcall` codegen.

```sh
go install github.com/eternal-flame-AD/gopheR/cmd/mkrcall
```

Then write the functions you want in a Go main package.

```golang
// example.go
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

```

Then run `go generate` to generate the bindings.

```sh
go generate
go build -buildmode=c-shared -o example.so example.go
```

Then you can load and call the shared library in R.

```R
dyn.load("example.so")
.Call("Sum", c(1,2,3), TRUE)
```