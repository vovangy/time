package main

import (
	"fmt"
	"net/http"
	"time"
)

func lol() int {
	return 7
}

func handler(w http.ResponseWriter, r *http.Request) {
	document := "<!DOCTYPE html><html><head><title>Время</title></head><body><h1>Время</h1><p>" + string(time.Now().Format("2006-01-02 15:04:05")) + "</p></body></html>"
	fmt.Fprintf(w, document)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
