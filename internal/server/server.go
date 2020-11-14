package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/mlan080/final-project/internal/database"
)

func Server(db *database.Database) {
	router := mux.NewRouter().StrictSlash(true)
	//router.HandleFunc("/airports", airportsGet).Methods("GET")

	router.HandleFunc("/airports", func(w http.ResponseWriter, r *http.Request) {
		airportsGetall(db, w, r)
	}).Methods("GET")

	router.HandleFunc("/airports/{iataCode}", func(w http.ResponseWriter, r *http.Request) {
		airportsIataCodeGet(db, w, r)
	}).Methods("GET")

	router.HandleFunc("/airports/{iatacode}", func(w http.ResponseWriter, r *http.Request) {
		airportsIataCodeGet(db, w, r)
	}).Methods("DELETE")

	router.HandleFunc("/airports", func(w http.ResponseWriter, r *http.Request) {
		airportsPost(db, w, r)
	}).Methods("POST")

	router.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		healthzGet(w, r)
	}).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

//endpoints
func airportsGetall(db *database.Database, w http.ResponseWriter, r *http.Request) bool {
	q := r.URL.Query().Get("q")

	if q == "maxlat" {
		tempLat := 0.0
		tempAirportIata := ""
		for _, airport := range db.GetAllAirports() {
			if tempLat < airport.Latitude {
				tempLat = airport.Latitude
				tempAirportIata = airport.IATA
			}
		}
		json.NewEncoder(w).Encode(db.GetAirport(tempAirportIata))
		return true
	} else if q == "minlat" {
		tempLat := 1000.0
		tempAirportIata := ""
		for _, airport := range db.GetAllAirports() {
			if tempLat > airport.Latitude {
				tempLat = airport.Latitude
				tempAirportIata = airport.IATA
			}
		}
		json.NewEncoder(w).Encode(db.GetAirport(tempAirportIata))
		return true
	} else if q == "maxlon" {
		tempLon := -100.0
		tempAirportIata := ""
		for _, airport := range db.GetAllAirports() {
			if tempLon < airport.Longitude {
				tempLon = airport.Longitude
				tempAirportIata = airport.IATA
			}
		}
		json.NewEncoder(w).Encode(db.GetAirport(tempAirportIata))
		return true
	} else if q == "minlon" {
		tempLon := 1000.0
		tempAirportIata := ""
		for _, airport := range db.GetAllAirports() {
			if tempLon > airport.Longitude {
				tempLon = airport.Longitude
				tempAirportIata = airport.IATA
			}
		}
		json.NewEncoder(w).Encode(db.GetAirport(tempAirportIata))
		return true
	}
	json.NewEncoder(w).Encode(db.GetAllAirports())
	return true
}

func airportsIataCodeGet(db *database.Database, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	iataCode := strings.ToUpper(vars["iataCode"])
	if len(iataCode) != 3 {
		log.Printf("Invalid IATA code for GET airportsIataCodeGet", iataCode)
		w.WriteHeader(http.StatusBadRequest)
	}
	// for _, airport := range db.GetAllAirports() {
	// 	if iataCode != db.DBAirport.IATA {
	//log.Printf("Airport not found for GET airportsIataCodeGet", iataCode)
	// 	w.WriteHeader(http.StatusPageNotFound)
	//}

	json.NewEncoder(w).Encode(db.GetAirport(iataCode))
}

func airportsIataCodeDelete(db *database.Database, w http.ResponseWriter, r *http.Request) {
	type airport struct {
		IATA      string `json:"IATA"`
		Latitude  string `json:"Latitude"`
		Longitude string `json:"Longitude"`
	}

	type allAirports []airport

	var airports = allAirports{
		{
			IATA:      "12345", //airport.IATA,
			Latitude:  "12345", //airport.Latitude,
			Longitude: "12345", //airport.Longitude,
		},
	}

	vars := mux.Vars(r)
	iataCode := strings.ToUpper(vars["iataCode"])
	for i, airport := range airports {
		if airport.IATA == iataCode {
			airports = append(airports[:i], airports[i+1:]...)
			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", iataCode)

		}
	}
}

func airportsPost(db *database.Database, w http.ResponseWriter, r *http.Request) bool {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading request body for POST airportsPost: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return true
	}
	var airport database.DBAirport
	err = json.Unmarshal(data, &airport)
	if err != nil {
		log.Printf("error parsing request for POST airportsPost: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return true
	}
	if len(airport.IATA) != 3 {
		log.Printf("Invalid IATA code for POST airportsPost", airport.IATA)
		w.WriteHeader(http.StatusBadRequest)
		return true
	}
	if _, ok := db.AirportsData[airport.IATA]; ok {
		log.Printf("IATA code already exists in the db for POST airportsPost", airport.IATA)
		w.WriteHeader(http.StatusConflict)
		return true
	}
	db.AddAirport(airport)
	w.WriteHeader(http.StatusCreated)
	return true
}

func healthzGet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
