package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", rootHandler)
	serve()
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := json.MarshalIndent(struct {
		Message string `json:"message"`
	}{
		Message: "Hello World",
	}, " ", " ")
	w.Write(data)
}

func serve() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	addr := ":" + port
	fmt.Printf("Listening %s\n", addr)
	http.ListenAndServe(addr, nil)
}
