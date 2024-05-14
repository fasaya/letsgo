package main

import (
	"fmt"
	"sync"
)

func challengeUnordered() {
	fmt.Println("--- Unordered ---")

	var wg sync.WaitGroup

	firstInterface := []interface{}{1, 2, 3, 4}
	secondInterface := []interface{}{1, 2, 3, 4}

	for i := 0; i <= 3; i++ {
		wg.Add(2)

		idx := i

		go func() {
			defer wg.Done()
			fmt.Println("[coba]", firstInterface[idx])
		}()

		go func() {
			defer wg.Done()
			fmt.Println("[bisa]", secondInterface[idx])
		}()

	}

	wg.Wait()
}
