package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":7070", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./templates/login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		r.ParseForm()
		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password: ", r.Form["password"])
	}
}
