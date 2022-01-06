package vaccineDemographics

import (
	"encoding/json"
	"fmt"
)

type Database struct {
	DataByDate     map[string][]*AreaData
	ComputedByDate map[string]ComputedData
}

func (database *Database) MarshalJSON() ([]byte, error) {
	return json.Marshal(Database{
		DataByDate:     database.DataByDate,
		ComputedByDate: *NewComputedData(database.DataByDate),
	})
}

func (database *Database) getAreaDataByCode(code int) (*AreaData, error) {
	for _, areaData := range database.Data {
		if areaData.Area.CensusCode == code {
			return areaData, nil
		}
	}
	return nil, fmt.Errorf("failed to find area data of code %v in database", code)
}

type AreaData struct {
	Area         Area
	CovidVaccine CovidVaccine
}
