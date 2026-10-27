package main

import "sync"

func merge(channels ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, channel := range channels {
		go func(ch <-chan int) {
			defer wg.Done()

			for value := range ch {
				out <- value
			}
		}(channel)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
