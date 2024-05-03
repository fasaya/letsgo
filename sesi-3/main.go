package main

import "fmt"

type Employee struct {
	name     string
	age      int
	position string
}

type Person struct {
	name string
	age  int
}

func main() {
	fmt.Println("--- Sesi 3 ---")

	// 1. Struct
	fmt.Println("\n\n-- Struct --")
	var employee Employee
	employee.name = "aya"
	employee.age = 20
	employee.position = "developer"
	fmt.Println(employee.name)

	var employee2 = Employee{name: "ade", age: 19, position: "designer"}
	fmt.Println(employee2.name)

	var employee3 = Employee{}
	employee3.name = "aco"
	employee3.age = 17
	employee3.position = "student"
	fmt.Println(employee3.name)

	employee4 := Employee{}
	employee4.name = "ira"
	fmt.Println(employee4.name)

	// Pointer to struct
	// ...

	// Anonymous struct
	fmt.Println("\n\n-- Anonymous struct --")
	var member = struct {
		person   Person
		position string
	}{}
	member.person = Person{name: "Kim Ji Soo", age: 20}
	member.position = "Main Vocalist"
	fmt.Printf("%+v", member)

	var member2 = struct {
		person   Person
		position string
	}{
		person:   Person{name: "La Lisa", age: 20},
		position: "Main Dancer",
	}
	fmt.Printf("%+v", member2)

	// Slice to struct
	fmt.Println("\n\n-- Slice to struct --")
	var girlBand = []Person{
		{name: "Kim Ji Soo", age: 20},
		{name: "La Lisa", age: 22},
		{name: "Rose", age: 21},
		{name: "Jennie", age: 24},
	}

	for _, value := range girlBand {
		fmt.Printf("%+v", value)
	}

	// Slice To Anonymous Struct
	fmt.Println("\n\n-- Slice To Anonymous Struct --")
	var girlBand2 = []struct {
		person   Person
		position string
	}{
		{person: Person{name: "Kim Ji Soo", age: 20}, position: "Main Vocalist"},
		{person: Person{name: "La Lisa", age: 22}, position: "Main Dancer"},
		{person: Person{name: "Rose", age: 21}, position: "Main Rapper"},
		{person: Person{name: "Jennie", age: 24}, position: "Main Vocalist"},
	}

	for _, value := range girlBand2 {
		fmt.Printf("%+v", value)
	}

	// 2. Function
	fmt.Println("\n\n\n--- Function ---")
	greet("aya", "makassar", 20)
	fmt.Println(greet2("aco", "malang", 19))

}

func greet(name, city string, age int) {
	fmt.Println("hello my name is", name, "i live in", city, "my age is", age)
}
func greet2(name, city string, age int) string {
	return fmt.Sprintf("hello my name is %s i live in %s my age is %d", name, city, age)
}
