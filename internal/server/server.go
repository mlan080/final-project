package server

import (
	"encoding/json"
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

		//router.HandleFunc("/airports/{iata}", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		// case "GET":
		// 	airportsIataCodeGet(db, w, r)
		case "GET":
			airportsGet(db, w, r)
			// // }
			// case "DELETE":
			// 	airportsIataCodeDelete(db, w, r)
			// 	case "POST":
			// 		airportsPost(w, req)
			// 	case "GET":
			// 		healthzGet(w, req)
			// 	default:
			// 		http.Error(w, "allowed methods for /airports: POST", http.StatusMethodNotAllowed)
		}
	}).Queries("maxLat", "{[0-9]*?}")
	log.Fatal(http.ListenAndServe(":8080", router))
}

//endpoints
func airportsIataCodeGet(db *database.Database, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	iataCode := strings.ToUpper(vars["iataCode"])
	json.NewEncoder(w).Encode(db.GetAirport(iataCode))
}

func airportsGet(db *database.Database, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	q := r.URL.Query().Get("q")
	var values map[string]DBAirport
	maxLat := values.Latitude[0]
	for _, airport := range db.AirportsData {
		if airport.Latitude > maxLat.Latitude {
			maxLat = airport
			continue
		}
		if q == "maxLat" {
			json.NewEncoder(w).Encode(airport)
		}
	}
	json.NewEncoder(w).Encode(db.AirportsData)
}

// func airportsIataCodeDelete(db *database.Database, w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	iataCode := strings.ToUpper(vars["iataCode"])
// 	airporttodelete := db.GetAirport(iataCode)
// 	airporttodelete.Longitude
// 	airporttodelete = append(airporttodelete[:1], airporttodelete[i+1:]...)
// 	fmt.Fprintf(w, "The airport IATA %s has been deleted successfully", iataCode)
// }

// func airportsPost(w http.ResponseWriter, r *http.Request) {
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
// }

// func healthzGet() {
// 	healthcheck.WithChecker(
// 		"project6", checkers.Heartbeat("/6"),
// 	)
// }

//extra code

// response := struct {
// 	iataCode  string
// 	latitude  float64
// 	name      string
// 	longitude float64
// }

// bytes, err := json.Marshal(response)
// if err != nil {
// 	log.Printf("error marshaling JSON: %v", err)
// 	http.Error(w, "internal error", http.StatusInternalServerError)
// 	return
// }
// minLat, maxLat, minLon, maxLon := airports[0], airports[0], airports[0], airports[0]

// for _, airport := range airports {
// 	if airport.Latitude > maxLat.Latitude {
// 		maxLat = airport
// 	}
// 	if airport.Latitude < minLat.Latitude {
// 		minLat = airport
// 	}
// 	if airport.Longitude < maxLon.Longitude {
// 		maxLon = airport
// 	}
// 	if airport.Longitude > minLon.Longitude {
// 		minLon = airport
// 	}
// 	return minLat, maxLat, minLon, maxLon

// }
//}

//define the minLat by iteratig through
//if q == minlat, return the slice
//if q == maxlat, return the slice

// for _, airport := range db.AirportsData {
// 	if strings.HasPrefix(strings.ToLower(airport.IATA), strings.ToLower(iataCode)) {
// 		db.AirportsData = append(db.AirportsData, airport)
// 		continue
// 	}

// for i, airport := range db.AirportData {
// 	if airport.IATA == iataCode {
// 		airports = append(airports[:1], airports[i+1:]...)
// 		fmt.Fprintf(w, "The airport IATA %s has been deleted successfully", iataCode)
// 		continue

// 	if airport.IATA == iataCode {
// 		//json.NewEncoder(w).Encode(airport)
// 		fmt.Println(airport)
