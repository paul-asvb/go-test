package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type request_query struct {
	Id         int      `json:"id"`
	Stringlist []string `json:"stringlist"`
}

func YourHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world")

	res1D := &request_query{
		Id:         1,
		Stringlist: []string{"apple", "peach", "pear"}}

	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	str := `{"id": 1, "stringlist": ["apple", "peach"]}`
	res := request_query{}

	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)

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


	g := gin.Default()
    	g.GET("/ping", func(c *gin.Context) {
    		c.JSON(200, gin.H{
    			"message": "pong",
    		})
    	})
    	g.Run()

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
