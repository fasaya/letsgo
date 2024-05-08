package main

import (
	"fmt"
	"sync"
)

type aksi interface {
	coba() int
	bisa() int
}

type Angka struct {
	angka int
}

func (j Angka) coba() int {
	return j.angka
}

func (j Angka) bisa() int {
	return j.angka
}

func challengeUnordered() {
	fmt.Println("--- Unordered ---")

	var wg sync.WaitGroup

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go fungsiSatu(&wg, i)
	}

	for j := 1; j <= 4; j++ {
		wg.Add(1)
		go fungsiDua(&wg, j)
	}

	wg.Wait()
}

func fungsiSatu(wg *sync.WaitGroup, i int) {
	defer wg.Done()

	var myAction aksi = Angka{i}
	fmt.Println("[coba]", myAction.coba())

}

func fungsiDua(wg *sync.WaitGroup, i int) {
	defer wg.Done()

	var myAction aksi = Angka{i}
	fmt.Println("[bisa]", myAction.bisa())

}
