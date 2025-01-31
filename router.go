package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

// define routes
func RequestHandler() {
	http.HandleFunc("/receipts/process", httpLogger(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			ProcessReceipt(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))
	http.HandleFunc("/receipts/{id}/points", httpLogger(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			GetReceipt(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))

	logger.Print("Starting server...")
	http.ListenAndServe(":8080", nil)
}

// define logger
var logger = log.New(os.Stdout, "[log] ", 0)

func httpLogger(fn func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		logger.Printf("Starting %s %s", r.Method, r.URL.Path)
		fn(w, r)
		logger.Printf("Completed in %v (%s %s)", time.Since(start), r.Method, r.URL.Path)
	}
}
