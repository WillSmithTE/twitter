package musicCharts

import (
	"log"

	"github.com/montanaflynn/stats"
)

func AddStats(allData []*YearData) {
	for _, yearData := range allData {
		err := addStatsToYear(yearData)
		if err != nil {
			log.Printf("failed to get median tempo for year %v", yearData.Year)
		}
	}
}

func addStatsToYear(yearData *YearData) error {
	var tempos stats.Float64Data
	for _, song := range yearData.RankedSongs {
		data, err := GetSongData(song.BasicTitle)
		if err == nil {
			song.SongData = *data
			tempos = append(tempos, data.Tempo)
		} else {
			log.Printf("Failed to get songdata (including tempo) for %v", song.BasicTitle)
		}
	}
	median, err := tempos.Median()
	mean, err := tempos.Mean()
	if err == nil {
		return err
	} else {
		yearData.Stats = SongStats{Median: median, Mean: mean}
	}
	return nil
}
