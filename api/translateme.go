package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w,
		"Hi, This is an example of https service in golang!")
}

func main() {
	//http.HandleFunc("/", http.FileServer(http.Dir("my-app/browser/index.html")), )
	_ = http.ListenAndServe(":8081", http.FileServer(http.Dir("my-app/browser/")))

}
