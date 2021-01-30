package main

import (
	"Github/Web-Dev-Golang/Web-Dev-Golang/banking/app"
	"Github/Web-Dev-Golang/Web-Dev-Golang/banking/logger"
)

func main() {

	//http.HandleFunc("/greet", func(writer http.ResponseWriter, request *http.Request) {
	//	fmt.Fprint(writer,"Hello World")
	//})
	logger.Info("Starting the Application")
	app.Start()

}

