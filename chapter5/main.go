package main

import "fmt"

func main() {
	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
	fmt.Println(4)
	fmt.Println(5)
	fmt.Println(6)
	fmt.Println(7)
	fmt.Println(8)
	fmt.Println(9)
	fmt.Println(10)

	fmt.Println(`1
2
3
4
5
6
7
8
9
10`)

	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "짝수")
		} else {
			fmt.Println(i, "홀수")
		}

	}

	i := 4
	switch i {
	case 0:
		fmt.Println("영")
	case 1:
		fmt.Println("일")
	case 2:
		fmt.Println("이")
	case 3:
		fmt.Println("삼")
	case 4:
		fmt.Println("사")
	case 5:
		fmt.Println("오")
	default:
		fmt.Println("알 수 없는 숫자")
	}

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println(i, "Buzz")
		} else if i%3 == 0 {
			fmt.Println(i, "Fizz")
		}
	}

}
