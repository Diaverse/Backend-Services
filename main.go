package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.DefaultServeMux
	mux.Handle("/", http.FileServer(http.Dir("./Frontend/Login")))
	mux.HandleFunc("/RegisterUser", handleLogin)
	mux.HandleFunc("/LoginUser", )
	mux.HandleFunc("/RegisterHardware", )
	mux.HandleFunc("/CheckForExistingHardwareToken", )
	mux.HandleFunc("/AuthenticateHardware", )
	mux.HandleFunc("/GetScritpsByHardwareToken", )
	mux.HandleFunc("/Logout", )
	mux.HandleFunc("/GetUserInfoByUsername", )
	mux.HandleFunc("/GetScriptResultsByHardwareToken", )
	mux.HandleFunc("/UpdateUserInfo", )
	mux.HandleFunc("/DeleteAccount", )
	mux.HandleFunc("/DeleteScript", )
	mux.HandleFunc("/UploadScript", )

	print("running")

	fmt.Println(http.ListenAndServe(":8080", mux))
}

func handleLogin(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	uname := r.FormValue("username")
	pass := r.FormValue("pass")

	//verify in db
	result := true

	if result {
		w.WriteHeader(http.StatusAccepted)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}
