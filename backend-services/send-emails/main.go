//this file is the sample code used to setup token emails in go for diaverse
package main

import (
    "log"
    "net/smtp"
)

func main() {
    send("HERE IS YOUR TOKEN")
}

func send(body string) {
    from := "diaversemail@gmail.com" //you must allow for less secure apps on your gmail account
    //https://support.google.com/accounts/answer/6010255?hl=en
    pass := "" //password goes here
    to := "shuew@wit.edu"

    msg := "From: " + from + "\n" +
            "To: " + to + "\n" +
            "Subject: Hello there\n\n" +
            body

    err := smtp.SendMail("smtp.gmail.com:587",
            smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
            from, []string{to}, []byte(msg))

    if err != nil {
            log.Print(err)
            return
    }

    log.Print("sent")
}
