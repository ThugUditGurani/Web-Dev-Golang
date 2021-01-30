package app

import (
	"Github/Web-Dev-Golang/Web-Dev-Golang/banking-auth/domain"
	"Github/Web-Dev-Golang/Web-Dev-Golang/banking-auth/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"time"
)

func Start() {
	router := mux.NewRouter()
	authRepository := domain.NewAuthRepository(getDBClient())
	ah := AuthHandler{service.NewLoginService(authRepository, domain.GetRolePermissions())}

	router.HandleFunc("/auth/login", ah.Login).Methods(http.MethodPost)
	router.HandleFunc("/auth/register", ah.NotImplementedHandler).Methods(http.MethodPost)
	router.HandleFunc("/auth/verify", ah.Verify).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}


func getDBClient() *sqlx.DB {
	client,err := sqlx.Open("mysql","root:Iamhotudit@123@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
