package main

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, q *http.Request) {
	fmt.Println(q.Method)
	fmt.Println(q.URL)
	fmt.Fprintf(w, "OK\n")
}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/hello", Index)

	s := http.Server{Addr: "127.0.0.1:8080"}
	err := s.ListenAndServe()
	if err != nil {
		fmt.Printf("err : %v\n", err)
	}
}
