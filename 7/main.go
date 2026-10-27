package main

import "fmt"

func main() {
	first := make(chan int)
	second := make(chan int)

	go func() {
		defer close(first)
		first <- 1
		first <- 2
	}()

	go func() {
		defer close(second)
		second <- 3
		second <- 4
	}()

	for value := range merge(first, second) {
		fmt.Println(value)
	}
}
