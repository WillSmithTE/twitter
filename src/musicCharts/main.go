package musicCharts

import (
	"log"

	"github.com/willsmithte/twitter/src/util"
)

func Main() {
	data, err := GetAllTop100SongsByYear()
	if err == nil {
		// AddStats(yearData)

		var found *YearData

		for _, yearData := range data {
			if yearData.Year == 1945 {
				found = yearData
				break
			}
		}

		util.PrintJson(found)
	} else {
		log.Printf("error getting top 100s - %-v", err)
	}
}
