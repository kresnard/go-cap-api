package app

import (
	"capi/domain"
	"capi/logger"
	"capi/service"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func sanityCheck() {
	envProps := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
	}
	for _, envKey := range envProps {
		if os.Getenv(envKey) == "" {
			logger.Fatal(fmt.Sprintf("environment variable %s not defined. terminating appliacation ...", os.Getenv(envKey)))
		}
	}
}

func Start() {

	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file")
	}
	sanityCheck()

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
	serverAddr := os.Getenv("SERVER_ADDRESS")
	ServerPort := os.Getenv("SERVER_PORT")

	logger.Info(fmt.Sprintf("Start server on %s:%s ...", serverAddr, ServerPort))
	http.ListenAndServe(fmt.Sprintf("%s:%s", serverAddr, ServerPort), mux)
}
