package main

// cube преобразует число uint8 в float64 и возводит его в куб.
func cube(value uint8) float64 {
	number := float64(value)
	return number * number * number
}

// cubePipeline читает числа из input, преобразует их и отправляет кубы в output.
func cubePipeline(input <-chan uint8, output chan<- float64) {
	defer close(output)

	for value := range input {
		output <- cube(value)
	}
}
