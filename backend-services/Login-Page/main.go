package main

import (
	"fmt"
	"net/http"
)


func main() {

	http.Handle("/", http.FileServer(http.Dir("./Login")))

	print("running")


	fmt.Println(http.ListenAndServe(":80", nil));
}
