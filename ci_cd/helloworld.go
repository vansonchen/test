package main

import (
    "fmt"
    "net/http"
    "log"  
)

func myHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello\n")
}

func main(){
    http.HandleFunc("/", myHandler)     
    log.Fatal(http.ListenAndServe(":8080", nil))
}
