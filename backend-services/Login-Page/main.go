package main

import (
	"fmt"
	"net/http"
)


func main() {

	http.Handle("/", http.FileServer(http.Dir("/Backend-Service/backend-services/Frontend/Login")))

	print("running")


	fmt.Println(http.ListenAndServe(":8080", nil));
}
