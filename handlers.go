package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "Root...")
}

func HandleApi(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "API Endpoint...")
}

func PostRequest(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var metadata MetaData
	err := decoder.Decode(&metadata)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	fmt.Fprintf(w, "Payload %v\n", metadata)
}

func UserPostRequest(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	response, err := user.ToJson()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
	// fmt.Println(user.FirstName)
	// fmt.Fprintf(w, "Payload %v\n", user)
}
