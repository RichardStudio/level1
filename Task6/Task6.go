package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Context Example")
	contextExample()

	fmt.Println("\nChannel Example")
	channelExample()

	fmt.Println("\nTimeout Example")
	timeoutExample()

	fmt.Println("\nGlobal Variable Example")
	globalVariableExample()

	fmt.Println("\nDeadline Example")
	deadlineExample()
}

// 1. Остановка через контекст
func contextExample() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go workerWithContext(ctx, &wg)
	time.Sleep(time.Second)
	cancel()
	wg.Wait()
}

func workerWithContext(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker with context stopped")
			return
		default:
			fmt.Print(".")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// 2. Остановка через канал
func channelExample() {
	var wg sync.WaitGroup
	stopCh := make(chan struct{})
	wg.Add(1)
	go workerWithChannel(stopCh, &wg)
	time.Sleep(time.Second)
	close(stopCh)
	wg.Wait()
}

func workerWithChannel(stopCh <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-stopCh:
			fmt.Println("Worker with channel stopped")
			return
		default:
			fmt.Print(".")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// 3. Остановка через тайм-аут
func timeoutExample() {
	var wg sync.WaitGroup
	wg.Add(1)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	go workerWithTimeout(ctx, &wg)
	time.Sleep(3 * time.Second)
	wg.Wait()
}

func workerWithTimeout(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker with timeout stopped")
			return
		default:
			fmt.Print(".")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// 4. Остановка через глобальную переменную
var stop bool

func globalVariableExample() {
	var wg sync.WaitGroup
	stop = false
	wg.Add(1)
	go workerWithGlobalVariable(&wg)
	time.Sleep(time.Second)
	stop = true
	wg.Wait()
}

func workerWithGlobalVariable(wg *sync.WaitGroup) {
	defer wg.Done()
	for !stop {
		fmt.Print(".")
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("Worker with global variable stopped")
}

// 5. Остановка через дедлайн
func deadlineExample() {
	var wg sync.WaitGroup
	deadline := time.Now().Add(2 * time.Second)
	wg.Add(1)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	go workerWithDeadline(ctx, &wg)
	time.Sleep(3 * time.Second)
	wg.Wait()
}

func workerWithDeadline(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker with deadline stopped")
			return
		default:
			fmt.Print(".")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
