package main

import (
	"fmt"
	"sync"
)

func challengeOrdered() {
	fmt.Println("--- Odered ---")

	var wg sync.WaitGroup
	var mu sync.Mutex

	firstInterface := []interface{}{1, 2, 3, 4}
	secondInterface := []interface{}{1, 2, 3, 4}

	for i := 0; i <= 3; i++ {
		wg.Add(1)
		idx := i

		go func() {
			defer wg.Done()
			mu.Lock()
			fmt.Println("[coba]", firstInterface[idx])
			mu.Unlock()
		}()

	}

	for j := 0; j <= 3; j++ {
		wg.Add(1)
		jdx := j

		go func() {
			defer wg.Done()
			mu.Lock()
			fmt.Println("[bisa]", secondInterface[jdx])
			mu.Unlock()
		}()
	}

	wg.Wait()
}
