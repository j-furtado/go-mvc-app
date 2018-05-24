package main

import (
	"fmt"
	"net/http"
	"os"
)

var db_host = ""

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "The host is ", db_host)
}

func main() {
	db_host = os.Args[1]
	http.HandleFunc("/", index)
	port := ":8080"
	http.ListenAndServe(port, nil)
}
