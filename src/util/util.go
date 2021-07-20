package util

import (
	"encoding/json"
	"log"
)

func PrintJson(toPrint interface{}) {
	jsonData, _ := json.Marshal(toPrint)
	log.Printf("%v", string(jsonData))
}
