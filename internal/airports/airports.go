package airports

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/mlan080/final-project/internal/database"
)

func ParseCSV(db *database.Database) {
	csvfile, err := os.Open("cmd/airport-data.csv")
	Err(err)
	defer csvfile.Close()

	records := csv.NewReader(csvfile)
	_, err = records.Read()
	Err(err)

	airportsData := make(map[string]database.DBAirport)
	for {
		record, err := records.Read()
		if err == io.EOF {
			break
		}
		Err(err)

		iata, lat, lon, name, typ := record[13], record[4], record[5], record[3], record[2]
		if len(iata) != 3 {
			if iata != "" && iata != "0" && iata != "-" {
				fmt.Errorf("skipping, wrong IATA %s", iata)
			}
			continue
		}
		if name == "" {
			fmt.Errorf("skipping, %s dosen't have name", iata)
			continue
		}
		name = strings.ReplaceAll(name, `"`, `\"`)
		if _, ok := airportsData[iata]; ok {
			fmt.Errorf("skipping %s, duplicated", iata)
			continue
		}
		latitude, err := strconv.ParseFloat(lat, 64)
		if err != nil {
			continue
		}
		longitude, err := strconv.ParseFloat(lon, 64)
		if err != nil {
			continue
		}
		if typ == "large_airport" {
			fmt.Errorf("skipping, %s is not the correct type", iata)
			continue
		}
		airportsData[iata] = database.DBAirport{Name: name, IATA: iata, Latitude: latitude, Longitude: longitude, Type: typ}
	}
	db.AirportsData = airportsData
}

func Err(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
