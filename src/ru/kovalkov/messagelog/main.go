package main

import (
    "net/http"
    "fmt"
    "log"
    "os"
)

const SERVER_HOST = ":8080"
const LOG_FILE = "paymessages.log"

type MyHandler struct {}
var Tube chan string
var file *os.File

func main() {
    var fooHandler *MyHandler
    Tube = make(chan string, 100)
    go saveMessage()

    var err error
    file, err = os.OpenFile(LOG_FILE, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
    if err != nil {
            log.Fatal(err)
    }

    http.Handle("/log", fooHandler)
    log.Fatal(http.ListenAndServe(SERVER_HOST, nil))
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    err := r.ParseForm()
    if err != nil {
        fmt.Fprintf(w, "fail")
        log.Fatal(err)
    }
    message := r.Form["msg"][0]
    fmt.Fprintf(w, "ok")
    if "" != message {
        Tube <- message
    }
}

func saveMessage() {
    for {
        file.WriteString(<-Tube + "\n")
    }
}
