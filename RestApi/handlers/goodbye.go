package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func (g *Goodbye) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer,"Good bye")
}

func NewGoodBye(l *log.Logger) *Goodbye  {
	return &Goodbye{l: l}
}
