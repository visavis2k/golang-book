package main

import "fmt"
import "io"
import "os"

// 이것은 주석이다

func main() {
	fmt.Printf("Hello, my name is %s\n", "daniel")
	fmt.Println("Hello, my name is", "daniel")
	s := fmt.Sprint("daniel", " is ", 22, " years old.\n")
	io.WriteString(os.Stdout, s)
}
