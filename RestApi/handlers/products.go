package handlers

import (
	"Github/Web-Dev-Golang/Web-Dev-Golang/RestApi/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func (p *Products) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		p.getProducts(writer,request)
		return
	}

	if request.Method == http.MethodPost {
		p.addProduct(writer,request)
		return
	}

	if request.Method == http.MethodPut {
		r := regexp.MustCompile(`/[0-9]+`)
		g := r.FindAllStringSubmatch(request.URL.Path,-1)
		if len(g[0]) != 1 {
			p.l.Println("Invalid URL more then one id")
			http.Error(writer,"Invalid URI",http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			p.l.Println("Invalid URL more then one capture group")
			http.Error(writer,"Invalid URL",http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			p.l.Println("Invalid URL unable to covert to numer",idString)
			http.Error(writer,"Invalid URL",http.StatusBadRequest)
			return
		}

		//p.l.Println(id)
		p.updateProducts(id, writer,request)
		return
	}

	//catch All
	writer.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(writer http.ResponseWriter,request *http.Request){
	lp := data.GetProducts()
	err := lp.ToJSON(writer)
	if err != nil {
		http.Error(writer,"Unable to marshal json",http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(writer http.ResponseWriter, request *http.Request) {
	p.l.Println("Handle POST Product")
	prod := &data.Product{}
	err := prod.FromJSON(request.Body)
	if err != nil {
		http.Error(writer,"Unable to unmarshal json",http.StatusBadRequest)
	}
	data.AddProduct(prod)

}

func (p *Products) updateProducts(id int, writer http.ResponseWriter,request *http.Request) {
	p.l.Println("Handle PUT Product")
	prod := &data.Product{}
	err := prod.FromJSON(request.Body)
	if err != nil {
		http.Error(writer,"Unable to unmarshal json",http.StatusBadRequest)
	}

	err = data.UpdateProduct(id,prod)
	if err == data.ErrProductNotFound {
		http.Error(writer,"Product not found",http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(writer,"Product not found",http.StatusNotFound)
		return
	}
}

func NewProduct(l *log.Logger) *Products{
	return &Products{l: l}
}