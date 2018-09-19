package main

import (
    "fmt"
    "net/http"
    "log"
    "github.com/gorilla/mux"
    "encoding/json"
)

type request_query struct {
    Id   int      `json:"page"`
    Stringlist []string `json:"fruits"`
}

func YourHandler(w http.ResponseWriter, r *http.Request) {


     res1D := &request_query{
                Id:   1,
                Stringlist: []string{"apple", "peach", "pear"}}

            res1B, _ := json.Marshal(res1D)
            fmt.Println(string(res1B))

            fmt.Println("hello world")
            resp, err := http.Get("http://example.com/")
            if err != nil {
            	// handle error
            	fmt.Println(err)

            }
            fmt.Println(resp)


    w.Write([]byte(res1B))
}


func main() {



    r := mux.NewRouter()
    // Routes consist of a path and a handler function.
    r.HandleFunc("/", YourHandler)

    // Bind to a port and pass our router in
    log.Fatal(http.ListenAndServe(":8000", r))
}
