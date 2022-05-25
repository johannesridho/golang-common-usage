package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	ID   int
	Name string
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/users", GetUsersHandler).Methods(http.MethodGet)
	router.HandleFunc("/api/users/{id}", GetUserHandler).Methods(http.MethodGet)
	router.HandleFunc("/api/users", CreateUserHandler).Methods(http.MethodPost)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	// get the query parameters
	fmt.Printf("%+v\n", r.URL.Query())

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	users := &[]User{
		{
			ID:   1,
			Name: "user1",
		}, {
			ID:   2,
			Name: "user2",
		},
	}

	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// get the id param from /api/users/{id}
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	user := &User{
		ID:   id,
		Name: fmt.Sprint("user", id),
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	user := &User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("%+v", user)
	w.WriteHeader(http.StatusCreated)
}
