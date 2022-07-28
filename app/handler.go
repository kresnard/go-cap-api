package app

import (
	"capi/service"
	"encoding/json"
	"encoding/xml"
	"net/http"

	// "github.com/go-delve/delve/service"
	"github.com/gorilla/mux"
)

// type Customer struct {
// 	ID      int    `json:"id" xml:"id"`
// 	Name    string `json:"name" xml:"name"`
// 	City    string `json:"city" xml:"city"`
// 	ZipCode string `json:"zip_code" xml:"zipcode"`
// }

// var customers []Customer = []Customer{
// 	{1, "User1", "Jakarta", "123456"},
// 	{2, "User2", "Surabaya", "67890"},
// }

// func greet(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Hello Celerates!")
// }

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "get customerendpoint")

	customers, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func (ch *CustomerHandler) getCustomerByID(w http.ResponseWriter, r *http.Request) {

	// get route variable
	vars := mux.Vars(r)

	customerId := vars["customer_id"]

	customer, err := ch.service.GetCustomerByID(customerId)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}

	// return customer data
	writeResponse(w, http.StatusOK, customer)

}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

// func addCustomer(w http.ResponseWriter, r *http.Request) {
// 	// decode request body
// 	var cust Customer
// 	json.NewDecoder(r.Body).Decode(&cust)

// 	// generate new id
// 	nextID := getNextID()
// 	cust.ID = nextID

// 	// save data to array
// 	customers = append(customers, cust)

// 	w.WriteHeader(http.StatusCreated)
// 	fmt.Fprintln(w, "customer succesfully created")

// 	// fmt.Println(cust)
// }

// func getNextID() int {
// 	cust := customers[len(customers)-1]

// 	return cust.ID + 1
// }

// func updateCustomers(w http.ResponseWriter, r *http.Request) {
// 	// get route variable
// 	vars := mux.Vars(r)

// 	customerId := vars["customer_id"]

// 	// convert str to int
// 	// id, _ := strconv.Atoi(customerId)
// 	id, err := strconv.Atoi(customerId)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		fmt.Fprint(w, "invalid customer id")
// 	}

// 	// searching customer data
// 	var cust Customer

// 	for customerIndex, data := range customers {
// 		if data.ID == id {

// 			// save temp data for validation
// 			cust = data

// 			// decode request body
// 			var newCust Customer
// 			json.NewDecoder(r.Body).Decode(&newCust)

// 			// do update
// 			customers[customerIndex].Name = newCust.Name
// 			customers[customerIndex].City = newCust.City
// 			customers[customerIndex].ZipCode = newCust.ZipCode

// 			w.WriteHeader(http.StatusOK)
// 			fmt.Fprint(w, "customer data updated")

// 		}
// 	}
// 	if cust.ID == 0 {
// 		w.WriteHeader(http.StatusNotFound)
// 		fmt.Fprint(w, "customer data not found")
// 		return
// 	}

// }

// func deleteCustomers(w http.ResponseWriter, r *http.Request) {
// 	// get route variable
// 	vars := mux.Vars(r)
// 	customerId := vars["customer_id"]

// 	// convert str to int
// 	id, _ := strconv.Atoi(customerId)
// 	// id, err := strconv.Atoi(customerId)
// 	// if err != nil {
// 	// 	w.WriteHeader(http.StatusBadRequest)
// 	// 	fmt.Fprint(w, "invalid customer id")
// 	// }

// 	// searching customer data
// 	var updateCustomers Customer
// 	json.NewDecoder(r.Body).Decode(&updateCustomers)

// 	for i, cust := range customers {
// 		if cust.ID == id {
// 			customers = append(customers[:i], customers[i+1:]...)
// 			// customers = append(customers, updateCustomers)
// 		}
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprint(w, "customer data deleted")
// }

// func updateCustomers(w http.ResponseWriter, r *http.Request) {
// 	// get route variable
// 	vars := mux.Vars(r)

// 	customerId := vars["customer_id"]

// 	// convert str to int
// 	id, _ := strconv.Atoi(customerId)
// 	// id, err := strconv.Atoi(customerId)
// 	// if err != nil {
// 	// 	w.WriteHeader(http.StatusBadRequest)
// 	// 	fmt.Fprint(w, "invalid customer id")
// 	// }

// 	// searching customer data
// 	var updateCustomers Customer

// 	json.NewDecoder(r.Body).Decode(&updateCustomers)

// 	for i, cust := range customers {
// 		if cust.ID == id {
// 			customers = append(customers[:i], customers[i+1:]...)
// 			customers = append(customers, updateCustomers)
// 		}
// 	}
// 	json.NewEncoder(w).Encode(customers)
// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintln(w, "customer succesfully updated")

// }
