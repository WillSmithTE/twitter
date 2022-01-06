package vaccineDemographics

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Main() {
	database := buildDatabase()

	serve(database)
}

// Sources
// https://impfdashboard.de/en/
// https://interaktiv.morgenpost.de/corona-virus-karte-infektionen-deutschland-weltweit/

func buildDatabase() *Database {
	database, err := GetIncidenceRates()
	if err != nil {
		log.Panicf("error getting incidence rate data - %v", err)
	}
	err = AddVaccinationData(database)
	if err != nil {
		log.Panicf("error adding vacc rates - %v", err)
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
