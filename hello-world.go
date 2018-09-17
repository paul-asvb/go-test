package main

import (
    "fmt"
    "net/http"
    "log"
    "github.com/gorilla/mux"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello World!\n"))
}

func main() {

        fmt.Println("hello world")
        resp, err := http.Get("http://example.com/")
        if err != nil {
        	// handle error
        	fmt.Println(err)

        }
        fmt.Println(resp)

    r := mux.NewRouter()
    // Routes consist of a path and a handler function.
    r.HandleFunc("/", YourHandler)

    // Bind to a port and pass our router in
    log.Fatal(http.ListenAndServe(":8000", r))
}
