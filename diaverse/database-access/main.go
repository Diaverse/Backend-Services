package main

import "database/sql"
import "fmt"
import _ "github.com/go-sql-driver/mysql"

func main() {
        print("running")

        //open the database connection
        db, err := sql.Open("mysql", "USERNAME-GOES-HERE:PASSWORD-GOES-HERE@tcp(ENDPOINT:PORT-NUMBER)/DATABSE-NAME")
        if err != nil {
            panic(err)
        }

        //select a single row
        //id := "username"
        var col string
        sqlStatment := `SELECT * FROM users`
        row  := db.QueryRow(sqlStatment)
        err  = row.Scan(&col)
        if err != nil {
            if err == sql.ErrNoRows {
                fmt.Println("Zero rows found")
            } else {
                panic(err)
        }
    }
}
