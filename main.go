package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"bytes"
)


func HomeHandler (w http.ResponseWriter, r *http.Request) {
	var buffer bytes.Buffer
	for i:=0; i < 1000000; i++ {
		buffer.WriteString("Hähß")
	}
	fmt.Fprint(w, buffer.String())
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", nil))
}