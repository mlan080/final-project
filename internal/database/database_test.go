package database

import (
	"testing"
)

func TestData(t *testing.T) {
	tests := []struct {
		IATA      string 
		arg map[string]AirportsData
		want []*db.DBAirport
	}{
		name: "POM", 
		arg: map[string]AirportsData{
	"POM": {IATA: "POM", Latitude: -9.44338036, Longitude: 147.22000122, Name: "Port Moresby Jacksons International Airport", Type: "large_airport"},
	"KEF": {IATA: "KEF", Latitude: 63.98500061, Longitude: -22.60560036, Name: "Keflavik International Airport", Type: "large_airport"},
	"PRN": {IATA: "PRN", Latitude: 42.57280000, Longitude: 21.03580100, Name: "Priština International Airport", Type: "large_airport"},
	"YEG": {IATA: "YEG", Latitude: 53.30970001, Longitude: -113.58000183, Name: "Edmonton International Airport", Type: "large_airport"},
		}, 
		want []*db.DBAirport{
			{
				IATA: 	

			}
		}

	}

}
