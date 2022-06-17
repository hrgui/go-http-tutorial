package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// write the JSON header
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	// make the JSON object
	resp := make(map[string]string)
	resp["message"] = "Hello World"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	// return back the JSON response
	w.Write(jsonResp)
	return
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/json", handleRequest)
	
	fmt.Println("Server started.")
	http.ListenAndServe(":8090", nil)
}
