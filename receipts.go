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
	logger.Println("Processing receipt")
	byteID, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}

	id := strings.TrimSpace(string(byteID))

	// TODO add actual points
	AllReceipts[id] = 100
	fmt.Println("ID: ", id)

	fmt.Println("AllReceipts: ", AllReceipts)

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
		w.WriteHeader(http.StatusNotFound)
		response := ErrorResp{Status: http.StatusNotFound, Message: "Receipt not found"}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			logger.Printf("Error encoding response in GetReceipt")
		}
		return
	}

	resp := getPointsResponse{
		Points: existingPoints,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

// Response structs
type createReceiptResponse struct {
	ID string
}

type getPointsResponse struct {
	Points int
}

type ErrorResp struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
