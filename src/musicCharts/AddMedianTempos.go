package musicCharts

import (
	"log"

	"github.com/montanaflynn/stats"
)

func AddMedianTempos(allData []*YearData) {
	for _, yearData := range allData {
		err := addTemposToYear(yearData)
		if err != nil {
			log.Printf("failed to get median tempo for year %v", yearData.Year)
		}
	}
}

func addTemposToYear(yearData *YearData) error {
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
	yearMedian, err := tempos.Median()
	if err == nil {
		return err
	} else {
		yearData.MedianTempo = yearMedian
	}
	return nil
}
