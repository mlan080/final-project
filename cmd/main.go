package main

import "github.com/fromatob/engineering/training/backend/6/server"

func main() {
	server.Server()
}

// set up HTTP handlers

//find airport by IATA code
// func airportsIataCodeGet(w http.ResponseWriter, req *http.Request) {
// 	iata := req.URL.Query().Get("iata")
// 	//iata invalid -
// 	if iata == "" {
// 		http.Error(w, "iatacode is required", http.StatusBadRequest)
// 		return
// 	}
// 	airports := Search(iataCode)

// 	bytes, err := json.Marshall(airports)
// 	if err != nil {
// 		http.Error(w, "invalid IATA code", http.StatusBadRequest)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	_, err = w.Write(bytes)
// 	if err != nil {
// 		http.Error(w, "iataCode not found in the database", http.StatusNotFound)
// 		return
// 	}
// }

// func Search(prefix string) []Airport {
// 	prefix = strings.ToLower(prefix)
// 	result := []Airport{}
// 	for _, city := range airports {
// 		if strings.HasPrefix(strings.ToLower(city.Name), prefix) {
// 			result = append(result, airport)
// 		}
// 	}
// 	return result
// }
