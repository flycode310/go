package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type Myfloat float64

func (f Myfloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	var a Abser
	f := Myfloat(-math.Sqrt2)
	a = f
	fmt.Println(a.Abs())
}
