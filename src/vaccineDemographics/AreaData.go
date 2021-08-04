package vaccineDemographics

import (
	"fmt"
)

type Database struct {
	Data []*AreaData
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
