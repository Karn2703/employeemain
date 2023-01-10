package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"sync"
)

// Service holds the map of items and provides methods CRUD operations on the map
type Service struct {
	connectionString string
	profiles           map[string]Profile
	sync.RWMutex
}

// NewService returns a Service with a connectionString configured and can be a map of items setup. The items map can be empty,
// or can contain items
func NewService(connectionString string, profiles map[string]Profile) *Service {
	return &Service{
		connectionString: connectionString,
		profiles:            profiles,
	}
}

func main(){
	router := mux.NewRouter()

	router.HandleFunc("/profiles", additem).Methods("POST")

	router.HandleFunc("/profiles", getAllProfiles).Methods("GET")

	router.HandleFunc("/profiles/{id}", getProfile).Methods("GET")

	router.HandleFunc("/profiles/{id}", updateProfile).Methods("PUT")

	router.HandleFunc("/profiles/{id}", deleteProfile).Methods("DELETE")

	http.ListenAndServe(":5000", router)

}