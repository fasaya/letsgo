package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
)

type User struct {
	ID      int
	Email   string
	Address string
	Job     string
}

var users = []User{
	{ID: 1, Email: "Walter Goyette MD", Address: "Jakarta", Job: "Direct Security Associate"},
	{ID: 2, Email: "Laura Hirthe", Address: "Bandung", Job: "Corporate Data Officer"},
	{ID: 3, Email: "Milton Wisoky", Address: "Surabaya", Job: "Direct Operations Consultant"},
}

var PORT = ":9090"

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/user", getUser)

	fmt.Println("Application is listening on port", PORT)
	http.ListenAndServe(PORT, nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tpl, err := template.ParseFiles("main-page.html")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tpl.Execute(w, users)
		return
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(users)
		return
	}

	// using html/template package
	// if r.Method == "GET" {
	// 	tpl, err := template.ParseFiles("template.html")
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}

	// 	tpl.Execute(w, users)
	// 	return
	// }

	http.Error(w, "Invalid method", http.StatusBadRequest)
}
