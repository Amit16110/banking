package check

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// there is something new if we want it to send json data with name changes it do it here
type Customer struct {
	Name string  `json:"Full_name"`
	City string	`json:"State"`	
	Zipcode string	`json:"Code"`
}
// marshalling data structures to xml representation



func greet (rw http.ResponseWriter, r *http.Request) {
	// use fmt.Fprint for send response to the frontend
	fmt.Fprint(rw, "hello work from the go")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request){
	customers := []Customer{
		{Name: "Amit", City: "New Delhi" , Zipcode: "110041"},
		{Name: "Tannu", City: "New Delhi" , Zipcode: "110041"},
		{Name: "Parwati", City: "New Delhi" , Zipcode: "110041"},
	}

	w.Header().Add("Content-Type", "application/json")  // content type for send response 

	json.NewEncoder(w).Encode(customers)  //it's default go lang json encoder it changes in json format
}