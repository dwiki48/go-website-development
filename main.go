package main

import (
	"log"
	"net/http"
	"website-development/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/product", handler.ProductHandler)

	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("starting on port 666")
	err := http.ListenAndServe("127.0.0.1:666", mux)

	log.Fatal(err)
}
