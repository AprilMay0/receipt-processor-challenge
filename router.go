package main

import (
	"net/http"
)

// define routes
func RequestHandler() {
	http.HandleFunc("/receipts/process", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			CreateReceipt(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/receipts/{id}/points", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			GetReceipt(w, r, r.PathValue("id"))
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}
