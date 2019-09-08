package main

import "fmt"

func main() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2
	fmt.Printf("%9.0f\n", output)

	fmt.Print("화씨를 입력하세요: ")
	fmt.Scanf("%f", &input)

	output = ((input - 32) * 5 / 9)
	fmt.Printf("%9.3f\n", output)

	fmt.Print("피트를 입력하세요: ")
	fmt.Scanf("%f", &input)

	output = (input * 0.3048)
	fmt.Printf("%9.4f m\n", output)
}
