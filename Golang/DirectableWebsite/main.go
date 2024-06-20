package main

import (
	"fmt"
	"log"
	"net/http"
)

func coryhandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST Request Successful")
	name := r.FormValue("name")
	fmt.Fprintf(w, "Name = %s\n", name)
}

func cxkhandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/cxk" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "metod is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "welcome")
}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/cory", coryhandler)
	http.HandleFunc("/cxk", cxkhandler)

	fmt.Printf("starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
