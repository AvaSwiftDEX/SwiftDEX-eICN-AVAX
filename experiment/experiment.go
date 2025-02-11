package main

import (
	"context"
	"time"
)

func main() {
	ctx := context.Background()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	ctx.Done()
	println("done")
}
