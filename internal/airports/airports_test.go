package airports

import (
	"testing"
)

func TestData(t *testing.T) {
	for iata, data := range airportsData {
		if iata != data.IATA {
			t.Errorf("expected data.IATA code to be: %s, got:%s", iata, data.IATA)
		}
		if data.Longitude == 0 {
			t.Errorf("expected non-0 longitude for %s", iata)
		}
		if data.Latitude == 0 {
			t.Errorf("expected non-0 latitude for %s", iata)
		}
		if data.Name == "" {
			t.Errorf("expected non-empty name for %s", iata)
		}
		if data != ParseCSV(iata) {
			t.Errorf("got different data from GetAirport for %s", iata)
		}
	}

	// if a := GetAirport("INVALID"); a != nil {
	// 	t.Errorf("expected 'INVALID' IATA code to return nil, got a result")
	// }
}
