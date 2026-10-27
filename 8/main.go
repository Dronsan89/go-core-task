package main

import "fmt"

func main() {
	wg := NewWaitGroup()

	for i := 1; i <= 3; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			fmt.Printf("Горутина %d завершилась\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("Все горутины завершились")
}
