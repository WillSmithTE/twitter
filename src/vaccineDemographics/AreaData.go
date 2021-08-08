package vaccineDemographics

import (
	"encoding/json"
	"fmt"
)

type Database struct {
	Data     []*AreaData
	Computed ComputedData
}

func (database *Database) MarshalJSON() ([]byte, error) {
	return json.Marshal(Database{
		Data:     database.Data,
		Computed: *NewComputedData(database.Data),
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
