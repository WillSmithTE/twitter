package vaccineTone

import "log"

func DoTwitterVaccineAnalysis() {
	res, err := SearchTwitter("Donald+Trump")
	if err == nil {
		tone := ComprehendTwitter(*res, "Trump")
		log.Printf("analysis complete - %-v", tone)
	} else {
		log.Printf("error - %-v", err)
	}
}
