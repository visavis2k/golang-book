package main

import "fmt"

func average(xs []float64) (total float64) {
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))

	fmt.Println(f())
}

func f() (int, int) {
	return 5, 6
}
