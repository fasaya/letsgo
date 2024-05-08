package main

import (
	"fmt"
	"math"
)

type hitung interface {
	luas() float64
	keliling() float64
}

type persegi struct {
	sisi float64
}

type lingkaran struct {
	diameter, jariJari float64
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari, 2)
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

func (p persegi) keliling() float64 {
	return p.sisi
}

func catatan() {

	fmt.Println("---- Notes ----")

	var lingkaran1 hitung = lingkaran{14, 7}
	var persegi1 hitung = persegi{5}

	fmt.Println("Luas lingkaran		: ", lingkaran1.luas())
	fmt.Println("Keliling lingkaran	: ", lingkaran1.keliling())
	fmt.Println("Luas persegi		: ", persegi1.luas())
	fmt.Println("Keliling persegi	: ", persegi1.keliling())

	// fmt.Printf("Type of lingkaran: %T \n", lingkaran1)
	// fmt.Printf("Type of persegi: %T \n", persegi1)
}
