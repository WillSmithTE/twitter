package musicCharts

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func GetSongData(query string) (*SongData, error) {
	filename := "gobs/songs/" + query

	songData := &SongData{}
	err := songData.Load(filename)

	if err == nil {
		log.Printf("Saved data found for song %v", query)
		return songData, nil
	}

	log.Printf("Failed to load saved data for song %v", query)

	searchId, err := getSearchId(query)

	if err != nil {
		return nil, err
	}

	songData, err = getFullSongData(*searchId)

	if err != nil {
		return nil, err
	}

	err = songData.Save(filename)

	if err == nil {
		log.Printf("Saved data for song %v", query)
	} else {
		log.Printf("Failed to save data for song %v - %v", query, err)
	}

	return songData, nil

}

func getSearchId(query string) (*string, error) {

	type Payload struct {
		Query string `json:"query"`
	}

	data := Payload{
		Query: query,
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://songbpm.com/api/searches", body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authority", "songbpm.com")
	req.Header.Set("Sec-Ch-Ua", "\" Not;A Brand\";v=\"99\", \"Google Chrome\";v=\"91\", \"Chromium\";v=\"91\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Origin", "https://songbpm.com")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://songbpm.com/searches/29c3ba4a-4406-4918-a582-38f04bd39f90")
	req.Header.Set("Accept-Language", "en-GB,en-US;q=0.9,en;q=0.8")
	req.Header.Set("Cookie", "_ga=GA1.2.1093947617.1626665372; _gid=GA1.2.1562353574.1626665372")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	idResponse := &IdResponse{}
	err = json.NewDecoder(resp.Body).Decode(idResponse)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &idResponse.ID, nil
}

func getFullSongData(searchId string) (*SongData, error) {

	req, err := http.NewRequest("GET", "https://songbpm.com/_next/data/gpiNoF_Bm5T8NtijFyVkm/searches/"+searchId+".json?id="+searchId, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authority", "songbpm.com")
	req.Header.Set("Sec-Ch-Ua", "\" Not;A Brand\";v=\"99\", \"Google Chrome\";v=\"91\", \"Chromium\";v=\"91\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://songbpm.com/searches/"+searchId)
	req.Header.Set("Accept-Language", "en-GB,en-US;q=0.9,en;q=0.8")
	req.Header.Set("Cookie", "_ga=GA1.2.1093947617.1626665372; _gid=GA1.2.1562353574.1626665372")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	songSearchResponse := &SongSearchResponse{}
	err = json.NewDecoder(resp.Body).Decode(songSearchResponse)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	songs := songSearchResponse.PageProps.Songs
	if len(songs) == 0 {
		return nil, errors.New("song not found")
	}

	return &songs[0], nil
}

type IdResponse struct {
	ID string `json:"id"`
}

type SongSearchResponse struct {
	PageProps struct {
		Search struct {
			ID    string `json:"id"`
			Query string `json:"query"`
		} `json:"search"`
		Songs []SongData `json:"songs"`
	} `json:"pageProps"`
	NSSP bool `json:"__N_SSP"`
}

type SongData struct {
	ID   string `json:"id"`
	Data struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Album struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"album"`
	} `json:"data"`
	Name   string  `json:"name"`
	Slug   string  `json:"slug"`
	Tempo  float64 `json:"tempo"`
	Artist struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"artist"`
}
