package util

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"os"
	"strconv"
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
	headerRow, err := csvr.Read()
	if err != nil {
		return err
	}

	for {
		row, err := csvr.Read()
		if err != nil {
			if err == io.EOF {
				return nil
			} else {
				return err
			}
		}

		err = consume(row, headerRow)
		if err != nil {
			return err
		}
	}
}

type ConsumeCsvRow func(row []string, headerRow []string) error

func StringToInt(s string) int {
	res, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("Failed to convert string to int - %v", s)
	}
	return res
}

func StringToFloat(s string) float64 {
	res, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Printf("Failed to convert string to int - %v", s)
	}
	return res
}
