package domain

// CustomerRespositoryStub ..
type CustomerRespositoryStub struct {
	customers []Customer
}

// FindAll ..
func (s CustomerRespositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

// NewCustomerRespositoryStub ..
func NewCustomerRespositoryStub() CustomerRespositoryStub {
	customers := []Customer{
		{
			Id:          "1",
			Name:        "Rahul",
			City:        "Kochi",
			Zipcode:     "682018",
			DateofBirth: "2000-01-01",
			Status:      "1",
		},
		{
			Id:          "2",
			Name:        "Raj",
			City:        "Kochi",
			Zipcode:     "682020",
			DateofBirth: "1998-01-01",
			Status:      "1",
		},
	}

	return CustomerRespositoryStub{customers: customers}
}
