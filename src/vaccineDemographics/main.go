package vaccineDemographics

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// sources
// https://www.health.gov.au/sites/default/files/documents/2021/08/covid-19-vaccination-geographic-vaccination-rates-2-august-2021.pdf
// https://www.health.gov.au/sites/default/files/documents/2021/08/covid-19-vaccination-doses-by-age-and-sex_1.pdf
// https://quickstats.censusdata.abs.gov.au/census_services/getproduct/census/2016/quickstat/121?opendocument

func Main() {
	// Parse vaccination rates csv.
	// Create all rows as AreaDatas
	// Add CensusCode to all from that csv
	// Add stats
	// ?? add national vacc-age data

	database := buildDatabase()

	serve(database)
}

func buildDatabase() *Database {
	database, err := GetGeographicalVaccData()
	if err != nil {
		log.Panicf("error getting data - %v", err)
	}
	err = AddCensusCodes(database)
	if err != nil {
		log.Panicf("error adding census codes - %v", err)
	}
	err = AddCensusData(database)
	if err != nil {
		log.Panicf("error adding census data - %v", err)
	}
	return database
}

func serve(data *Database) {
	handleRequests(data)
}

func handleRequests(data *Database) {
	r := mux.NewRouter()
	r.HandleFunc(
		"/api/areas",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			err := json.NewEncoder(w).Encode(data)
			if err != nil {
				log.Fatalf("error encoding data - %v", err)
			}
		},
	)
	http.Handle("/", r)

	log.Print("Starting server")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
