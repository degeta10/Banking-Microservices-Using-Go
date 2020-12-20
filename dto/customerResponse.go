package dto

// CustomerResponse ..
type CustomerResponse struct {
	Id          string `json:"customer_id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateofBirth string `json:"dob"`
	Status      string `json:"status"`
}
