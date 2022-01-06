package vaccineDemographics

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

func AddVaccinationData(database *Database) error {
	filename := "src/germanCovid/assets/vaccines_10-11-2021.csv"
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	csvr := csv.NewReader(f)
	csvr.Read() // skip headers row

	for {
		row, err := csvr.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return nil
		}

		name := row[0]
		areaData := findAreaData(database, name)
		if areaData == nil {
			log.Printf("Failed to find %v in database (adding vaccine data)", name)
		} else {
			vaccinationRate := row[1]
			asFloat, err := strconv.ParseFloat(vaccinationRate, 64)
			if err != nil {
				log.Printf("Failed to convert string to float - %v", err)
				asFloat = 0
			}
			areaData.VaccinatedPercentage = asFloat
		}
	}
}

func findAreaData(db *Database, name string) *AreaData {
	for _, areaData := range db.Data {
		if areaData.Area.Name == name {
			return areaData
		}
	}
	return nil
}
