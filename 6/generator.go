package main

import (
	"context"
	"math/rand/v2"
)

const randomLimit = 100

// randomInts возвращает канал со случайными числами.
func randomInts(ctx context.Context, count int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for i := 0; i < count; i++ {
			value := rand.IntN(randomLimit)

			select {
			case out <- value:
			case <-ctx.Done():
				return
			}
		}
	}()

	return out
}
