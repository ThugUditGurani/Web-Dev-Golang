package domain

import (
	"Github/Web-Dev-Golang/Web-Dev-Golang/banking/errs"
	"Github/Web-Dev-Golang/Web-Dev-Golang/banking/logger"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}


func (d CustomerRepositoryDB) FindAll(status string) ([]Customer,*errs.AppError)  {


	//var rows *sql.Rows
	var err error
	customers := make([]Customer,0)

	if status == "" {
		findAllSql := "select customer_id, full_name, city, zipcode, date_of_birth, status from customer"
		err = d.client.Select(&customers,findAllSql)
		//rows,err = d.client.Query(findAllSql)
	}else {
		findAllSql := "select customer_id, full_name, city, zipcode, date_of_birth, status from customer where status = ?"
		err = d.client.Select(&customers,findAllSql,status)
		//rows,err = d.client.Query(findAllSql, status)
	}


	if err != nil {
		logger.Error("Error while quering customer tables"+err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}


	//err = sqlx.StructScan(rows,&customers)
	//if err != nil {
	//	logger.Error("Error while scanning customer"+err.Error())
	//	return nil, errs.NewUnexpectedError("Unexpected database error")
	//}

	//for rows.Next() {
	//	var c Customer
	//	err := rows.Scan(&c.Id , &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
	//
	//	if err != nil {
	//		logger.Error("Error while scanning customer"+err.Error())
	//		return nil, errs.NewUnexpectedError("Unexpected database error")
	//	}
	//	customers = append(customers,c)
	//}
	return  customers,nil
}

func (d CustomerRepositoryDB) ById(id string) (*Customer,*errs.AppError)  {
	customerSql := "select customer_id, full_name, city, zipcode, date_of_birth, status from customer where customer_id=?"
	//row := d.client.QueryRow(customerSql,id)
	var c Customer
	err :=d.client.Get(&c,customerSql,id)
	//err := row.Scan(&c.Id , &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NerNotFoundError("Customer not found")
		}else{
			log.Println("Error while scanning customer "+err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &c,nil
}

func NewCustomerRepositoryDB(dbClient *sqlx.DB) CustomerRepositoryDB {
	return CustomerRepositoryDB{client: dbClient}
}