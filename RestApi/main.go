package main

import (
	"Github/Web-Dev-Golang/Web-Dev-Golang/RestApi/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout,"product-api",log.LstdFlags)
	hh := handlers.NewHello(l)
	gg := handlers.NewGoodBye(l)
	pp := handlers.NewProduct(l)


	sm := http.NewServeMux()
	sm.Handle("/",hh)
	sm.Handle("/products",pp)
	sm.Handle("/goodbye",gg)
	//
	//s := &http.Server{
	//	Addr:              "8080",
	//	Handler:           sm,
	//	TLSConfig:         nil,
	//	ReadTimeout:       1 * time.Second,
	//	ReadHeaderTimeout: 1 * time.Second,
	//	WriteTimeout:      1 * time.Second,
	//	IdleTimeout:       120 * time.Second,
	//}
	//
	//go func() {
	//	err := s.ListenAndServe()
	//	if err != nil {
	//		l.Fatal(err)
	//	}
	//}()
	//s.ListenAndServe()
	//sigChan := make(chan os.Signal)
	//signal.Notify(sigChan,os.Interrupt)
	//signal.Notify(sigChan,os.Kill)
	//
	//signal := <- sigChan
	//l.Println("Received terminate, graceful shutdown",signal)

	http.ListenAndServe("localhost:8080",sm)

	//tc, _ := context.WithTimeout(context.Background(),30 * time.Second)
	//s.Shutdown(tc)
}
