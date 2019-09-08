package main

import "fmt"

func main() {
	// var x [5]int
	// x[4] = 100
	// fmt.Println(x)
	//
	// var y [5]float64
	// y[0] = 98
	// y[1] = 93
	// y[2] = 77
	// y[3] = 82
	// y[4] = 83

	// var total float64 = 0
	// for i := 0; i < len(y); i++ {
	// 	total += y[i]
	// }
	// fmt.Println(total / float64(len(y)))

	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))

	y := [5]float64{
		98,
		93,
		77,
		82,
		83,
	}

	fmt.Println(y)

	var z []float64
	z = x[0:1]
	// z[0] = 1

	fmt.Println(z)

	arr := []float64{1, 2, 3, 4, 5}
	z = arr[:4]

	fmt.Println(z)

	slice1 := []float64{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println(slice3, slice4)
}
