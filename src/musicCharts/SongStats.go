package musicCharts

import (
	"log"

	"github.com/montanaflynn/stats"
)

func AddStats(allData []*YearData) {
	for _, yearData := range allData {
		addStatsToYear(yearData)
	}
}

func addStatsToYear(yearData *YearData) {
	var tempos stats.Float64Data
	for _, song := range yearData.RankedSongs {
		data, err := GetSongData(song.BasicTitle)
		if err == nil {
			song.SongData = *data
			if song.SongData.Tempo != 0 {
				tempos = append(tempos, data.Tempo)
			}
		} else {
			log.Printf("Failed to get songdata for %v - %v", song.BasicTitle, err)
		}
	}
	median, err := tempos.Median()
	if err != nil {
		log.Printf("failed to get median for year %v - %v", yearData.Year, err)
	}
	mean, err := tempos.Mean()
	if err != nil {
		log.Printf("failed to get mean for year %v - %v", yearData.Year, err)
	}
	yearData.Stats = SongStats{Median: median, Mean: mean}
}
