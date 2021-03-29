package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           string   `json:"id"`
	FirstName    string   `json:"first_name"`
	LastName     string   `json:"last_name"`
	Emails       []string `json:"emails"`
	PasswordHash string   `json:"-"`
	Salt         string   `json:"-"`
}

func init() {
	db := getDatabaseConnection()
	defer db.Close()
	createSchema(db)
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*User)(nil),
	}
	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *User) String() string {
	if userBytes, err := json.Marshal(u); err == nil {
		return string(userBytes)
	} else {
		return fmt.Sprintf("An error occured while serializing user to JSON: %v", err)
	}
}

var (
	createUserRequiredArguments []string = []string{
		"id", "{id}",
		"first", "{first}",
		"last", "{last}",
		"email", "{email}",
		"password", "{password}",
	}
	authUserRequiredArguments []string = []string{
		"id", "{id}",
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
	user := User{
		ID:           r.PostFormValue("id"),
		FirstName:    r.PostFormValue("first"),
		LastName:     r.PostFormValue("last"),
		Emails:       []string{r.PostFormValue("email")},
		PasswordHash: string(hash),
	}

	db := getDatabaseConnection()
	defer db.Close()
	db.Model(user).Insert()
	log.Info("Created user", zap.String("user", user.String()))
}

func authUserhandler(w http.ResponseWriter, r *http.Request) {
	log, _ := zap.NewProduction()
	defer log.Sync()
	user := User{ID: r.PostFormValue("id")}
	db := getDatabaseConnection()
	defer db.Close()
	err := db.Model(&user).Select()
	if err != nil {
		log.Info("Does not exist", zap.String("user", user.ID))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	passwordBytes := []byte(r.PostFormValue("password"))
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), passwordBytes); err != nil {
		log.Info("Bad Password", zap.String("user", user.ID))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	log.Info("Successful Authentication", zap.String("user", user.ID))
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
	subRouter := router.PathPrefix("/user").Subrouter()
	subRouter.HandleFunc("/new", createUserhandler).Methods("POST").Queries(createUserRequiredArguments...)
	subRouter.HandleFunc("/auth", authUserhandler).Methods("POST")
	router.PathPrefix("/users").Subrouter().HandleFunc("", listUsersHandler).Methods("GET")
}
