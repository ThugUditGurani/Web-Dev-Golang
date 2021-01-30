package app

import (
	"Github/Web-Dev-Golang/Web-Dev-Golang/banking/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Customer struct {
	Name string  `json:"full_name" xml:"fullName"`
	City string  `json:"city_name" xml:"cityName"`
	ZipCode string `json:"zip_code" xml:"zipCode"`
}


type CustomerHandlers struct {
	service service.CustomerService
}


func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter,r *http.Request){

	status := r.URL.Query().Get("status")
	customers,err := ch.service.GetAllCustomer(status)

	if err != nil{
		writeResponse(w,err.Code,err.AsMessage())
	}else{
		writeResponse(w,http.StatusOK,customers)
	}
}

func (ch *CustomerHandlers) getCustomers(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["customer_id"]
	customer,err := ch.service.GetCustomer(id)
	if err != nil{
		writeResponse(writer,err.Code,err.AsMessage())
	}else{
		writeResponse(writer,http.StatusOK,customer)
	}
}

func writeResponse(writer http.ResponseWriter,code int,data interface{})  {
	writer.Header().Add("Content-Type","application/json")
	writer.WriteHeader(code)
	if err := json.NewEncoder(writer).Encode(data); err != nil {
		panic(err)
	}
}
