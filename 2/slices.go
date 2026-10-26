package main

import "math/rand/v2"

const randomLimit = 100

// generateRandomSlice возвращает новый слайс из n случайных чисел в диапазоне [0, randomLimit).
func generateRandomSlice(n int) []int {
	numbers := make([]int, n)
	for i := range numbers {
		numbers[i] = rand.IntN(randomLimit)
	}
	return numbers
}

// sliceExample возвращает новый слайс, содержащий только чётные числа из исходного слайса.
func sliceExample(numbers []int) []int {
	evenNumbers := make([]int, 0, len(numbers))
	for _, number := range numbers {
		if number%2 == 0 {
			evenNumbers = append(evenNumbers, number)
		}
	}
	return evenNumbers
}

// addElements возвращает новый слайс: копию numbers с добавленным в конец
// значением value. Исходный слайс не изменяется.
func addElements(numbers []int, value int) []int {
	result := make([]int, len(numbers), len(numbers)+1)
	copy(result, numbers)
	return append(result, value)
}

// copySlice возвращает копию слайса, не разделяющую память с оригиналом.
func copySlice(numbers []int) []int {
	result := make([]int, len(numbers))
	copy(result, numbers)
	return result
}

// removeElement возвращает новый слайс без элемента по указанному индексу.
// Если индекс выходит за границы слайса, возвращается независимая копия исходного слайса.
func removeElement(numbers []int, index int) []int {
	if index < 0 || index >= len(numbers) {
		return copySlice(numbers)
	}

	result := make([]int, 0, len(numbers)-1)
	result = append(result, numbers[:index]...)
	result = append(result, numbers[index+1:]...)
	return result
}
