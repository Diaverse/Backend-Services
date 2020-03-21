package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.DefaultServeMux

	mux.Handle("/", http.FileServer(http.Dir("./Frontend/Login")))

	print("running")

	fmt.Println(http.ListenAndServe(":8080", mux))
}
