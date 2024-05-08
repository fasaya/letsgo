package main

import (
	"fmt"
	"strconv"
	"sync"
)

type aksi interface {
	coba() string
	bisa() string
}

type Angka struct {
	angka int
}

func (j Angka) coba() string {
	return "[coba] " + strconv.Itoa(j.angka)
}

func (j Angka) bisa() string {
	return "[bisa] " + strconv.Itoa(j.angka)
}

func challengeUnordered() {
	fmt.Println("--- Unordered ---")

	var wg sync.WaitGroup

	for i := 1; i <= 4; i++ {
		wg.Add(2)
		go fungsiSatu(&wg, i)
		go fungsiDua(&wg, i)
	}

	wg.Wait()
}

func fungsiSatu(wg *sync.WaitGroup, i int) {
	defer wg.Done()

	var myAction aksi = Angka{i}
	fmt.Println(myAction.coba())

}

func fungsiDua(wg *sync.WaitGroup, i int) {
	defer wg.Done()

	var myAction aksi = Angka{i}
	fmt.Println(myAction.bisa())

}
