package musicCharts

import (
	"log"

	"github.com/willsmithte/twitter/src/util"
)

func Main() {
	data, err := GetAllTop100SongsByYear()
	if err == nil {
		AddStats(data)

		util.PrintJson(data)
	} else {
		log.Printf("error getting top 100s - %-v", err)
	}
}
