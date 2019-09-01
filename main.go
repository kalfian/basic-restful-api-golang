package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// @export Article Title, Desc, Content
type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{"Title article", "Desc article", "Content Article"},
	}

	fmt.Println("Endpoint hit: All Article Endpoint")
	json.NewEncoder(w).Encode(articles)

}

func postArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: Post Article Endpoint")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page Endpoint")
}

func handleRequest() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/articles", allArticles).Methods("GET")
	router.HandleFunc("/articles", postArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequest()
}
