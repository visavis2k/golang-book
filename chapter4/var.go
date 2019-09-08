package main

import "fmt"

var x = "Hello World"

const y = "Hello World"

func main() {
	fmt.Println(x)
	fmt.Println(y)
	v()
}

var (
	a = 5
	b = 10
	c = 15
)

func v() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
