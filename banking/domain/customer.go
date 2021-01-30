package domain

import (
	"Github/Web-Dev-Golang/Web-Dev-Golang/banking/dto"
	"Github/Web-Dev-Golang/Web-Dev-Golang/banking/errs"
)

type Customer struct {
	Id string  `db:"customer_id"`
	Name string `db:"full_name"`
	City string
	ZipCode string
	DateOfBirth string  `db:"date_of_birth"`
	Status string
}

func (c Customer) statusAsText() string  {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
 	return statusAsText
}

func (c Customer) ToDto() dto.CustomerResponse{


	return dto.CustomerResponse{
		Id: c.Id,
		Name: c.Name,
		City: c.City,
		ZipCode: c.ZipCode,
		DateOfBirth: c.DateOfBirth,
		Status: c.statusAsText(),
	}
}


type CustomerRepository interface {
	FindAll(status string) ([]Customer,*errs.AppError)
	ById(string) (*Customer,*errs.AppError)
}

