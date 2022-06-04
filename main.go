package main

import (
	"assignment3/controllers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", controllers.ChangeValue)
	http.Handle("/static/", http.FileServer(http.Dir("assets")))

	address := "localhost:8080"
	log.Println("Server started at; ", address)
	err := http.ListenAndServe(address, nil)

	if err != nil {
		fmt.Println(err)
	}
}
