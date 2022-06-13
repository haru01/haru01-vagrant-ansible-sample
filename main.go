package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Go App")
}

func main() {
	http.HandleFunc("/hello", handler)
	fmt.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
