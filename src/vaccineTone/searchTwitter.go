package vaccineTone

import (
	"compress/gzip"
	"encoding/gob"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func SearchAstra() (*TwitterSearchResponse, error) {
	return SearchTwitter("astrazeneca")
}

func SearchPfizer() (*TwitterSearchResponse, error) {
	return SearchTwitter("pfizer")
}

func SearchTwitter(query string) (*TwitterSearchResponse, error) {
	filename := "gobs/" + query

	gob.Register(map[string]interface{}{})
	gob.Register([]interface{}{})

	result := &TwitterSearchResponse{}
	err := result.Load(filename)

	if err == nil {
		log.Print("returning saved result")
		return result, nil
	}

	log.Print("no saved result found")

	searchUrl := "https://api.twitter.com/1.1/search/tweets.json?q=" + query + "&lang=en&result_type=recent&count=100&geocode=-33.865143,151.209900,25km&tweet_mode=extended"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", searchUrl, nil)
	req.Header.Set("Authorization", "Bearer "+os.Getenv("BEARER_TOKEN"))
	res, err := client.Do(req)

	if err == nil {
		defer res.Body.Close()
		result := &TwitterSearchResponse{}
		json.NewDecoder(res.Body).Decode(result)

		err = result.Save(filename)
		if err != nil {
			log.Printf("error saving result - %-v", err)
		}

		return result, nil
	} else {
		return nil, err
	}

}

func (t *TwitterSearchResponse) Load(filename string) error {

	fi, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fi.Close()

	fz, err := gzip.NewReader(fi)
	if err != nil {
		return err
	}
	defer fz.Close()

	decoder := gob.NewDecoder(fz)
	err = decoder.Decode(&t)
	if err != nil {
		return err
	}

	return nil
}

func (data *TwitterSearchResponse) Save(filename string) error {

	fi, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer fi.Close()

	fz := gzip.NewWriter(fi)
	defer fz.Close()

	encoder := gob.NewEncoder(fz)
	err = encoder.Encode(data)
	if err != nil {
		return err
	}

	return nil
}
