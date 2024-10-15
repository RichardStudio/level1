package main

import (
	"fmt"
	"time"
)

func main() {
	var n int
	fmt.Println("Seconds: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}

	ch := make(chan int)

	go func() {
		for i := 0; i < n; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	for i := range ch {
		fmt.Println(i)
	}
}
