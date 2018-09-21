package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type request_query struct {
	Id         int      `json:"id"`
	Stringlist []string `json:"stringlist"`
}

func Pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Pong",
	})
}

func YourHandler(c *gin.Context) {
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

	c.JSON(200, res1B)
}

func main() {

	router := gin.Default()
	router.GET("/ping", Pong)

	router.GET("/", YourHandler)
	router.Run()

	// Bind to a port and pass our router in
	//log.Fatal(http.ListenAndServe(":8000", router))
}
