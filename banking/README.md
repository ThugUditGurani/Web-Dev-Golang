
Architecture :- Hexagonal Architecture

Steps:- 

1. Run database.sql file in MySql WorkBench.
2. Go Get Command to install modules.

Tools :- 
1. MySql
2. Gorilla/mux
3. Mockito

API Documentation :-

1. Get Customer API

Request-type -- GET
URL -- http://localhost:8000/customers


2.  Account Creation

Requesttype -- POST
URL -- http://localhost:8000/customers/{customer_id:[0-9]+}/account
RequestBody -
{
    "account_type" : "saving",
    "amount": 5000
}

3. Make Transaction

RequestType -- POST
URL -- http://localhost:8000/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}
Request body --
{
    "transaction_type" : "deposit",
    "amount" : 1000
}
