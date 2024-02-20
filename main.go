package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// PaymentService represents the payment service
type PaymentService struct {
	Router *mux.Router
}

// NewPaymentService creates a new payment service
func NewPaymentService() *PaymentService {
	service := &PaymentService{
		Router: mux.NewRouter(),
	}

	service.Routes()

	return service
}

// Routes defines the routes for the payment service
func (s *PaymentService) Routes() {
	s.Router.HandleFunc("/payments", s.CreatePayment).Methods("POST")
	s.Router.HandleFunc("/payments/{id}", s.GetPayment).Methods("GET")
}

// CreatePayment handles the creation of a new payment
func (s *PaymentService) CreatePayment(w http.ResponseWriter, r *http.Request) {
	// Implement your logic here
	fmt.Fprint(w, "Payment created")
}

// GetPayment handles the retrieval of a payment by its ID
func (s *PaymentService) GetPayment(w http.ResponseWriter, r *http.Request) {
	// Implement your logic here
	fmt.Fprint(w, "Payment retrieved")
}

// Start starts the payment service
func (s *PaymentService) Start() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST"},
	})

	handler := c.Handler(s.Router)

	go func() {
		log.Fatal(http.ListenAndServe(":8080", handler))
	}()

	fmt.Println("Payment service started on port 8080")

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	fmt.Println("Shutting down payment service")
}

func main() {
	service := NewPaymentService()
	service.Start()
}