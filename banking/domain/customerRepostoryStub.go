package domain


type CustomerRespositoryStub struct {
	customer []Customer
}

func (s CustomerRespositoryStub) FindAll() ([]Customer, error) {
	return s.customer,nil
}


func NewCustomerRepositoryStub() CustomerRespositoryStub  {
	customer := []Customer{
		{"10001","Ashish","Nainital","4234","324","23424"},
		{"10001","Ashish","Nainital","4234","324","23424"},
	}
	return CustomerRespositoryStub{customer: customer}
}
