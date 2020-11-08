package database

type Database struct {
	GetAirport(iata string) Airport  
	DeleteAirport(iata string) Airport
	airportsData map[string]Airport
	//handler is validation 
}
//called from main
 func New() *Database {
// 	return &Database
// }


//convert my csv
//

// // GetAirport fetches an Airport given an IATA code - how does this help my rest api?
func (db *Database) GetAirport(iata string) Airport {
	airportsData := make(map[string]Airport)

	data, ok := airportsData[iata]
	if !ok {
		return nil
	}
	return (&data)
}
