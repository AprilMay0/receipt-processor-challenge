package main

import (
	"encoding/json"
	"fmt"
	"log"
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

	// TODO: REMOVE
	marshaled, err := json.MarshalIndent(receipt, "", "   ")
	if err != nil {
		log.Fatalf("marshaling error: %s", err)
	}
	fmt.Println(string(marshaled))

	// create unique ID to assign to receipt
	byteID, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}

	id := strings.TrimSpace(string(byteID))

	// TODO add actual points
	AllReceipts[id] = 100

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
