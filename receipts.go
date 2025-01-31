package main

import (
	"encoding/json"
	"net/http"
	"os/exec"
	"strings"
)

// AllReceipts[ID]Points
var AllReceipts = make(map[string]int)

// Takes reciept details, processes, and returns new ID
func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	// validate body before assigning ID and processing points
	var receipt Receipt
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&receipt); err != nil {
		http.Error(w, "The receipt is invalid.", http.StatusBadRequest)
		return
	}

	// get points total for receipt
	totalPoints, err := Process(receipt)
	if err != nil {
		http.Error(w, "The receipt is invalid.", http.StatusBadRequest)
		return
	}

	// create unique ID to assign to receipt
	byteID, err := exec.Command("uuidgen").Output()
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		logger.Printf("Error creating receipt ID: %v", err)
		return
	}

	id := strings.TrimSpace(string(byteID))

	// add new id and total points earned to the map
	AllReceipts[id] = totalPoints

	resp := createReceiptResponse{
		ID: id,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

// Takes ID for receipt and returns points
func GetReceipt(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	existingPoints, ok := AllReceipts[id]

	// if receipt with id doesn't exist, return with NotFound status
	if !ok {
		http.Error(w, "No receipt found for that ID.", http.StatusNotFound)
		return
	}

	resp := getPointsResponse{
		Points: existingPoints,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
