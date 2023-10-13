package main

import (
	"context"
	"log"
	"time"
)

func HardThing(ctx context.Context) (int, error) {
	log.Print("Hard thing start")
	i := 0
	for i = 0; i < 60; i++ {
		select {
		case <-ctx.Done():
			return 0, ctx.Err()
		default:
			log.Printf("Hard thing loop %d", i)
			time.Sleep(1 * time.Second)
		}
	}

	return i, nil
}
