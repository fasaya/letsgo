package main

import "fmt"

func main() {
	fmt.Println("--- Sesi 2 ---")
	miniChallengeTwo()

	/**
	// 1. Array
	// has fixed length

	var animals = [3]string{"rabbit", "horse", "hippo"}
	animals[0] = "cat" // modify string
	animals[1] = "dog"
	animals[2] = "snake"
	fmt.Printf("%#v \n", animals)
	// or
	numbers := [3]int{1, 2, 3}
	fmt.Println(numbers)

	// loop
	fmt.Println("cara pertama")
	for key, animal := range animals {
		fmt.Printf("Index %d: value %s\n", key, animal)
	}

	fmt.Println("cara ketiga")
	for key := range animals { // better performance than other way of looping
		fmt.Printf("Index %d: value %s\n", key, animals[key])
	}

	// multidimensional array
	balances := [2][3]int{{1, 2, 3}, {5, 6, 7}}
	for _, valBal1 := range balances {
		for _, valBal2 := range valBal1 {
			fmt.Printf("%d ", valBal2)
		}
		fmt.Println("")
	}

	// 2. Slice
	// doesn't have fixed length

	// var fruits = []string{"apple", "orange", "banana"}
	// or
	var furniture = make([]string, 3)
	furniture[0] = "table"
	furniture[1] = "chair"
	furniture[2] = "drawer"

	// append
	furniture = append(furniture, "side table", "desk")
	fmt.Printf("%#v\n", furniture)

	// append function with ellipsis
	furniture2 := []string{"office chair"}
	furniture = append(furniture, furniture2...)

	fmt.Printf("%#v\n", furniture)

	// slice copy function
	animals_b1 := []string{"cat", "bird", "rabbit"}
	animals_b2 := []string{"horse", "duck", "butterfly"}

	copyFunc := copy(animals_b1, animals_b2)
	fmt.Println("Animals 1 =>", animals_b1)
	fmt.Println("Animals 2 =>", animals_b2)
	fmt.Println("Copied Animals =>", copyFunc)

	// Slice (slicing)
	fruits1 := []string{"apple", "mango", "orange", "banana", "grape"}
	var fruits2 = fruits1[1:4] // ambil index ke 1 sampe urutan ke-4
	// var fruits2 = fruits1[2:]  // ambil index ke 2 sampe urutan terakhir
	// var fruits2 = fruits1[:3]  // ambil index ke 0 sampe urutan ke-3
	// var fruits2 = fruits1[:]   // ambil index ke 0 sampe urutan terakhir (semua)
	fmt.Println("Sliced =>", fruits2)

	// slice (backing array)

	// 3. Map

	// 4. Pointer

	// ...

	**/

}
