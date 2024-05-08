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
		fmt.Println("Please provide a name or ID as a command-line argument.")
		return
	}

	input := os.Args[1]

	var people = map[string]Person{
		"1": {"Budi", "Jakarta", "Programmer", "Belajar lebih banyak"},
		"2": {"Joko", "Bandung", "Pengusaha", "Banyak waktu kosong"},
		"3": {"Rudi", "Makassar", "Mahasiswa", "Mau cari kerja yang menggunakan bahasa go"},
		"4": {"Andi", "Malang", "Programmer", "Mau cari kerja yang menggunakan bahasa go"},
		"5": {"Andi", "Jakarta", "Mahasiswa", "Belajar lebih banyak"},
	}

	var foundPeople = map[string]Person{}

	// fmt.Println("People", people)

	for key, value := range people {
		if input == value.name || input == key {
			// fmt.Println("value", value.name)
			// append value to foundPeople
			foundPeople[key] = value
		}

	}

	if len(foundPeople) == 0 {
		fmt.Println("Data dengan nama/absen tsb tidak tersedia")
		return
	}

	for key, value := range foundPeople {
		fmt.Println("ID   		:", key)
		fmt.Println("Nama 		:", value.name)
		fmt.Println("Alamat		:", value.address)
		fmt.Println("Pekerjaan	:", value.job)
		fmt.Println("Alasan		:", value.reason)
		fmt.Println("--------------------------------------------------------------------------")
	}
}
