package main

import (
	"encoding/json"
	"fmt"
	"gorawg/services"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func search(w http.ResponseWriter, r *http.Request) {
	rawg := services.RawgService()
	resp, err := rawg.SearchGame("")

	if err != nil {
		fmt.Printf("ERROR: %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
	}
	w.Write(jsonResponse)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	fmt.Println("hello from gorawg")

	r := mux.NewRouter()
	r.StrictSlash(true).HandleFunc("/", hello)
	r.StrictSlash(true).HandleFunc("/search", search)
	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
