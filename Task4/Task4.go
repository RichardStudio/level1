package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	var n int
	fmt.Println("Number of workers: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}
	fmt.Println("Press any key to exit")
	ch := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		i := 0
		for {
			select {
			case <-ctx.Done():
				close(ch)
				return
			default:
				ch <- i
				i++
				time.Sleep(1 * time.Second)
			}
		}
	}()

	for i := 0; i < n; i++ {
		go worker(i, ch, ctx)
	}

	fmt.Scanln()
	defer cancel()
}

func worker(id int, jobs <-chan int, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}
			fmt.Printf("id: %d job: %d\n", id, job)
		}
	}
}
