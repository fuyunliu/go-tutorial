package main

import "fmt"

func mysum1() {
	sum := 0
	for i:= 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func mysum2() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}

func mysum3() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func forever() {
	for {}
}
