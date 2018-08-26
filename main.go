package main

import (
	"net/http"

	"fmt"
)

func main()  {
	http.HandleFunc("/",test)
	fmt.Println("starting")
	http.ListenAndServe(":8080",nil)

}
func test(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w,"hello world")
}
