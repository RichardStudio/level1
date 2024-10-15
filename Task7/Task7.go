package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	dataMap := make(map[string]int)

	// Количество горутин
	numRoutines := 5

	for i := 0; i < numRoutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			value := i
			mu.Lock()
			dataMap[key] = value
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	for key, value := range dataMap {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
	}
}
