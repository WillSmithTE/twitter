package musicCharts

import (
	"encoding/json"
	"log"
)

func Main() {
	yearData, err := GetAllTop100SongsByYear()
	if err == nil {
		AddMedianTempos(yearData)

		jsonData, _ := json.Marshal(yearData)
		log.Printf("Success - %v", string(jsonData))
	} else {
		log.Printf("error getting top 100s - %-v", err)
	}
}
