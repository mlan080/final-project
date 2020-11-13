package server

import (
	"encoding/json"
	"fmt"
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
	}).Methods("GET").Queries("maxLat", "{[0-9]*?}")

	router.HandleFunc("/airports/{iataCode}", func(w http.ResponseWriter, r *http.Request) {
		airportsIataCodeGet(db, w, r)
	}).Methods("GET")

	router.HandleFunc("/airports/{iatacode}", func(w http.ResponseWriter, r *http.Request) {
		airportsIataCodeGet(db, w, r)
	}).Methods("DELETE")

	router.HandleFunc("/airports", func(w http.ResponseWriter, r *http.Request) {
		airportsPost(w, r)
	}).Methods("POST")

	router.HandleFunc("/airports", func(w http.ResponseWriter, r *http.Request) {
		healthzGet()
	}).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

//endpoints
func airportsGetall(db *database.Database, w http.ResponseWriter, r *http.Request) {
	//db = database.New()
	//q := r.URL.Query().Get("q")
	// if q != "" {
	// 	airportsGetQueries(db, w, r)
	// }
	json.NewEncoder(w).Encode(db.AirportsData)
}

//convert map into an array and response
func airportsGetQueries(db *database.Database, w http.ResponseWriter, r *http.Request) {
	// 	// 	w.Header().Set("Content-Type", "application/json")
	// 	q := r.URL.Query().Get("q")

	// 	// 	//convert map to slice
	// 	// response := struct {
	// 	// 	Latitude float64 `json:"latitude"`
	// 	// }{
	// 	// 	Latitude: database.DBAirport.float64(Latitude),
	// 	// }
	// 	//convert map to slice
	// 	airports := []float64{}
	// 	airportsData := make(map[string]database.DBAirport)

	// 	if q == "maxLat" {
	// 		tempLat := 0.0
	// 		tempAirport := ""

	// 		for airport := range db.AirportsData {
	// 			if tempLat < dblatitude {
	// 				tempLat = append(airport.Latitude
	// 				fmt.Println(tempLat)

	// 				pairs := [][]string{}
	// 				for key, value := range m {
	// 					pairs = append(pairs, []string{key, value})
	// 				}
	// 			}
	// 		}
	// 	}
}

func airportsIataCodeGet(db *database.Database, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	iataCode := strings.ToUpper(vars["iataCode"])
	json.NewEncoder(w).Encode(db.GetAirport(iataCode))
}

//json.NewEncoder(w).Encode(db.AirportsData)

func airportsIataCodeDelete(db *database.Database, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	iataCode := strings.ToUpper(vars["iataCode"])
	airporttodelete := db.GetAirport(iataCode)
	fmt.Println(airporttodelete)
}

//in built delete function for maps -

func airportsPost(w http.ResponseWriter, r *http.Request) {
	// 	data, err := ioutil.ReadAll(r.Body)
	// 	if err != nil {
	// 		log.Printf("error reading request body: %v", err)
	// 		return
	// 	}

	// 	var airport Airport
	// 	err = json.Unmarshal(data, &airport)
	// 	if err != nil {
	// 		log.Printf("error parsing request: %v", err)
	// 		http.Error(w, "parameter 'q' is required", http.StatusBadRequest)
	// 		return
	// 	}

	// 	GetAirport(airport.IATA)
}

//if ok = checks if the airport already exists

func healthzGet() {
	// 	healthcheck.WithChecker(
	// 		"project6", checkers.Heartbeat("/6"),
	// 	)
}
