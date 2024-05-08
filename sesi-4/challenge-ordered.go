package main

import (
	"fmt"
	"sync"
)

type aksi2 interface {
	coba() int
	bisa() int
}

type Angka2 struct {
	angka int
}

func (j Angka2) coba() int {
	return j.angka
}

func (j Angka2) bisa() int {
	return j.angka
}

func challengeOrdered() {
	fmt.Println("--- Ordered ---")

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go fungsiSatuOrdered(&wg, i, &mu)
	}

	for j := 1; j <= 4; j++ {
		wg.Add(1)
		go fungsiDuaOrdered(&wg, j, &mu)
	}

	wg.Wait()
}

func fungsiSatuOrdered(wg *sync.WaitGroup, i int, mu *sync.Mutex) {
	defer wg.Done()
	var myAction aksi2 = Angka2{i}

	mu.Lock()
	fmt.Println("[coba]", myAction.coba())
	mu.Unlock()
}

func fungsiDuaOrdered(wg *sync.WaitGroup, i int, mu *sync.Mutex) {
	defer wg.Done()
	var myAction aksi2 = Angka2{i}

	mu.Lock()
	fmt.Println("[bisa]", myAction.bisa())
	mu.Unlock()
}
