package domain

// Customer ..
type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string
}

// CustomerRespository ..
type CustomerRespository interface {
	FindAll() ([]Customer, error)
}
