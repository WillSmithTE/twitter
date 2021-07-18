package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/comprehend"
)

func GetToneAnalysis() {

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-2"),
	}))

	client := comprehend.New(sess)

	listTexts := []string{
		"Just had my second AstraZeneca shot. Feeling safer already. #GetTheJab",
	}

	for _, text := range listTexts {
		params := comprehend.DetectSentimentInput{}
		params.SetLanguageCode("en")
		params.SetText(text)

		req, resp := client.DetectSentimentRequest(&params)

		err := req.Send()
		if err == nil {
			fmt.Println(*resp)
		} else {
			fmt.Println(err)
		}
	}

}
