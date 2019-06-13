package main

import (
	"fmt"
	mathy "github.com/overanalyze/awsome-mathy"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// Fxxk
	fmt.Println(mathy.Fib(10))
	fmt.Println(mathy.Fact(10))
	person := Person{"zhuguo", 24}
	fmt.Println(person)
}
