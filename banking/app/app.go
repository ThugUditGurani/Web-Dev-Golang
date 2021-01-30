package app

import (
	"Github/Web-Dev-Golang/Web-Dev-Golang/banking/domain"
	"Github/Web-Dev-Golang/Web-Dev-Golang/banking/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"time"
)

func Start() {

	//mux := http.NewServeMux();
	router := mux.NewRouter()

	//wiring
	dbClient := getDBClient()
	customerRepositoryDb := domain.NewCustomerRepositoryDB(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)
	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandler{service.NewAccountService(accountRepositoryDb)}
	// define Routes
	router.HandleFunc("/customers",ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}",ch.getCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account",ah.NewAccount).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}",ah.MakeTransaction).Methods(http.MethodPost)


	//starting Server
	log.Fatal(http.ListenAndServe("localhost:8000",router))
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