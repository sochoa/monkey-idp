package api

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/antonlindstrom/pgstore"
	"github.com/gorilla/sessions"
	"go.uber.org/zap"
	"net/http"
	"os"
	"time"
)

var (
	sessionKey       []byte
	connectionString string
	cookieStore      *pgstore.PGStore
)

func init() {
	log, _ := zap.NewProduction()
	defer log.Sync()

	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	database := os.Getenv("POSTGRES_DB")

	host := os.Getenv("POSTGRES_HOST")
	if len(host) == 0 {
		host = "127.0.0.1"
	}

	port := os.Getenv("POSTGRES_PORT")
	if len(port) == 0 {
		port = "5432"
	}

	connectionString = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, database)
	redactedConnectionString := fmt.Sprintf("postgres://%s:<password>@%s:%s/%s", user, host, port, database)
	sessionKey = []byte(os.Getenv("SESSION_KEY"))
	if len(sessionKey) == 0 {
		log.Warn("Using randomly generated session key")
		sessionKey = generateSessionKey()
	}
	log.Info(fmt.Sprintf("Using postgres for session storage: %s", redactedConnectionString))
}

func generateSessionKey() []byte {
	size := 32
	randomBytes := make([]byte, size)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	return []byte(base64.URLEncoding.EncodeToString(randomBytes))
}

func getSession(r *http.Request, user User) {
	var (
		err     error = nil
		session *sessions.Session
	)
	log, _ := zap.NewProduction()
	defer log.Sync()
	cookieStore, err = pgstore.NewPGStore(connectionString, sessionKey)
	if err != nil {
		panic(err)
	}
	defer cookieStore.Close()
	defer cookieStore.StopCleanup(cookieStore.Cleanup(time.Minute * 5))

	session, err = cookieStore.Get(r)
	if err != nil {
		log.Error(fmt.Sprintf("An error occurred while creating a session: %v", err))
	}

	// Serialize user to JSON
	// Add signature to JSON
}
