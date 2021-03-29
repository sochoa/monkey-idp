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
	log, _ := zap.NewProduction()
	defer log.Sync()

	r := mux.NewRouter()
	addUserRoutes(r)

	endpoint := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	log.Info(fmt.Sprintf("Running API server on %s", endpoint))
	srv := &http.Server{
		Handler:      r,
		Addr:         endpoint,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	return srv
}
