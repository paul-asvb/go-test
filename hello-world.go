package main
import "fmt"
import "net/http"

func main() {
    fmt.Println("hello world")
    resp, err := http.Get("http://example.com/")
    if err != nil {
    	// handle error
    	fmt.Println(err)

    }
    fmt.Println(resp)
}


