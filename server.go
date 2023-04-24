package main


import (
    "fmt"
    "log"
    "net/http"
)


func main() {
    fileServer := http.FileServer(http.Dir("./static"))
    http.Handle("/", fileServer)



    fmt.Printf("Starting server at port 127.0.0.1:5555\nOpen 127.0.0.1:5555/index.htm in your browser to run the application.\nVisit https://www.elearningllc.com/scormwraptest/docs to see how to use this application.")
    if err := http.ListenAndServe(":5555", nil); err != nil {
        log.Fatal(err)
    }
}