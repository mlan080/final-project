package database

import (
	"encoding/json"
	"strings"
)

type Airport struct {
	IATA      string  `json:"iata"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Name      string  `json:"name"`
	Type      string  `json:"type"`
}

type Database struct {
	airportsData map[string]Airport
	//handler is validation
}

//called from main or airport
func New() *Database {
	var db Database
	return &db
}

//convert my csv

// // GetAirport fetches an Airport given an IATA code - how does this help my rest api?
func (db *Database) GetAirport(iata string) *Airport {
	airportsData := make(map[string]Airport)
	//iataCode := mux.Vars(r)["iataCode"]

	for _, airport := range airportsData {
		if airport.IATA == iataCode {
			json.NewEncoder(w).Encode(airport)
			continue
		}

		if strings.HasPrefix(strings.ToLower(airport.IATA), strings.ToLower(iataCode)) {
			airports = append(airports, airport)
			continue
		}
	}
	// data, ok := airportsData[iata]
	// if !ok {
	// 	return nil
	// }
	// return &data
}
