package main

import "fmt"

func main() {
	res, err := SearchAstra()
	if err == nil {
		fmt.Printf("success! - %-v", res)
	} else {
		fmt.Printf("error - %-v", err)
	}
}
