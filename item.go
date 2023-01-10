package server

import (
	"encoding/json"
	"net/http"

	"strconv"

	"github.com/gorilla/mux"
)

var profiles []Profile = []Profile{}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Profile struct {
	Employee User `json:"employee"`
}
func additem(q http.ResponseWriter, r *http.Request) {

	var newProfile Profile
	json.NewDecoder(r.Body).Decode(&newProfile)

	q.Header().Set("Content-Type", "application/json")

	profiles = append(profiles, newProfile)
	json.NewEncoder(q).Encode(profiles)
}

func getAllProfiles(q http.ResponseWriter, r *http.Request) {
	q.Header().Set("Content-Type", "application/json")
	json.NewEncoder(q).Encode(profiles)
}

func getProfile(q http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		q.WriteHeader(400)
		q.Write([]byte("ID could not be converted to int"))
		return
	}

	if id >= len(profiles) {
		q.WriteHeader(404)
		q.Write([]byte("No profile found with specify ID"))
	}

	profile := profiles[id]
	q.Header().Set("Content-Type", "application/json")
	json.NewEncoder(q).Encode(profile)

}

func updateProfile(q http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		q.WriteHeader(400)
		q.Write([]byte("ID could not be converted to int"))
		return
	}

	if id >= len(profiles) {
		q.WriteHeader(404)
		q.Write([]byte("No profile found with specify ID"))
	}

	var updateProfile Profile
	json.NewDecoder(r.Body).Decode(&updateProfile)

	profiles[id] = updateProfile

	q.Header().Set("Content-Type", "appliation/json")
	json.NewEncoder(q).Encode(updateProfile)
}

func deleteProfile(q http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		q.WriteHeader(400)
		q.Write([]byte("ID could not be converted to int"))
		return
	}

	if id >= len(profiles) {
		q.WriteHeader(404)
		q.Write([]byte("No profile found with specify ID"))
	}

	profiles = append(profiles[:id], profiles[:id+1]...)
	q.WriteHeader(200)
}
