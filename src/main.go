package main

import (
	"log"

	"github.com/willsmithte/twitter/src/vaccineTone"
)

func main() {

}

func DoTwitterAnalysis() {
	res, err := vaccineTone.SearchTwitter("Donald+Trump")
	if err == nil {
		tone := vaccineTone.ComprehendTwitter(*res, "Trump")
		log.Printf("analysis complete - %-v", tone)
	} else {
		log.Printf("error - %-v", err)
	}
}
