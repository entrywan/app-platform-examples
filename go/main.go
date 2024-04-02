package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r)
		w.Write([]byte("hello, world!"))
	})
	fmt.Println("listening on port 8000")
	http.ListenAndServe(":8000", nil)
}
