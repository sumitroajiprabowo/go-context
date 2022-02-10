package go_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second) // simulate slow processing
			}
		}
	}()
	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total Goroutines:", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)
	defer cancel()
	destination := CreateCounter(ctx)

	for n := range destination {
		fmt.Println("Counter:", n)
		if n == 10 {
			break
		}
	}
	time.Sleep(time.Second * 2)
	fmt.Println("Total Goroutines:", runtime.NumGoroutine())
}

func TestWithTimeOut(t *testing.T) {
	fmt.Println("Total Goroutines:", runtime.NumGoroutine())
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	destination := CreateCounter(ctx)
	fmt.Println("Total Goroutines:", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("Counter:", n)
	}
	time.Sleep(time.Second * 2)

	fmt.Println("Total Goroutines:", runtime.NumGoroutine())
}

func TestWithDeadline(t *testing.T) {
	fmt.Println("Total Goroutines:", runtime.NumGoroutine())
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))
	defer cancel()
	destination := CreateCounter(ctx)
	fmt.Println("Total Goroutines:", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("Counter:", n)
	}
	time.Sleep(time.Second * 2)

	fmt.Println("Total Goroutines:", runtime.NumGoroutine())
}
