package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"os"
	"time"
)

func CreateServer() *http.Server {
	var (
		host string
		port string
	)
	log, _ := zap.NewProduction()
	defer log.Sync()

	r := mux.NewRouter().PathPrefix("/api/v1").Subrouter()
	addUserRoutes(r)
	addHealthRoutes(r)

	host = "0.0.0.0"
	if h, ok := os.LookupEnv("HOST"); ok {
		host = h
	}

	port = "4000"
	if p, ok := os.LookupEnv("PORT"); ok {
		port = p
	}

	bind_address := fmt.Sprintf("%s:%s", host, port)
	log.Info(fmt.Sprintf("Running API server on %s", bind_address))
	srv := &http.Server{
		Handler:      r,
		Addr:         bind_address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	return srv
}
