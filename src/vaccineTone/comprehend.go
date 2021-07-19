package vaccineTone

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/comprehend"
)

func GetToneAnalysis(texts []string) []comprehend.DetectSentimentOutput {

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-2"),
	}))

	client := comprehend.New(sess)

	var scores []comprehend.DetectSentimentOutput

	for _, text := range texts {
		params := comprehend.DetectSentimentInput{}
		params.SetLanguageCode("en")
		params.SetText(text)

		req, resp := client.DetectSentimentRequest(&params)

		err := req.Send()
		if err == nil {
			scores = append(scores, *resp)
		} else {
			log.Printf("error getting sentiment for '%v' - %v", text, err)
		}
	}

	return scores
}
