package main

import "fmt"

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	result := difference(slice1, slice2)

	fmt.Println("Первый слайс:", slice1)
	fmt.Println("Второй слайс:", slice2)
	fmt.Println("Элементы первого слайса, которых нет во втором: ", result)
}
