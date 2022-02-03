package main

import "fmt"

func fibonacci() func() int {
	f := 0
	s := 1
	return func() int {
		r := f
		f = s
		s += r
		return r
	}
}

func main() {
	fib := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}
}
