package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

var defaultTimeout = 1 * time.Second

func longExternalCall() int64 {
	rnd := rand.Int63n(5000)
	time.Sleep(time.Duration(rnd) * time.Millisecond)
	return rnd
}

// код выше нельзя менять

func wrappedCall(ctx context.Context) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	ch := make(chan int64, 1)

	go func() {
		ch <- longExternalCall()
		close(ch)
	}()

	select {
	case <-ctx.Done():
		return -1, ctx.Err()
	case res := <-ch:
		return res, nil
	}
}

func main() {
	fmt.Println("started")
	ctx := context.Background()

	start := time.Now()

	res, err := wrappedCall(ctx)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	fmt.Println("duration: ", time.Since(start))

	fmt.Println("result: ", res)
}
