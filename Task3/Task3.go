package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}
	ch := make(chan int)
	var wg sync.WaitGroup

	for _, num := range nums {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			ch <- num * num
		}(num)
	}
	//Запускаем отдельную горутину, которая закроет канал, чтобы мейн не останавливал работу
	go func() {
		wg.Wait()
		close(ch)
	}()
	sum := 0
	for num := range ch {
		sum += num
	}
	fmt.Println(sum)
}
