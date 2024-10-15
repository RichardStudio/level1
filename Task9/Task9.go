package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch1 := make(chan int)
	ch2 := make(chan int)
	var wg sync.WaitGroup

	// Запись в первый канал
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, i := range arr {
			ch1 <- i
		}
		close(ch1)
	}()

	// Чтение из первого канала и запись во второй
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range ch1 {
			ch2 <- i * 2
		}
		close(ch2)
	}()

	// Чтение из второго канала и вывод
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range ch2 {
			fmt.Println(i)
		}
	}()

	wg.Wait()
}
