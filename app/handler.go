package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Customer struct {
	ID      int    `json:"id" xml:"id"`
	Name    string `json:"name" xml:"name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zipcode"`
}

var customers []Customer = []Customer{
	{1, "User1", "Jakarta", "123456"},
	{2, "User2", "Surabaya", "67890"},
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Celerates!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "get customerendpoint")

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func getCustomers(w http.ResponseWriter, r *http.Request) {

	// get route variable
	vars := mux.Vars(r)

	customerId := vars["customer_id"]

	// convert str to int
	id, _ := strconv.Atoi(customerId)

	// searching customer data
	var cust Customer

	for _, data := range customers {
		if data.ID == id {
			cust = data
		}
	}

	// return customer data
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cust)

}
