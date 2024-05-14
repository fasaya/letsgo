package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type User struct {
	ID      int
	Email   string
	Name    string
	Address string
	Job     string
}

var users = []User{
	{ID: 1, Email: "Pink_Goldner@hotmail.com", Name: "Walter Goyette MD", Address: "Jakarta", Job: "Direct Security Associate"},
	{ID: 2, Email: "Jordane67@gmail.com", Name: "Laura Hirthe", Address: "Bandung", Job: "Corporate Data Officer"},
	{ID: 3, Email: "Christopher_Roob32@gmail.com", Name: "Milton Wisoky", Address: "Surabaya", Job: "Direct Operations Consultant"},
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
	if r.Method == "GET" {
		tpl, err := template.ParseFiles("detail-page.html")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		email := r.URL.Query().Get("email")

		user, err := getUserByEmail(email)

		if err != nil {
			showNotFoundPage(w)
			return
		}

		tpl.Execute(w, user)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}

func getUserByEmail(email string) (*User, error) {
	for _, user := range users {
		if user.Email == email {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("user with email %s not found", email)
}

func showNotFoundPage(w http.ResponseWriter) {

	tpl, err := template.ParseFiles("notfound-page.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	tpl.Execute(w, struct{}{})
}
