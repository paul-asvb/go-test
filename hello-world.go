package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"strconv"
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

	// new object
	res1D := &request_query{
		Id:         1,
		Stringlist: []string{"apple", "peach", "pear"}}

	// obj -> json
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// string -> obj
	str := `{"id": 1, "stringlist": ["apple", "peach"]}`
	res := request_query{}
	json.Unmarshal([]byte(str), &res)

	fmt.Println("Print res", res)

	resp, err := http.Get("http://example.com/")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)

	c.JSON(200, res1D)
}

func main() {

	database, _ := sql.Open("sqlite3", "./test.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")
	statement.Exec("Nic", "Raboy")
	rows, _ := database.Query("SELECT id, firstname, lastname FROM people")
	var id int
	var firstname string
	var lastname string
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname)
		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
	}


	router := gin.Default()
	router.GET("/ping", Pong)

	router.GET("/", YourHandler)
	router.Run()


	// Bind to a port and pass our router in
	//log.Fatal(http.ListenAndServe(":8000", router))
}
