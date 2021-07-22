package musicCharts

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Main() {
	// getAndServeSongs()
	data, _ := GetAllTop100SongsByYear()
	var problemYearsData []*YearData
	for _, data := range data {
		if data.Year == 1943 || data.Year == 1944 || data.Year == 2013 || data.Year == 2015 || data.Year == 2020 {
			problemYearsData = append(problemYearsData, data)
		}
	}

	for _, data := range problemYearsData {
		log.Print(*data)

	}
}

func getAndServeSongs() {
	data, err := GetAllTop100SongsByYear()
	if err == nil {
		AddStats(data)
	} else {
		log.Printf("error getting top 100s - %-v", err)
	}
	Serve(data)

}

func handleRequests(data []*YearData) {
	r := mux.NewRouter()
	r.HandleFunc(
		"/api/years",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			json.NewEncoder(w).Encode(data)
		},
	)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":5000", nil))
}

func Serve(data []*YearData) {
	handleRequests(data)
}
