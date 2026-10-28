package main

import (
	"errors"
	"expenses-service/internal/budget"
	"expenses-service/internal/db/postgres"
	"log"
	"net/http"
	"time"
)

func main() {
	driver := &postgres.PostgresDriver{}
	db, err := driver.Connect()

	if err != nil {
		log.Fatalf("database initialization failed: %v", err)
	}
	defer db.Close()

	log.Println("Connected to PostgreSQL")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`OK`))
	})

	budget.RegisterBudgetModule(mux, db)

	server := &http.Server{
		Addr:              ":50053",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Printf("Server listening on %s", server.Addr)

	if err := server.ListenAndServe(); err != nil &&
		!errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
