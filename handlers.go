package main

import (
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "Root...")
}

func HandleApi(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "API Endpoint...")
}
