package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Second)

	parse(ctx)
}

func parse(ctx context.Context) {
	for {
		select {
		case <-time.After(time.Second * 2):
			fmt.Println("parsing complited")
			return
		case <-ctx.Done():
			fmt.Println("dedline excede")
			return
		}
	}
}
