package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func getRequestAsMap(r *http.Request) (map[string]interface{}, error) {

	var (
		parsedRequest    map[string]interface{} = make(map[string]interface{})
		requestBodyJson  map[string]interface{} = make(map[string]interface{})
		requestBodyBytes []byte                 = make([]byte, 0)
		contentType      string                 = r.Header.Get("Content-Type")
		err              error
	)

	log, _ := zap.NewProduction()
	defer log.Sync()

	if r.Body != nil {
		requestBodyBytes, _ = ioutil.ReadAll(r.Body)
	}
	r.Body = ioutil.NopCloser(bytes.NewBuffer(requestBodyBytes))

	if contentType == "application/json" {
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&requestBodyJson); err == nil {
			for arg, val := range requestBodyJson {
				parsedRequest[arg] = val
				log.Info(fmt.Sprintf("JSON request body:  %v = %v", arg, val)) /* TODO:  remove password log */
			}
		}
	} else if err = json.Unmarshal(requestBodyBytes, &requestBodyJson); err == nil {
		for arg, val := range requestBodyJson {
			log.Info(fmt.Sprintf("raw request body as JSON:  %v = %v", arg, val)) /* TODO:  remove password log */
			parsedRequest[arg] = val
		}
	} else {
		log.Error(fmt.Sprintf("An error occurred while unmarshalling the request to map[string]interface{}: %v", err))
	}

	log.Info(fmt.Sprintf("Request:  %v", parsedRequest)) /* TODO:  remove password log */
	return parsedRequest, nil
}

func authUserhandler(w http.ResponseWriter, r *http.Request) {
	var (
		requestArgs map[string]interface{} = nil
		password    string                 = ""
		err         error                  = nil
		user        User
	)

	log, _ := zap.NewProduction()
	defer log.Sync()

	if requestArgs, err = getRequestAsMap(r); err != nil {
		log.Error(fmt.Sprintf("An error occurred while parsing the request arguments: %v", err))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if userId, ok := requestArgs["userid"]; ok {
		user = User{UserID: userId.(string)}
		log.Info(fmt.Sprintf("Parsed UserID:  %s", user.UserID))
		db := getDatabaseConnection()
		defer db.Close()
		err := db.Model(&user).Select()
		if err != nil {
			log.Info("Does not exist", zap.String("user", user.UserID))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	} else {
		log.Info(fmt.Sprintf("Could not find the userId in the request: %v", requestArgs))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if p, ok := requestArgs["password"]; ok {
		password = p.(string)
		log.Info(fmt.Sprintf("Parsed Password:  %v", password)) /* TODO:  remove password log */
	} else {
		log.Info(fmt.Sprintf("Could not find the password in the request: %v", requestArgs))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	passwordBytes := []byte(password)
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), passwordBytes); err != nil {
		log.Info("Bad Password", zap.String("user", user.UserID), zap.String("password", password)) /* TODO:  remove password log */
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
	router.HandleFunc("/auth", authUserhandler).Methods("POST")

	log.Info("Adding /api/v1/users GET handler for listing users")
	router.HandleFunc("/users", listUsersHandler).Methods("GET")
}
