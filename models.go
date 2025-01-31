package main

// Request structs
type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

// Response structs
type createReceiptResponse struct {
	ID string `json:"id"`
}

type getPointsResponse struct {
	Points int `json:"points"`
}
