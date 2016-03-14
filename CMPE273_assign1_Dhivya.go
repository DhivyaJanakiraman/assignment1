package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Profile struct {
	Email          string `json:"email"`
	Zip            string `json:"zip"`
	Country        string `json:"country"`
	Profession     string `json:"profession"`
	Favorite_color string `json:"favorite_color"`
	Is_smoking     string `json:"is_smoking"`
	Favorite_sport string `json:"favorite_sport"`
	Food           struct {
		Type          string `json:"type"`
		Drink_alcohol string `json:"drink_alcohol"`
	} `json:"food"`
	Music struct {
		SpotifyUserID string `json:"spotify_user_id"`
	} `json:"music"`
	Movie struct {
		Movies  []string `json:"movies"`
		TvShows []string `json:"tv_shows"`
	} `json:"movie"`
	Travel struct {
		Flight struct {
			Seat string `json:"seat"`
		} `json:"flight"`
	} `json:"travel"`
}

var profile map[string]*Profile

func HandlePost(rw http.ResponseWriter, req *http.Request) {

	users := new(Profile)
	//employee.Email = GenUUID()

	dec := json.NewDecoder(req.Body)
	err := dec.Decode(&users)
	if err != nil {
		//http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	profile[users.Email] = users

	retjs, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		//http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println(string(retjs))
	//fmt.Fprint(rw, string(retjs))
	rw.WriteHeader(201)
	return

}

func HandleProfile(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	switch req.Method {
	case "GET": //READ users
		fmt.Println("GET /profile/" + vars["email"])

		users := profile[vars["email"]]

		js, err := json.MarshalIndent(users, "", "  ")
		if err != nil {
			//http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Println(string(js))
		fmt.Fprint(rw, string(js))
		//rw.WriteHeader(http.StatusOK)
		return

	case "PUT": // UPDATE users
		fmt.Println("PUT /profile/" + vars["email"])

		users := profile[vars["email"]]

		dec := json.NewDecoder(req.Body)
		err := dec.Decode(&users)
		if err != nil {
			//http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		users.Email = vars["email"]

		retjs, err := json.Marshal(users)
		if err != nil {
			//http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Println(string(retjs))
		//fmt.Fprint(rw, string(retjs))
		rw.WriteHeader(http.StatusNoContent)
		return

	case "DELETE":
		fmt.Println("DELETE /profile/" + vars["email"])

		delete(profile, vars["email"])

		//fmt.Fprint(rw, "Success")
		rw.WriteHeader(http.StatusNoContent)
		return
	}

}

func main() {
	profile = map[string]*Profile{

		//"abcde": &Profile{Email: "abcde", Zip: "Matthew Brown", Country: "Gopher", Profession: "singer", Favorite_color: "blue", Is_smoking: "no", Favorite_sport: "badminton", Food: {Type: "veg", Drink_alcohol: "no"}},
		"xyz": &Profile{Email: "xyz", Zip: "Alexander Brown", Country: "Gopher's Assistant"},
	}

	router := mux.NewRouter()
	router.HandleFunc("/profile", HandlePost).Methods("POST")
	router.HandleFunc("/profile/{email}", HandleProfile).Methods("GET", "PUT", "DELETE")
	log.Println("listening")

	http.ListenAndServe(":3000", router)
}
