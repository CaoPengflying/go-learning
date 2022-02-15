package main

import "math"

func main() {
	a,b,c := 2.0,1.0,0.0
	x,y := a/c,b/c
	n := math.NaN()
	m := math.Sqrt(-1.0)
	println(x)
	println(y)
	println(m)
	println(n)
	println(x == y, m==n)
}
