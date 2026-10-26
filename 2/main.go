package main

import "fmt"

func main() {
	originalSlice := generateRandomSlice(10)
	fmt.Println("Исходный слайс:", originalSlice)

	evenSlice := sliceExample(originalSlice)
	fmt.Println("\nТолько чётные числа:", evenSlice)

	extendedSlice := addElements(originalSlice, 42)
	fmt.Println("\nСлайс после добавления 42:", extendedSlice)
	fmt.Println("Исходный слайс не изменился:", originalSlice)

	copiedSlice := copySlice(originalSlice)
	if len(copiedSlice) > 0 {
		copiedSlice[0] = 101
	}
	fmt.Println("\nКопия после изменения первого элемента:", copiedSlice)
	fmt.Println("Изменение копии не затронуло оригинал:", originalSlice)

	if len(originalSlice) > 0 {
		index := len(originalSlice) / 2
		withoutElement := removeElement(originalSlice, index)
		fmt.Printf("\nСлайс без элемента с индексом %d: %v\n", index, withoutElement)
		fmt.Println("Исходный слайс не изменился:", originalSlice)
	}
}
