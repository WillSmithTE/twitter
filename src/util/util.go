package util

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"os"
)

func PrintJson(toPrint interface{}) {
	jsonData, _ := json.Marshal(toPrint)
	log.Printf("%v", string(jsonData))
}

func ExecuteOnEachCsvRow(filename string, consume ConsumeCsvRow) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	csvr := csv.NewReader(f)
	csvr.Read() // skip headers row

	for {
		row, err := csvr.Read()
		if err != nil {
			if err == io.EOF {
				return nil
			} else {
				return err
			}
		}

		err = consume(row)
		if err != nil {
			return err
		}
	}
}

type ConsumeCsvRow func([]string) error
