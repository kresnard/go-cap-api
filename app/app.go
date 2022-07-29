package app

import (
	"capi/domain"
	"capi/service"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	mux := mux.NewRouter()

	// wiring
	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryDB())}

	// mux := http.NewServeMux()

	// * defining routes
	// mux.HandleFunc("/greet", greet)
	// mux.HandleFunc("/customer", getCustomers)
	// mux.HandleFunc("/greet", greet).Methods(http.MethodGet)
	mux.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	// mux.HandleFunc("/customer", addCustomer).Methods(http.MethodPost)
	// mux.HandleFunc("/customer/{customer_id}", getCustomers).Methods(http.MethodGet)
	mux.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomerByID).Methods(http.MethodGet)
	// mux.HandleFunc("/customer/{customer_id:[0-9]+}", updateCustomers).Methods(http.MethodPut)
	// mux.HandleFunc("/customer/{customer_id:[0-9]+}", deleteCustomers).Methods(http.MethodDelete)

	// * starting the server
	http.ListenAndServe(":8080", mux)
}
