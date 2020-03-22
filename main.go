package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

var db *sql.DB
var err error

//open the database connection
func openDBConnection(){
	_, err := sql.Open("mysql",  "USERNAME-GOES-HERE:PASSWORD-GOES-HERE@tcp(ENDPOINT:PORT-NUMBER)/DATABSE-NAME")
	if err != nil {
		fmt.Println(err.Error())
	}
} //TODO: structure logic to also close connection when func is terminated

//struct to store row of user data from DB
type UserData struct {
	Username, Password, Email, Hardwaretoken, FirstName, LastName, UserBio string
}

//struct to store hw token data from DB
type hardwareTokenData struct {
	Username, Hardwaretoken, Authtoken string
}

//struct to store script data from DB
type scriptData struct {
	Hardwaretoken, ScriptID, Script string
}

//retrieves entire row of user data based on username
func getUserDataByUsername(usersName string) UserData {
	openDBConnection()
	var userDataStruct UserData
	err = db.QueryRow("SELECT username, password, email, hardwaretoken, first_name, last_name, user_bio FROM users WHERE username = ?", usersName).Scan(&userDataStruct.Username, &userDataStruct.Password, &userDataStruct.Email, &userDataStruct.Hardwaretoken, &userDataStruct.FirstName, &userDataStruct.LastName, &userDataStruct.UserBio)
	if err != nil {
		fmt.Println(err.Error())
	}
	return userDataStruct
}

//retrieves full row of hw token data based on username
func getHardwareTokenDataByUsername(usersName string) hardwareTokenData {
	openDBConnection()
	var hwTDataStruct hardwareTokenData
	err = db.QueryRow("SELECT username, hardwareToken, AuthToken FROM hardwareToken WHERE username = ?", usersName).Scan(&hwTDataStruct.Username, &hwTDataStruct.Hardwaretoken, &hwTDataStruct.Authtoken)
	if err != nil {
		fmt.Println(err.Error())
	}
	return hwTDataStruct
}

//retrieves row of script data based on script_id
func getScriptsByHWToken(scriptID string) scriptData {
	openDBConnection()
	var scriptStruct scriptData
	err = db.QueryRow("SELECT hardwareToken, script_id, script FROM scripts WHERE script_id = ?", scriptID).Scan(&scriptStruct.Hardwaretoken, &scriptStruct.ScriptID, &scriptStruct.Script)
	if err != nil {
		fmt.Println(err.Error())
	}
	return scriptStruct
}

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
	//get the username and password from the request/form
	uname := r.FormValue("username")
	passwd := r.FormValue("pass")

	var result bool

	//check that the password matches the username
	uDataStruct := getUserDataByUsername(uname)
	if uDataStruct.Password == passwd {
		result = true
	} else {
		result = false
	}

	//TODO: update login state in database, we need to setup a field for this

	if result {
		w.WriteHeader(http.StatusAccepted) //let the user in
	} else {
		w.WriteHeader(http.StatusForbidden) //forbid the user from entering
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
	r.ParseForm()
	uname := r.FormValue("username")
	hwToken := r.FormValue("hwtoken")

	var result bool
	hwTokenStruct := getHardwareTokenDataByUsername(uname)
	if hwTokenStruct.Hardwaretoken == hwToken {
		result = true
	} else {
		result = false
	}

	if result {
		w.WriteHeader(http.StatusAccepted) // token is in DB
	} else {
		w.WriteHeader(http.StatusForbidden) // token is not in DB.
	}
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


