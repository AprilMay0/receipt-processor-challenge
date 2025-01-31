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

type createReceiptResponse struct {
	ID string
}

type getPointsResponse struct {
	Points int
}

// Takes reciept details, processes, and returns new ID
func CreateReceipt(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Receipt")
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
func GetReceipt(w http.ResponseWriter, r *http.Request, id string) {
	fmt.Println("Checking AllReceipts: ", AllReceipts)
	fmt.Printf("Receipt id %s generated %d points\n", id, AllReceipts[id])

	resp := getPointsResponse{
		Points: AllReceipts[id],
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
