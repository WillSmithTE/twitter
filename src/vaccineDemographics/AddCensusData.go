package vaccineDemographics

import (
	"log"
	"strconv"

	"github.com/willsmithte/twitter/src/util"
)

var sheetsDirPath = "src/vaccineDemographics/assets/2016_GCP_ALL_for_NSW_short-header/2016 Census GCP All Geographies for NSW/SA4/NSW/"

func AddCensusData(database *Database) error {
	err := addPopulation(database)
	if err != nil {
		return err
	}
	err = addMedians(database)
	if err != nil {
		return err
	}
	return nil
}

func addPopulation(database *Database) error {
	filename := sheetsDirPath + "2016Census_G01_NSW_SA4.csv"
	err := util.ExecuteOnEachCsvRow(filename, func(row []string) error {
		censusCode := row[0]
		codeInt, err := strconv.Atoi(censusCode)
		if err != nil {
			return err
		}
		areaData, err := database.getAreaDataByCode(codeInt)
		if err != nil {
			log.Print(err.Error())
			return nil
		}
		addPopAndAgesFromRow(row, areaData)
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func addPopAndAgesFromRow(row []string, areaData *AreaData) {
	areaData.Area.CensusStats.Population = stringToInt(row[3])
	areaData.Area.CensusStats.Age.Num0to4 = stringToInt(row[6])
	areaData.Area.CensusStats.Age.Num5to14 = stringToInt(row[9])
	areaData.Area.CensusStats.Age.Num15to19 = stringToInt(row[12])
	areaData.Area.CensusStats.Age.Num20to24 = stringToInt(row[15])
	areaData.Area.CensusStats.Age.Num25to34 = stringToInt(row[18])
	areaData.Area.CensusStats.Age.Num35to44 = stringToInt(row[21])
	areaData.Area.CensusStats.Age.Num45to54 = stringToInt(row[24])
	areaData.Area.CensusStats.Age.Num55to64 = stringToInt(row[27])
	areaData.Area.CensusStats.Age.Num65to74 = stringToInt(row[30])
	areaData.Area.CensusStats.Age.Num75to84 = stringToInt(row[33])
	areaData.Area.CensusStats.Age.Num85Plus = stringToInt(row[36])

}

func addMedians(database *Database) error {
	filename := sheetsDirPath + "2016Census_G02_NSW_SA4.csv"
	return util.ExecuteOnEachCsvRow(filename, func(row []string) error {
		censusCode := row[0]
		codeInt, err := strconv.Atoi(censusCode)
		if err != nil {
			return err
		}
		areaData, err := database.getAreaDataByCode(codeInt)
		if err != nil {
			log.Print(err.Error())
			return nil
		}
		areaData.Area.CensusStats.Age.Median = stringToInt(row[1])
		areaData.Area.CensusStats.Income.MedianPersonal = stringToFloat(row[3])
		areaData.Area.CensusStats.Income.MedianFamily = stringToFloat(row[5])
		areaData.Area.CensusStats.Income.MedianHousehold = stringToFloat(row[7])
		areaData.Area.CensusStats.AvgPeoplePerHousehold = stringToFloat(row[8])
		return nil
	})
}

func stringToInt(s string) int {
	res, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("Failed to convert string to int - %v", s)
	}
	return res
}

func stringToFloat(s string) float64 {
	res, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Printf("Failed to convert string to int - %v", s)
	}
	return res
}
