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
        db, err := sql.Open("mysql", "")
        if err != nil {
            fmt.Println(err.Error())
        }

        var userDataStruct UserData
        usersName := "william"
        err = db.QueryRow("SELECT username, password, email, hardwaretoken, first_name, last_name, user_bio FROM users WHERE username = ?", usersName).Scan(&userDataStruct.Username, &userDataStruct.Password, &userDataStruct.Email, &userDataStruct.Hardwaretoken, &userDataStruct.FirstName, &userDataStruct.LastName, &userDataStruct.UserBio)
        if err != nil {
            fmt.Println(err.Error())
        }

        fmt.Printf("Username: %s\n Password: %s\n Email: %s\n HardwareToken: %s\n FirstName: %s\n LastName: %s\n UserBio: %s\n", userDataStruct.Username, userDataStruct.Password, userDataStruct.Email, userDataStruct.Hardwaretoken, userDataStruct.FirstName, userDataStruct.LastName, userDataStruct.UserBio)


        //return err
}
