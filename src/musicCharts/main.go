package musicCharts

import "log"

func Main() {
	tempo, err := GetTempo()
	if err == nil {
		log.Printf("success - %-v", *tempo)
	} else {
		log.Printf("error - %-v", err)
	}
}
