package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	bc := context.Background()
	t := 5000 * time.Millisecond
	ctx, cancel := context.WithTimeout(bc, t)
	defer cancel()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
