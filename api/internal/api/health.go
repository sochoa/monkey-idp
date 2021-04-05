package api

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

func addHealthRoutes(router *mux.Router) {
	log, _ := zap.NewProduction()
	defer log.Sync()
	log.Info("Adding /healthcheck GET handler for checking API health")
	router.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("OK\n"))
	}).Methods("GET")
}
