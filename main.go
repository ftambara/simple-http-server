package main

import (
    "net/http"
    "log"
)

func home(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello World!"))
}

func main(){
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)

    port := ":4000"
    log.Printf("Starting server on %v", port)
    err := http.ListenAndServe(port, mux)
    if err != nil {
        log.Fatal(err)
    }
}

