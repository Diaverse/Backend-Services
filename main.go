package main
 
import (
	"encoding/json"
	"database/sql"
	"fmt"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

//open the database connection
func openDBConnection(){
	db, err = sql.Open("mysql", "DB CONNECTION STRING")
	if err != nil {
		fmt.Println(err.Error())
	}
} //TODO: structure logic to also close connection when func is terminated

//struct to store row of user data from DB
type UserData struct {
	Username, Password, Email, FirstName, LastName, UserBio string
}

//struct to store hw token data from DB
type hardwareTokenData struct {
	Username, Hardwaretoken string
}

//struct to store script data from DB
type scriptData struct {
	Hardwaretoken, Script string
	ScriptID int
}

//retrieves entire row of user data based on username
func getUserDataByUsername(usersName string) UserData {
	openDBConnection()
	var userDataStruct UserData
	err = db.QueryRow("SELECT username, password, email, first_name, last_name, user_bio FROM users WHERE username = ?", usersName).Scan(&userDataStruct.Username, &userDataStruct.Password, &userDataStruct.Email, &userDataStruct.FirstName, &userDataStruct.LastName, &userDataStruct.UserBio)
	if err != nil {
		fmt.Println(err.Error())
	}
	return userDataStruct
}

//retrieves full row of hw token data based on username
func getHardwareTokenDataByUsername(usersName string) hardwareTokenData {
	openDBConnection()
	var hwTDataStruct hardwareTokenData
	err = db.QueryRow("SELECT username, hardwareToken FROM hardwareToken WHERE username = ?", usersName).Scan(&hwTDataStruct.Username, &hwTDataStruct.Hardwaretoken)
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
	//mux.Handle("/Dashboard", http.FileServer(http.Dir("./Frontend/Dashboard/examples")))
	mux.HandleFunc("/RegisterUser", handleRegister)
	mux.HandleFunc("/LoginUser", handleLogin)
	mux.HandleFunc("/RegisterHardware", handleHardwareRegister)
	mux.HandleFunc("/CheckForExistingHardwareToken",handleExistingHardwareToken)
	mux.HandleFunc("/AuthenticateHardware", handleAuthHardware)
	mux.HandleFunc("/GetScriptsByHardwareToken", handleGetScripts)
	mux.HandleFunc("/Logout", handleLogout)
	mux.HandleFunc("/GetUserInfoByUsername", handleGetUserInfo)
	mux.HandleFunc("/GetScriptResultsByHardwareToken", handleScriptResults)
	mux.HandleFunc("/UpdateUserInfo", handleUpdateUserInfo)
	mux.HandleFunc("/DeleteAccount", handleDeleteAccount)
	mux.HandleFunc("/DeleteScript", handleDeleteScript)
	mux.HandleFunc("/UploadScript", handleUploadScript)

	print("running\n")

	fmt.Println(http.ListenAndServe(":80", mux)) //change to 8080 for localhost
}

func handleLogin(w http.ResponseWriter, r *http.Request){
	(w).Header().Set("Access-Control-Allow-Origin", "*")
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
		fmt.Print("ACCEPTED\n") //TODO: add better logging
		w.WriteHeader(http.StatusAccepted) //let the user in
	} else {
		fmt.Print("FAILED ATTEMPTED LOGIN\n")
		w.WriteHeader(http.StatusForbidden) //forbid the user from entering
	}
}

//add the users entered information into the database
func handleRegister(w http.ResponseWriter, r *http.Request){
	//TODO: send verification email?
	r.ParseForm()
	uname := r.FormValue("username")
	passwd := r.FormValue("pass")
	email := r.FormValue("email")

	first_name := ""
	last_name := ""
	user_bio := ""
	insForm, err := db.Prepare("INSERT INTO users (username, password, email, first_name, last_name, user_bio) VALUES(?,?,?,?,?,?)")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Print(err.Error())
	} else {
		_, err = insForm.Exec(uname, passwd, email, first_name, last_name, user_bio)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Print(err.Error())
		}
		w.WriteHeader(http.StatusOK)
	}

}

func handleHardwareRegister(w http.ResponseWriter, r *http.Request){
	// perform some kind of verification and addition to the database

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
		fmt.Print("ACCEPTED\n")
		w.WriteHeader(http.StatusAccepted) // token is in DB
	} else {
		fmt.Print("ACCEPTED\n")
		w.WriteHeader(http.StatusForbidden) // token is not in DB.
	}
}

func handleAuthHardware(w http.ResponseWriter, r *http.Request){
	//I'm imaganing we will enter a registration number into the DB manually
	//verify that the value passed matches the value in the DB
	//if yes return accepted, if false return forbidden
}

//TODO: fix issue with retrieving data form DB
func handleGetScripts(w http.ResponseWriter, r *http.Request){
	//return all the scripts in the database associated with a given hardware token
	openDBConnection()
	r.ParseForm()
	hwid := r.FormValue("hardwareToken")
	selDB, err := db.Query("SELECT * FROM scripts WHERE hardwareToken =  ?", hwid)
	if err != nil {
	    fmt.Println(err.Error())
	}

	//aScript := scriptData{}
	scripts := []scriptData{}
	i := 0
	for selDB.Next(){
		var scriptID int
		var hwToken, script string
		err = selDB.Scan(&scriptID, &hwToken, &script)
		if err != nil {
			fmt.Println(err.Error())
		}
		scripts[i].Hardwaretoken = hwToken
		scripts[i].ScriptID = scriptID
		scripts[i].Script = script
		i++
	}

	j, e := json.Marshal(scripts)
	if e != nil {
		fmt.Println(e.Error())
	}
	jString := string(j)

	w.Write([]byte(jString))
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
