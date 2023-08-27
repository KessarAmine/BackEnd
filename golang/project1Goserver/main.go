package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello_route(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Not allowed", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "herzlich willkommen")
}

func form_route(w http.ResponseWriter, r *http.Request) {
	if error := r.ParseForm(); error != nil {
		fmt.Fprintf(w, "ParseForm error: %v\n", error)
		return
	}
	fmt.Fprintf(w, "Post request done\n")
	name := r.FormValue("name")
	adresse := r.FormValue("adresse")
	fmt.Fprintf(w, "name = %s\nadresse= %s\n", name, adresse)
}

func main() {
	file_server := http.FileServer(http.Dir("./static"))
	http.Handle("/", file_server)
	http.HandleFunc("/form", form_route)
	http.HandleFunc("/hello", hello_route)

	fmt.Printf("server started\n")
	if error := http.ListenAndServe(":8080", nil); error != nil {
		log.Fatal(error)
	}

}
