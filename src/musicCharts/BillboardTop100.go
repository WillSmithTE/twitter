package musicCharts

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/anaskhan96/soup"
)

func GetAllTop100SongsByYear() ([]*YearData, error) {
	currentYear := time.Now().Format("2006")
	yearAsInt, _ := strconv.Atoi(currentYear)

	var allYearsData []*YearData

	for i := 1941; i <= int(yearAsInt); i++ {
		data := YearData{Year: i}

		filename := "gobs/years/" + fmt.Sprint(i)
		err := data.Load(filename)

		if err == nil {
			log.Printf("Saved data found for year %v", i)
		} else {
			log.Printf("Failed to load saved data for year %v", i)

			songs, err := getSongsForYear(&data)
			if err == nil {
				log.Printf("Data collected successfully for year %v", i)
				data.RankedSongs = songs
			} else {
				log.Printf("Failed to get top 100 for year %v - %v", i, err)
			}
			err = data.Save(filename)
			if err == nil {
				log.Printf("Saved data for year %v", i)
			} else {
				log.Printf("Failed to save data for year %v - %v", i, err)
			}
		}
		allYearsData = append(allYearsData, &data)
	}

	return allYearsData, nil
}

func getSongsForYear(yearData *YearData) ([]*Song, error) {
	resp, err := soup.Get("http://billboardtop100of.com/" + fmt.Sprint(yearData.Year) + "-2/")
	if err != nil {
		return nil, err
	}
	doc := soup.HTMLParse(resp)

	songs, err := scrapeListTable(doc)
	if err == nil {
		return songs, nil
	}
	songs, err = scrapeListPTags(doc)
	if err == nil {
		return songs, nil
	}
	songs, err = scrapeListArticles(doc)
	if err == nil {
		return songs, nil
	}
	return nil, err
}

func scrapeListTable(doc soup.Root) ([]*Song, error) {
	tbody := doc.Find("tbody")
	if tbody.Error != nil {
		return nil, tbody.Error
	}
	rows := tbody.FindAll("tr")
	if len(rows) == 0 {
		return nil, errors.New("no rows found on table")
	}
	var songs []*Song
	for _, row := range rows {
		tds := row.FindAll("td")
		if len(tds) == 3 {
			author := tds[1].Children()[0].Text()
			if author == "" {
				author = tds[1].Text()
			}
			title := tds[2].Text()
			combinedText := title + " " + author
			cleaned := cleanString(combinedText)
			songs = append(songs, &Song{BasicTitle: cleaned})
		} else {
			return nil, errors.New("expected 3 cells in table row, found " + fmt.Sprint(len(tds)))
		}
	}
	return songs, nil
}

func scrapeListPTags(doc soup.Root) ([]*Song, error) {
	div := doc.Find("div", "class", "entry-content")
	if div.Error != nil {
		return nil, div.Error
	}
	pTags := div.FindAll("p")
	if len(pTags) == 0 {
		return nil, errors.New("no ptags found")
	}

	var songs []*Song
	for _, pTag := range pTags {
		text := pTag.Text()
		indexOfDot := strings.Index(text, ".")
		withoutRank := text[indexOfDot+1:]
		withoutBy := strings.ReplaceAll(withoutRank, "by ", "")
		cleaned := cleanString(withoutBy)
		songs = append(songs, &Song{BasicTitle: cleaned})
	}

	return songs, nil
}

func scrapeListArticles(doc soup.Root) ([]*Song, error) {
	divs := doc.FindAll("div", "class", "ye-chart-item__text")
	if len(divs) == 0 {
		return nil, errors.New("no divs found")
	}
	var songs []*Song
	for _, div := range divs {
		insideDivs := div.FindAll("div")
		if len(insideDivs) == 0 {
			return nil, errors.New("no divs found")
		}
		title, artist := insideDivs[0].Text(), insideDivs[1].Text()
		text := cleanString(title + " " + artist)
		songs = append(songs, &Song{BasicTitle: text})
	}
	return songs, nil
}

func cleanString(before string) string {
	cleaned := strings.Replace(before, " and His Orchestra", "", 1)
	cleaned = strings.ReplaceAll(cleaned, " featuring", "")
	cleaned = strings.ReplaceAll(cleaned, " Featuring", "")
	cleaned = strings.ReplaceAll(cleaned, " &", "")
	return cleaned
}

type YearData struct {
	Stats       SongStats
	RankedSongs []*Song
	Year        int
}

type Song struct {
	BasicTitle string
	SongData   SongData
}

type SongStats struct {
	Median float64
	Mean   float64
}