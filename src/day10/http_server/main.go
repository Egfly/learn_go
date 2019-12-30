package main

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle hello")
	_, _ = fmt.Fprintf(w, "hello ")
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle login")
	_, _ = fmt.Fprintf(w, "login ")
}

func main() {
	http.HandleFunc("/a", Hello)
	http.HandleFunc("/user/login", Login)
	err := http.ListenAndServe("0.0.0.0:8880", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}
