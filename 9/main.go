package main

import "fmt"

func main() {
	input := make(chan uint8)
	output := make(chan float64)

	go cubePipeline(input, output)

	values := []uint8{1, 2, 3, 4, 5, 10, 255}

	go func() {
		defer close(input)

		for _, value := range values {
			input <- value
		}
	}()

	fmt.Println("Исходные значения:", values)

	fmt.Println("Кубы чисел:")
	for value := range output {
		fmt.Printf("%.0f\n", value)
	}
}
