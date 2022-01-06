package vaccineDemographics

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetGeographicalVaccData() (*Database, error) {
	filename := "src/vaccineDemographics/assets/geographicalCensusData/9-8-2021/covid-19-vaccination-geographic-vaccination-rates-9-august-2021.csv"
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

		areaData := &AreaData{}
		state := row[0]
		name4 := row[1]
		num1Dose := row[2]
		num2Doses := row[3]

		num1Float, num2Float := PctToDecimal(num1Dose), PctToDecimal(num2Doses)
		areaData.CovidVaccine = CovidVaccine{num1Float, num2Float}
		areaData.Area = *NewArea(state, name4)

		database.Data = append(database.Data, areaData)
	}
}

func PctToDecimal(pct string) float64 {
	withoutPct := strings.Trim(pct, "%")
	converted, err := strconv.ParseFloat(withoutPct, 64)
	if err != nil {
		log.Printf("Failed to convert percentage to float - %v", err)
		return 0
	} else {
		return converted
	}

}
