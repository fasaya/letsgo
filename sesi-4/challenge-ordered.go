package main

import (
	"fmt"
	"strconv"
	"sync"
)

type aksi2 interface {
	coba() string
	bisa() string
}

type Angka2 struct {
	angka int
}

func (j Angka2) coba() string {
	return "[coba] " + strconv.Itoa(j.angka)
}

func (j Angka2) bisa() string {
	return "[bisa] " + strconv.Itoa(j.angka)
}

func challengeOrdered() {
	fmt.Println("--- Ordered ---")

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 1; i <= 4; i++ {
		wg.Add(2)
		go fungsiSatuOrdered(&wg, i, &mu)
		go fungsiDuaOrdered(&wg, i, &mu)
	}

	wg.Wait()
}

func fungsiSatuOrdered(wg *sync.WaitGroup, i int, mu *sync.Mutex) {
	defer wg.Done()

	mu.Lock()

	var myAction aksi2 = Angka2{i}
	fmt.Println(myAction.coba())

	mu.Unlock()
}

func fungsiDuaOrdered(wg *sync.WaitGroup, i int, mu *sync.Mutex) {
	defer wg.Done()

	mu.Lock()

	var myAction aksi2 = Angka2{i}
	fmt.Println(myAction.bisa())

	mu.Unlock()
}
