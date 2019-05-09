package main

import (
	"fmt"
	"math"
)

func NewtonSqrt(num float64) float64 {
	x := num / 2.0
	var y float64 = 0
	count := 1
	for math.Abs(x-y) > 0.00000001 {
		fmt.Println(count, x)
		count += 1
		y = x
		x = (1.0/2.0)*x + (num*1.0)/(x*2.0)
	}
	return x
}

func BinarySqrt(num float64) float64 {
	y := num / 2.0
	low := 0.0
	up := num
	count := 1
	for math.Abs(y*y - num) > 0.00000001 {
		fmt.Println(count, y)
		count += 1
		if y*y > num {
			up = y
			y = low + (y-low)/2
		} else {
			low = y
			y = up - (up-y)/2
		}
	}
	return y
}

func main() {
	fmt.Println("math sqrt", math.Sqrt(5))
	fmt.Println("Newton sqrt", NewtonSqrt(5))
	fmt.Println("Binary sqrt", BinarySqrt(5))
}
