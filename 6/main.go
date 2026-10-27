package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	for value := range randomInts(ctx, 10) {
		fmt.Println(value)
	}
}
