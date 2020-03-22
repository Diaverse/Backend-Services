package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

//open the database connection
func openDBConnection(){
	db, err := sql.Open("mysql",  "USERNAME-GOES-HERE:PASSWORD-GOES-HERE@tcp(ENDPOINT:PORT-NUMBER)/DATABSE-NAME")
	if err != nil {
		panic(err)
	}
} //TODO: structure logic to also close connection when func is terminated


func main() {

	mux := http.DefaultServeMux
	mux.Handle("/", http.FileServer(http.Dir("./Frontend/Login")))
	mux.HandleFunc("/RegisterUser", handleRegister)
	mux.HandleFunc("/LoginUser", handleLogin)
	mux.HandleFunc("/RegisterHardware", handleHardwareRegister)
	mux.HandleFunc("/CheckForExistingHardwareToken", handleExistingHardwareToken)
	mux.HandleFunc("/AuthenticateHardware", handleAuthHardware)
	mux.HandleFunc("/GetScriptsByHardwareToken", handleGetScripts)
	mux.HandleFunc("/Logout", handleLogout)
	mux.HandleFunc("/GetUserInfoByUsername", handleGetUserInfo)
	mux.HandleFunc("/GetScriptResultsByHardwareToken", handleScriptResults)
	mux.HandleFunc("/UpdateUserInfo", handleUpdateUserInfo)
	mux.HandleFunc("/DeleteAccount", handleDeleteAccount)
	mux.HandleFunc("/DeleteScript", handleDeleteScript)
	mux.HandleFunc("/UploadScript", handleUploadScript)

	print("running")

	fmt.Println(http.ListenAndServe(":8080", mux))
}

func handleLogin(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	//uname := r.FormValue("username")
	//pass := r.FormValue("pass")

	//verify in db
	result := true

	//update login state in database

	if result {
		w.WriteHeader(http.StatusAccepted)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

func handleRegister(w http.ResponseWriter, r *http.Request){
	//add the users entered information into the database
	//send verification email?
}

func handleHardwareRegister(w http.ResponseWriter, r *http.Request){
	// ???
}

func handleExistingHardwareToken(w http.ResponseWriter, r *http.Request){
	//simply check if token is already in DB, if so return true, else return false
}

func handleAuthHardware(w http.ResponseWriter, r *http.Request){
	//I'm imaganing we will enter a registration number into the DB manually
	//verify that the value passed matches the value in the DB
	//if yes return accepted, if false return forbidden
}

func handleGetScripts(w http.ResponseWriter, r *http.Request){
	//return all the scripts in the database associated with a given hardware token
}

func handleLogout(w http.ResponseWriter, r *http.Request){
	// save a logged in or logged out state in the database
	// if this is called, set the state to logged out
}

func handleGetUserInfo(w http.ResponseWriter, r *http.Request){
	// based on the user's username - return all info about them from DB
}

func handleScriptResults(w http.ResponseWriter, r *http.Request){
	// based on the hardware token - return the results of the test run.
}

func handleUpdateUserInfo(w http.ResponseWriter, r *http.Request){
	// the writer should contain all the info pertaining ot the update
	// simply update all fields in the databased based on this
}

func handleDeleteAccount(w http.ResponseWriter, r *http.Request){
	// based on the users username delete all assocaited information from the database.
}

func handleDeleteScript(w http.ResponseWriter, r *http.Request){
	// based on a unique value, delete the script from the database
}

func handleUploadScript(w http.ResponseWriter, r *http.Request){
	// the script will be in the text of the response writer
	// add the script as a new entry in the database.
}


