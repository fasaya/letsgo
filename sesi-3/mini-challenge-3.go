package main

import (
	"fmt"
	"os"
)

type Person struct {
	name    string
	address string
	job     string
	reason  string
}

func minChallengeTiga() {

	// Check if there is at least one command-line argument
	if len(os.Args) < 2 {
		fmt.Println("Please provide a name as a command-line argument.")
		return
	}

	name := os.Args[1]

	var people = []Person{
		{"Budi", "Jakarta", "Programmer", "Belajar lebih banyak"},
		{"Joko", "Bandung", "Pengusaha", "Banyak waktu kosong"},
		{"Rudi", "Makassar", "Mahasiswa", "Mau cari kerja yang menggunakan bahasa go"},
		{"Andi", "Malang", "Programmer", "Mau cari kerja yang menggunakan bahasa go"},
		{"Andi", "Jakarta", "Mahasiswa", "Belajar lebih banyak"},
	}

	var foundPeople = []Person{}

	// fmt.Println("People", people)

	for _, value := range people {

		if name == value.name {
			// fmt.Println("value", value.name)
			// append value to foundPeople
			foundPeople = append(foundPeople, value)
		}

	}

	if len(foundPeople) == 0 {
		fmt.Println("Data dengan nama/absen tsb tidak tersedia")
		return
	}

	for _, value := range foundPeople {
		fmt.Println("Nama :", value.name)
		fmt.Println("Alamat: ", value.address)
		fmt.Println("Pekerjaan: ", value.job)
		fmt.Println("Alasan: ", value.reason)
		fmt.Println("--------------------------------------------------------------------------")
	}
}
