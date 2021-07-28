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

	for i := 1941; i < int(yearAsInt); i++ {
		data := GetYearData(i)
		allYearsData = append(allYearsData, &data)
	}

	return allYearsData, nil
}

func GetYearData(year int) YearData {
	data := YearData{Year: year}

	filename := "gobs/years/" + fmt.Sprint(year)
	err := data.Load(filename)

	if err == nil {
		// log.Printf("Saved data found for year %v", i)
	} else {
		log.Printf("Failed to load saved data for year %v", year)

		songs, err := getSongsForYear(&data)
		if err == nil {
			log.Printf("Data collected successfully for year %v", year)
			data.RankedSongs = songs
			err = data.Save(filename)
			if err == nil {
				log.Printf("Saved data for year %v", year)
			} else {
				log.Printf("Failed to save data for year %v - %v", year, err)
			}
		} else {
			log.Printf("Failed to get top 100 for year %v - %v", year, err)
		}
	}
	return data
}

func getSongsForYear(yearData *YearData) ([]*Song, error) {
	resp, err := soup.Get("http://billboardtop100of.com/" + fmt.Sprint(yearData.Year) + "-2/")
	if err != nil {
		return nil, err
	}
	doc := soup.HTMLParse(resp)

	songs, err := scrapeListTable(doc)
	if err == nil {
		log.Print("found in table")
		return songs, nil
	}
	songs, err = scrapeListPTags(doc)
	if err == nil {
		log.Print("found in ptags")
		return songs, nil
	}
	songs, err = scrapeListArticles(doc)
	if err == nil {
		log.Print("found in articles")
		return songs, nil
	}
	songs, err = scrapeListPTags2013(doc)
	if err == nil {
		log.Print("found in ptags 2013")
		return songs, nil
	}
	songs, err = scrapeListTable2015(doc)
	if err == nil {
		log.Print("found in table 2015")
		return songs, nil
	}
	songs, err = scrapeListDivs2020(doc)
	if err == nil {
		log.Print("found in listdivs2020")
		return songs, nil
	}
	songs, err = scrapeListPTags1942(doc)
	if err == nil {
		log.Print("found in listptags1942")
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
			cleaned := CleanString(combinedText)
			if strings.Trim(cleaned, " ") == "" {
				return nil, errors.New("got empty string in table")
			}
			songs = append(songs, &Song{BasicTitle: cleaned})
		} else {
			return nil, errors.New("expected 3 cells in table row, found " + fmt.Sprint(len(tds)))
		}
	}
	return songs, nil
}

// 1943, 1944
func scrapeListPTags(doc soup.Root) ([]*Song, error) {
	div := doc.Find("div", "class", "entry-content")
	if div.Error != nil {
		return nil, div.Error
	}
	ptag := div.Find("p")
	if ptag.Error != nil {
		return nil, ptag.Error
	}

	var songs []*Song

	for _, song := range ptag.Children() {
		childText := song.NodeValue
		if childText != "br" {
			indexOfDot := strings.Index(childText, ".")
			withoutRank := childText[indexOfDot+1:]
			withoutDash := strings.Replace(withoutRank, " –", "", 1)
			cleaned := CleanString(withoutDash)
			songs = append(songs, &Song{BasicTitle: cleaned})
		}
	}

	if len(songs) < 5 { // 2013
		return nil, errors.New("less than 5 songs found")
	}

	return songs, nil
}

// 2013
func scrapeListPTags2013(doc soup.Root) ([]*Song, error) {
	div := doc.Find("div", "class", "entry-content")
	if div.Error != nil {
		return nil, div.Error
	}
	ptags := div.FindAll("p")
	if len(ptags) != 3 {
		return nil, errors.New("expected 3 ptags, was " + fmt.Sprint(len(ptags)))
	}
	lastP := ptags[len(ptags)-1]

	listOfSongs := lastP.Children()[0].Children()
	var songs []*Song
	log.Printf("listofsongs: %v", listOfSongs)

	for _, song := range listOfSongs {
		childText := song.NodeValue
		if childText != "br" {
			indexOfDot := strings.Index(childText, ".")
			withoutRank := childText[indexOfDot+1:]
			withoutDash := strings.Replace(withoutRank, " –", "", 1)
			cleaned := CleanString(withoutDash)
			songs = append(songs, &Song{BasicTitle: cleaned})
		}
	}

	return songs, nil
}

// 2015
func scrapeListTable2015(doc soup.Root) ([]*Song, error) {
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
			author := tds[1].Find("h6").Text()
			title := tds[2].Find("h6").Text()
			combinedText := title + " " + author
			cleaned := CleanString(combinedText)
			songs = append(songs, &Song{BasicTitle: cleaned})
		} else {
			return nil, errors.New("expected 3 cells in table row, found " + fmt.Sprint(len(tds)))
		}
	}
	return songs, nil
}

// 2020
func scrapeListDivs2020(doc soup.Root) ([]*Song, error) {
	entryContent := doc.Find("div", "class", "entry-content")
	if entryContent.Error != nil {
		return nil, entryContent.Error
	}
	divs := entryContent.FindAll("div", "class", "wp-block-media-text")
	if len(divs) == 0 {
		return nil, errors.New("didn't find any divs")
	}
	var songs []*Song
	for _, row := range divs {
		insideDiv := row.Find("div")
		if insideDiv.Error != nil {
			return nil, insideDiv.Error
		}
		ptags := insideDiv.FindAll("p")
		if len(ptags[0].Children()) == 3 {
			title := ptags[0].Children()[2].Text()
			var author string
			if len(ptags[1].Children()) == 1 {
				author = ptags[1].Children()[0].NodeValue
			} else {
				author = ptags[1].Text()
			}
			combinedText := title + " " + author
			cleaned := CleanString(combinedText)
			songs = append(songs, &Song{BasicTitle: cleaned})
		}
	}
	return songs, nil
}

// 1942
func scrapeListPTags1942(doc soup.Root) ([]*Song, error) {
	entryContent := doc.Find("div", "class", "entry-content")
	if entryContent.Error != nil {
		return nil, entryContent.Error
	}
	ptags := entryContent.FindAll("p")
	if len(ptags) == 0 {
		return nil, errors.New("didn't find any ptags")
	}
	var songs []*Song
	for _, ptag := range ptags {
		fullText := ptag.Text()
		indexOfDot := strings.Index(fullText, ".")
		withoutRank := fullText[indexOfDot+1:]
		cleaned := CleanString(withoutRank)
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
		text := CleanString(title + " " + artist)
		songs = append(songs, &Song{BasicTitle: text})
	}
	return songs, nil
}

func CleanString(before string) string {
	cleaned := strings.Replace(before, " and His Orchestra", "", 1)
	cleaned = strings.ReplaceAll(cleaned, " featuring", "")
	cleaned = strings.ReplaceAll(cleaned, " Featuring", "")
	cleaned = strings.ReplaceAll(cleaned, " Feat.", "")
	cleaned = strings.ReplaceAll(cleaned, " feat.", "")
	cleaned = strings.ReplaceAll(cleaned, " &", "")
	cleaned = strings.ReplaceAll(cleaned, "/", " ")
	cleaned = strings.ReplaceAll(cleaned, "/", " ")
	cleaned = strings.ReplaceAll(cleaned, "by ", "")
	cleaned = strings.ReplaceAll(cleaned, "&nbsp;", " ")
	cleaned = strings.ReplaceAll(cleaned, "– ", " ")
	cleaned = strings.ReplaceAll(cleaned, "- ", " ")
	cleaned = strings.ReplaceAll(cleaned, " by ", " ")
	cleaned = strings.ReplaceAll(cleaned, " x ", " ")
	cleaned = strings.ReplaceAll(cleaned, " X ", " ")
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
