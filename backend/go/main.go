package main

import (
	"fmt"
	// "gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/blas/blas64"
	"gonum.org/v1/gonum/mat"

	// Import netlib BLAS to register the C BLAS implementation
	"gonum.org/v1/netlib/blas/netlib"
)

func main() {
	blas64.Use(netlib.Implementation{})

	// Example: create two dense matrices and multiply them
	A := mat.NewDense(2, 3, []float64{
		1, 2, 3,
		4, 5, 6,
	})
	B := mat.NewDense(3, 2, []float64{
		7, 8,
		9, 10,
		11, 12,
	})

	var C mat.Dense
	C.Mul(A, B)

	fmt.Printf("Result:\n%v\n", mat.Formatted(&C, mat.Prefix(""), mat.Squeeze()))
}
