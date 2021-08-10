package vaccineDemographics

import (
	"log"
	"strconv"

	"github.com/willsmithte/twitter/src/util"
)

var sheetsDirPath = "src/vaccineDemographics/assets/2016_GCP_ALL_for_NSW_short-header/2016 Census GCP All Geographies for NSW/SA4/NSW/"

type AddData func(*Database) error

var AddDataFuncs = []AddData{addPopulation, addMedians, addReligion, addMotorVehicles, addHoursWorked, addAncestry}

func AddCensusData(database *Database) error {
	for _, addDataFunc := range AddDataFuncs {
		err := addDataFunc(database)
		if err != nil {
			return err
		}
	}
	return nil
}

func addPopulation(database *Database) error {
	filename := sheetsDirPath + "2016Census_G01_NSW_SA4.csv"
	err := util.ExecuteOnEachCsvRow(filename, func(row []string, _ []string) error {
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
		areaData.Area.CensusStats.Population = util.StringToInt(row[3])
		areaData.Area.CensusStats.Age.Num0to4 = util.StringToInt(row[6])
		areaData.Area.CensusStats.Age.Num5to14 = util.StringToInt(row[9])
		areaData.Area.CensusStats.Age.Num15to19 = util.StringToInt(row[12])
		areaData.Area.CensusStats.Age.Num20to24 = util.StringToInt(row[15])
		areaData.Area.CensusStats.Age.Num25to34 = util.StringToInt(row[18])
		areaData.Area.CensusStats.Age.Num35to44 = util.StringToInt(row[21])
		areaData.Area.CensusStats.Age.Num45to54 = util.StringToInt(row[24])
		areaData.Area.CensusStats.Age.Num55to64 = util.StringToInt(row[27])
		areaData.Area.CensusStats.Age.Num65to74 = util.StringToInt(row[30])
		areaData.Area.CensusStats.Age.Num75to84 = util.StringToInt(row[33])
		areaData.Area.CensusStats.Age.Num85Plus = util.StringToInt(row[36])
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func addMedians(database *Database) error {
	filename := sheetsDirPath + "2016Census_G02_NSW_SA4.csv"
	return util.ExecuteOnEachCsvRow(filename, func(row []string, _ []string) error {
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
		areaData.Area.CensusStats.Age.Median = util.StringToInt(row[1])
		areaData.Area.CensusStats.Income.MedianPersonal = util.StringToFloat(row[3])
		areaData.Area.CensusStats.Income.MedianFamily = util.StringToFloat(row[5])
		areaData.Area.CensusStats.Income.MedianHousehold = util.StringToFloat(row[7])
		areaData.Area.CensusStats.AvgPeoplePerHousehold = util.StringToFloat(row[8])
		return nil
	})
}

func addReligion(database *Database) error {
	filename := sheetsDirPath + "2016Census_G14_NSW_SA4.csv"
	return util.ExecuteOnEachCsvRow(filename, func(row []string, headerRow []string) error {
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

		for i := 3; i < len(row)-3; i += 3 {
			peopleCount := util.StringToInt(row[i])
			areaData.Area.CensusStats.Religion.Raw[headerRow[i]] = peopleCount
		}
		return nil
	})
}

func addMotorVehicles(database *Database) error {
	filename := sheetsDirPath + "2016Census_G30_NSW_SA4.csv"
	return util.ExecuteOnEachCsvRow(filename, func(row []string, headerRow []string) error {
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

		num0 := util.StringToInt(row[1])
		num1 := util.StringToInt(row[2])
		num2 := util.StringToInt(row[3])
		num3 := util.StringToInt(row[4])
		num4Plus := util.StringToInt(row[5])

		var totalNumDwellings int
		var totalNumVehicles int
		for numVehicles, numDwellings := range []int{num0, num1, num2, num3, num4Plus} {
			totalNumDwellings += numDwellings
			for i := 0; i < numDwellings; i++ {
				totalNumVehicles += numVehicles
			}
		}

		average := float64(totalNumVehicles) / float64(totalNumDwellings)

		areaData.Area.CensusStats.AverageMotorVehiclesPerDwelling = average

		return nil
	})
}

func addHoursWorked(database *Database) error {
	filename := sheetsDirPath + "2016Census_G52D_NSW_SA4.csv"
	return util.ExecuteOnEachCsvRow(filename, func(row []string, _ []string) error {
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
		totalNumPeople := util.StringToInt(row[30])

		areaData.Area.CensusStats.HoursWorked.Num0 = util.StringToInt(row[21])
		areaData.Area.CensusStats.HoursWorked.Num1to15 = util.StringToInt(row[22])
		areaData.Area.CensusStats.HoursWorked.Num16to24 = util.StringToInt(row[23])
		areaData.Area.CensusStats.HoursWorked.Num25to34 = util.StringToInt(row[24])
		areaData.Area.CensusStats.HoursWorked.Num35to39 = util.StringToInt(row[25])
		areaData.Area.CensusStats.HoursWorked.Num40 = util.StringToInt(row[26])
		areaData.Area.CensusStats.HoursWorked.Num41to48 = util.StringToInt(row[27])
		areaData.Area.CensusStats.HoursWorked.Num49Plus = util.StringToInt(row[28])
		areaData.Area.CensusStats.HoursWorked.NumNotSpecified = util.StringToInt(row[29])

		areaData.Area.CensusStats.HoursWorked.SetAverage(totalNumPeople)

		return nil
	})
}

func addAncestry(database *Database) error {
	filename := sheetsDirPath + "2016Census_G08_NSW_SA4.csv"
	return util.ExecuteOnEachCsvRow(filename, func(row []string, headerRow []string) error {
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

		for i := 6; i < len(row)-6; i += 6 {
			peopleCount := util.StringToInt(row[i])
			areaData.Area.CensusStats.Ancestry.Raw[headerRow[i]] = peopleCount
		}
		return nil
	})
}
