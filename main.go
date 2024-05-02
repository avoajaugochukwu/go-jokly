package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", getHome)
	http.HandleFunc("/joke", getJoke)
	http.HandleFunc("/madlib", getMadlib)

	log.Println("Listening on port http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func getHome(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"joke":   "http://localhost:8080/joke",
		"madlib": "http://localhost:8080/madlib",
	}
	writeJson(w, response)
}

func getJoke(w http.ResponseWriter, r *http.Request) {
	writeJson(w, GetRandomJoke())
}

func getMadlib(w http.ResponseWriter, r *http.Request) {
	writeJson(w, GetRandomMadlib())
}

func writeJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	jsonResponse, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}
