package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func (h Hello) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	h.l.Println("Hello World")
	data, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(writer,"Hello there no body data in request",http.StatusBadRequest)
		return
	}

	fmt.Fprint(writer,string(data))
}

func NewHello(l *log.Logger) *Hello{
	return &Hello{l: l}
}
