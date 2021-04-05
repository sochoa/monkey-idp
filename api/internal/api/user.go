package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID       string
	Email        string
	PasswordHash string `json:"-"`
	Salt         string `json:"-"`
}

func (u *User) String() string {
	var userJson map[string]string = make(map[string]string)
	userJson["userid"] = u.UserID
	userJson["email"] = u.Email
	if userBytes, err := json.Marshal(userJson); err == nil {
		return string(userBytes)
	} else {
		return fmt.Sprintf("An error occured while serializing user to JSON: %v", err)
	}
}

var (
	createUserRequiredArguments []string = []string{
		"userid", "{userid}",
		"email", "{email}",
		"password", "{password}",
	}
	authUserRequiredArguments []string = []string{
		"userid", "{userid}",
		"password", "{password}",
	}
)

func createUserhandler(w http.ResponseWriter, r *http.Request) {
	log, _ := zap.NewProduction()
	defer log.Sync()

	passwordBytes := []byte(r.PostFormValue("password"))
	hash, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)
	if err != nil {
		msg := fmt.Sprintf("An error occurred while hashing the password: %v", err)
		log.Error(msg)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(msg))
		return
	}
	user := &User{
		UserID:       r.PostFormValue("userid"),
		Email:        r.PostFormValue("email"),
		PasswordHash: string(hash),
	}

	db := getDatabaseConnection()
	defer db.Close()
	if _, err := db.Model(user).Insert(); err != nil {
		log.Error(fmt.Sprintf("An error occurred while inserting a user: user=%s.  Details:  %v", user.String(), err))
	} else {
		log.Info("Successfully created user", zap.String("user", user.String()))
	}
}

func authUserhandler(w http.ResponseWriter, r *http.Request) {
	log, _ := zap.NewProduction()
	defer log.Sync()
	user := User{UserID: r.PostFormValue("userid")}
	db := getDatabaseConnection()
	defer db.Close()
	err := db.Model(&user).Select()
	if err != nil {
		log.Info("Does not exist", zap.String("user", user.UserID))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	passwordBytes := []byte(r.PostFormValue("password"))
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), passwordBytes); err != nil {
		log.Info("Bad Password", zap.String("user", user.UserID))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	log.Info("Successful Authentication", zap.String("user", user.UserID))
	w.WriteHeader(http.StatusOK)
}

func listUsersHandler(w http.ResponseWriter, r *http.Request) {
	log, _ := zap.NewProduction()
	defer log.Sync()
	db := getDatabaseConnection()
	defer db.Close()
	var users []User

	if err := db.Model(&users).Select(); err != nil {
		log.Info(fmt.Sprintf("Failed to list users from database: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if userBytes, err := json.Marshal(users); err == nil {
		log.Info(fmt.Sprintf("Found %d users", len(users)))
		w.Write(userBytes)
	} else {
		log.Info(fmt.Sprintf("Failed to serialize the list of users from database: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func addUserRoutes(router *mux.Router) {
	log, _ := zap.NewProduction()
	defer log.Sync()

	log.Info("Adding /api/v1/user POST handler for creating users")
	router.HandleFunc("/user", createUserhandler).Methods("POST").Queries(createUserRequiredArguments...)

	log.Info("Adding /api/v1/authenticate POST handler for authenticating users")
	router.HandleFunc("/authenticate", authUserhandler).Methods("POST")

	log.Info("Adding /api/v1/users GET handler for listing users")
	router.HandleFunc("/users", listUsersHandler).Methods("GET")
}
