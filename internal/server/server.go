package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mlan080/final-project/internal/database"
)

func Server(db *database.Database) {
	router := mux.NewRouter().StrictSlash(true)
	//router.HandleFunc("/airports", airportsGet).Methods("GET")
	router.HandleFunc("/airports", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			airportsGet(db, w, r)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", router))
}

// http.HandleFunc("/airports", func(w http.ResponseWriter, req *http.Request) {
// 	switch req.Method {
// 	case "GET":
// 		airportsGet(w, req)
// 	case "DELETE":
// 		airportsIataCodeDelete(w, req)
// 	case "GET":
// 		airportsIataCodeGet(w, req)
// 	case "POST":
// 		airportsPost(w, req)
// 	case "GET":
// 		healthzGet(w, req)
// 	default:
// 		http.Error(w, "allowed methods for /airports: POST", http.StatusMethodNotAllowed)
// 	}

// }

//endpoints
func airportsGet(db *database.Database, w http.ResponseWriter, r *http.Request) {
	//json.NewEncoder(w).Encode(airports)
	fmt.Println(db.GetAirport("POM"))
	//fmt.Println("hello")
	// 	//GetAirport(iataCode)
	//q := r.URL.Query().Get("q")

}

//q := r.URL.Query().Get("q")
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

// func getAllAirports(w http.ResponseWriter, r *http.Request) {
// 	json.NewEncoder(w).Encode(airports)
// 	// if err != nil {
// 	// 	http.Error(w, "Invalid query parameters", http.StatusBadRequest)
// 	//}
// }

// func getOneAirport(w http.ResponseWriter, r *http.Request) {

// 	//The names are used to create a map of route variables which can be retrieved calling mux.Vars()
//how do I get it to print a response for iatacode pom?
// }

// func deleteAirport(w http.ResponseWriter, r *http.Request) {
// 	iataCode := mux.Vars(r)["iataCode"]

// 	for i, oneAirport := range airports {
// 		if oneAirport.IATA == iataCode {
// 			airports = append(airports[:1], airports[i+1:]...)
// 			fmt.Fprintf(w, "The airport IATA %s has been deleted successfully", iataCode)
// 			continue
// 		}
// 		if strings.HasPrefix(strings.ToLower(oneAirport.IATA), strings.ToLower(iataCode)) {
// 			airports = append(airports, oneAirport)
// 			continue
// 		}
// 	}
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
