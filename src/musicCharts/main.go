package musicCharts

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Main() {
	// songs := GetYearData(2019)
	// log.Printf("%v", songs.RankedSongs[1])
	getAndServeSongs()
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

	err := http.ListenAndServe(":5000", nil)
	if err == nil {
		log.Print("serving now")
	} else {
		log.Fatal(err)
	}
}

func Serve(data []*YearData) {
	handleRequests(data)
}
