package database

type DBAirport struct {
	IATA      string  `json:"iata"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Name      string  `json:"name"`
	Type      string  `json:"type"`
}

type Database struct {
	AirportsData map[string]DBAirport
	//handler is validation
}

//called from main or airport
func New() *Database {
	var db Database
	return &db
}

//convert my csv

// // GetAirport fetches an Airport given an IATA code - how does this help my rest api?
func (db *Database) GetAirport(iata string) DBAirport {
	return db.AirportsData[iata] //returns the airport struct

	// if !ok {
	// 	return nil
	// }

}

//iataCode := mux.Vars(r)["iataCode"]

// for _, airport := range airportsData {
// 	if airport.IATA == iataCode {
// 		json.NewEncoder(w).Encode(airport)
// 		continue
// 	}

// 	if strings.HasPrefix(strings.ToLower(airport.IATA), strings.ToLower(iataCode)) {
// 		airports = append(airports, airport)
// 		continue
// 	}
