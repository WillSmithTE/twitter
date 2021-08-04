package vaccineDemographics

import (
	"log"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func AddCensusCodes(database *Database) error {
	found := 0
	filePath := "src/vaccineDemographics/assets/2016_GCP_ALL_for_NSW_short-header/Metadata/2016Census_geog_desc_1st_2nd_3rd_release.xlsx"
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return err
	}
	sheetName := f.WorkBook.Sheets.Sheet[0].Name
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return err
	}
	for _, row := range rows {
		if found == len(database.Data) {
			return nil
		}
		name := row[3]
		for _, areaData := range database.Data {
			if areaData.Area.Name4 == name {
				found += 1
				code, err := strconv.Atoi(row[2])
				if err != nil {
					log.Printf("error converting census code to int for %v - %v", name, err)
				} else {
					areaData.Area.CensusCode = code
				}
			}
		}
	}
	return nil
}
