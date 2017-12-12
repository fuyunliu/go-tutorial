package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello, 世界")
	fmt.Println("The time is", time.Now())
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(math.Pi)
	fmt.Printf("Now you have %g problems.", math.Sqrt(7))
}
