package main

import (
        "database/sql"
        "fmt"
        _ "github.com/go-sql-driver/mysql"
)

type UserData struct {
    Username, Password, Email, Hardwaretoken, FirstName, LastName, UserBio string
}

var db *sql.DB
var err error

func main() {
        println("running")

        //open the database connection
        db, err := sql.Open("mysql", "DB CONNECTION STRING")
        if err != nil {
            fmt.Println(err.Error())
        }

        //example of querying a row from the database
        var userDataStruct UserData
        usersName := "william"
        err = db.QueryRow("SELECT username, password, email, hardwaretoken, first_name, last_name, user_bio FROM users WHERE username = ?", usersName).Scan(&userDataStruct.Username, &userDataStruct.Password, &userDataStruct.Email, &userDataStruct.Hardwaretoken, &userDataStruct.FirstName, &userDataStruct.LastName, &userDataStruct.UserBio)
        if err != nil {
            fmt.Println(err.Error())
        }
        fmt.Printf("Username: %s\n Password: %s\n Email: %s\n HardwareToken: %s\n FirstName: %s\n LastName: %s\n UserBio: %s\n", userDataStruct.Username, userDataStruct.Password, userDataStruct.Email, userDataStruct.Hardwaretoken, userDataStruct.FirstName, userDataStruct.LastName, userDataStruct.UserBio)

        //example of inserting data into the databse
        username := "fakeuser2"
	      password := "password"
	      email := "fakeuser2@wit.edu"
	      first_name := "fake"
	      last_name := "user2"
	      user_bio := "he's not real"
	      insForm, err := db.Prepare("INSERT INTO users (username, password, email, first_name, last_name, user_bio) VALUES(?,?,?,?,?,?)")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(username, password, email, first_name, last_name, user_bio)

	      print("done\n")
}
