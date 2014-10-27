package main

import (
	"fmt"
	"github.com/mrbech/goremotefunc"
)

func main() {
	r := remotefunc.New()
	r.AddFunc("Test", Test)
	r.AddFunc("Printer", Printer)
	r.AddFunc("Fibonacci", Fibonacci)
	r.Start()

}

func Test() {
	fmt.Println("Test called")
}

func Printer(s string) {
	fmt.Println("Printer called")
	fmt.Println(s)
}

func Fibonacci(n int) int {
	if n < 3 {
		return 1
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}
