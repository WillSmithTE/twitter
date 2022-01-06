package vaccineDemographics

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

func GetIncidenceRates() (*Database, error) {
	filename := "src/germanCovid/assets/incidences_10-11-2021.csv"
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvr := csv.NewReader(f)
	csvr.Read() // skip headers row

	database := Database{}
	for {
		row, err := csvr.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return &database, err
		}

		name := row[0]
		incidenceRate := row[2]
		incidenceAsFloat, err := strconv.ParseFloat(incidenceRate, 64)
		if err != nil {
			log.Printf("Failed to convert string to float - %v", err)
			incidenceAsFloat = 0
		}
		areaData := &AreaData{Area: Area{name}, IncidenceRate: incidenceAsFloat, VaccinatedPercentage: 0}

		database.Data = append(database.Data, areaData)
	}
}
