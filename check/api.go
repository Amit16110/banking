package check

import (
	"log"
	"net/http"
)

func Start(){

	mux:= http.NewServeMux()

	mux.HandleFunc("/greet", greet)  //define routes
	mux.HandleFunc("/customers", getAllCustomers)  //define routes

	println("It's running")
	// for error handling we use the log.fatel for showing the and kind of errors 
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}