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

func New() *Database {
	var db Database
	return &db
}

//Fetches an Airport given an IATA code
func (db *Database) GetAirport(iata string) DBAirport {
	return db.AirportsData[iata]
}
