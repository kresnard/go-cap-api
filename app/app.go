package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	mux := mux.NewRouter()

	// mux := http.NewServeMux()

	// * defining routes
	// mux.HandleFunc("/greet", greet)
	// mux.HandleFunc("/customer", getCustomers)
	mux.HandleFunc("/greet", greet).Methods(http.MethodGet)
	mux.HandleFunc("/customer", getAllCustomers).Methods(http.MethodGet)
	mux.HandleFunc("/customer/{customer_id}", getCustomers).Methods(http.MethodGet)

	// * starting the server
	http.ListenAndServe(":8080", mux)
}
