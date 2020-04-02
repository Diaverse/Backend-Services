package main

import (
  	"github.com/gorilla/handlers"
	"encoding/json"
	"database/sql"
	"fmt"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/smtp"
	"math/rand"
	"time"
	"strconv"
	"io/ioutil"
	"strings"
)

var db *sql.DB
var err error
var dbCred = ""
var emailPass = ""

//open the database connection
func openDBConnection(){
	db, err = sql.Open("mysql", dbCred)
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


	dat, _ := ioutil.ReadFile("./creds.txt")
	creds := strings.Split(string(dat), "\n")
	dbCred = creds[0]
	emailPass = creds[1]

	mux := http.DefaultServeMux
	mux.Handle("/", http.FileServer(http.Dir("/home/ec2-user/Dashboard/examples")))
	//mux.Handle("/Dashboard", http.FileServer(http.Dir("./Frontend/Dashboard/examples")))
	mux.HandleFunc("/RegisterUser", handleRegister)
	mux.HandleFunc("/LoginUser", handleLogin)
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

  //uses old school gorilla package to handle mux
  headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}) //only allowed headers
  methods := handlers.AllowedMethods([]string{"GET", "POST"}) //only allowed requests
  origins := handlers.AllowedOrigins([]string{"*"}) //any possible domain origin

	print("running\n")

  //Cross Origin Resource Sharing passed to listen and server
  fmt.Println(http.ListenAndServe(":80", handlers.CORS(headers, methods, origins) (mux))) //change to 8080 for localhost
}

func handleLogin(w http.ResponseWriter, r *http.Request){
	//(w).Header().Set("Access-Control-Allow-Origin", "*") //may not need this
	r.ParseForm()
	//get the username and password from the request/form
	uname := r.FormValue("username")
	passwd := r.FormValue("pass")

	result := false
	fmt.Println(uname)
	fmt.Println(passwd)
	//check that the password matches the username
	uDataStruct := getUserDataByUsername(uname)
	if uDataStruct.Password == passwd && uDataStruct.Password != "" {
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

	first_name := "_"
	last_name := "_"
	user_bio := "_"

	fmt.Print(uname + passwd + email + first_name + last_name + user_bio)

	openDBConnection() //don't forget to open your connection boys and girls
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
		sendToken(email)
		w.WriteHeader(http.StatusOK)
	}
}

func sendToken(recipient string){
	from := "diaversemail@gmail.com" //you must allow for less secure apps on your gmail account
    	//https://support.google.com/accounts/answer/6010255?hl=en
    	pass := emailPass //password goes here

	rand.Seed(time.Now().UnixNano())
    	min := 0
    	max := 10000
    	//fmt.Println(rand.Intn(max - min + 1) + min
	token := (rand.Intn(max - min + 1) + min)

    	msg := "From: " + from + "\n" +
            "To: " + recipient + "\n" +
            "Subject: WELCOME TO DIAVERSE - HERE IS YOUR HARDWARE TOKEN\n\n" +
            strconv.Itoa(token)

    	err := smtp.SendMail("smtp.gmail.com:587",
            smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
            from, []string{recipient}, []byte(msg))

    	if err != nil {
            log.Print(err)
            return
    	}

    	log.Print("sent")
}

func handleExistingHardwareToken(w http.ResponseWriter, r *http.Request){
  //(w).Header().Set("Access-Control-Allow-Origin", "*") //may not need this given CORS above

	//simply check if token is already in DB, if so return true, else return false
	err := r.ParseForm()
  if err != nil {
    fmt.Println(err.Error())
  }
	uname := r.FormValue("username")
	hwToken := r.FormValue("hwtoken")
  fmt.Println(uname);
  fmt.Println(hwToken);

	var result bool
  fmt.Println("Made DB Req")
	hwTokenStruct := getHardwareTokenDataByUsername(uname)
  fmt.Println(hwTokenStruct.Hardwaretoken)
	if hwTokenStruct.Hardwaretoken == hwToken && hwToken != "" {
		result = true
	} else {
		result = false
	}

  fmt.Println("Posted Results to caller")
	if result {
		fmt.Print("ACCEPTED\n")
		w.WriteHeader(http.StatusAccepted) // token is in DB
	} else {
		fmt.Print("INVALID HARDWARE TOKEN\n")
		w.WriteHeader(http.StatusForbidden) // token is not in DB.
	}
}

func handleAuthHardware(w http.ResponseWriter, r *http.Request){
	//I'm imaganing we will enter a registration number into the DB manually
	//verify that the value passed matches the value in the DB
	//if yes return accepted, if false return forbidden
}

//return all the scripts in the database associated with a given hardware token
func handleGetScripts(w http.ResponseWriter, r *http.Request){
	openDBConnection()
	r.ParseForm()

	hwid := r.FormValue("hardwareToken")
	selDB, err := db.Query("SELECT * FROM scripts WHERE hardwareToken =  ?", hwid)
	if err != nil {
		fmt.Println(err.Error())
	}

	scripts := make(map[int]scriptData)
	i := 0
	for selDB.Next(){
		var scriptID int
		var hwToken, script string
		err = selDB.Scan(&scriptID, &hwToken, &script)
		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
		}
		var tempScript scriptData
		tempScript.Hardwaretoken = hwToken
		tempScript.ScriptID = scriptID
		tempScript.Script = script
		scripts[i] = tempScript
		i++
	}

	j, e := json.Marshal(scripts)
	if e != nil {
		fmt.Println(e.Error())
		w.WriteHeader(http.StatusBadRequest)
	}

	jString := string(j)

	w.Write([]byte(jString))
}


func handleLogout(w http.ResponseWriter, r *http.Request){
	// save a logged in or logged out state in the database
	// if this is called, set the state to logged out
}

// based on the user's username - return all info about them from DB
func handleGetUserInfo(w http.ResponseWriter, r *http.Request){
	openDBConnection()
	r.ParseForm()
	uname := r.FormValue("username")
	userInfo := getUserDataByUsername(uname)

	j, e := json.Marshal(userInfo)
	if e != nil {
		fmt.Println(e.Error())
		w.WriteHeader(http.StatusBadRequest)
	}

	jString := string(j)
	w.Write([]byte(jString))
}

func handleScriptResults(w http.ResponseWriter, r *http.Request){
	// based on the hardware token - return the results of the test run.
}

func handleUpdateUserInfo(w http.ResponseWriter, r *http.Request){
	// the writer should contain all the info pertaining ot the update
	// simply update all fields in the databased based on this
}

// based on the users username delete all assocaited information from the database.
func handleDeleteAccount(w http.ResponseWriter, r *http.Request){

}

//delete a script based on the scriptID
func handleDeleteScript(w http.ResponseWriter, r *http.Request){

}

// the script will be in the text of the response writer
// add the script as a new entry in the database.
func handleUploadScript(w http.ResponseWriter, r *http.Request){

}
