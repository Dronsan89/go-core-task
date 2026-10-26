package main

import "fmt"

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	hasIntersection, values := intersection(a, b)
	fmt.Println(hasIntersection, values)
}
