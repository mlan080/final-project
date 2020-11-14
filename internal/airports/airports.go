package airports

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/mlan080/final-project/internal/database"
)

type airport struct {
	IATA      string
	Latitude  float64
	Longitude float64
	Name      string
	Type      string
}

func ParseCSV(db *database.Database) {
	csvfile, err := os.Open("airport-data.csv")
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
				log.Printf("skipping, wrong IATA %s", iata)
			}
			continue
		}

		if name == "" {
			log.Printf("skipping, %s dosen't have name", iata)
		}

		name = strings.ReplaceAll(name, `"`, `\"`)
		if _, ok := airportsData[iata]; ok {
			log.Printf("skipping %s, duplicated", iata)
			continue
		}

		latitude, err := strconv.ParseFloat(lat, 64)
		if err != nil {
			log.Println(err)
			continue
		}
		longitude, err := strconv.ParseFloat(lon, 64)
		if err != nil {
			log.Println(err)
			continue
		}
		if typ != "large_airport" {
			log.Printf("skipping, %s is not the correct type", iata)
			continue
		}
		airportsData[iata] = database.DBAirport{Name: name, IATA: iata, Latitude: latitude, Longitude: longitude, Type: typ}
	}
	//populate the datastore - fill up empty db
	db.AirportsData = airportsData
}

func Err(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
