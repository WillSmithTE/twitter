package vaccineTone

import (
	"github.com/aws/aws-sdk-go/service/comprehend"
)

func ComprehendTwitter(twitterResponse TwitterSearchResponse, title string) ComprehendedTwitter {
	neutral := 0.0
	mixed := 0.0
	positive := 0.0
	negative := 0.0

	var allTwitterTexts []string

	for _, status := range twitterResponse.Statuses {
		allTwitterTexts = append(allTwitterTexts, status.FullText)
	}

	allToneAnalysis := GetToneAnalysis(allTwitterTexts)

	for _, tone := range allToneAnalysis {
		neutral += *tone.SentimentScore.Neutral
		mixed += *tone.SentimentScore.Mixed
		positive += *tone.SentimentScore.Positive
		negative += *tone.SentimentScore.Negative
	}

	return ComprehendedTwitter{
		Title: title,
		Score: comprehend.SentimentScore{
			Mixed:    &mixed,
			Neutral:  &neutral,
			Negative: &negative,
			Positive: &positive,
		},
	}
}

type ComprehendedTwitter struct {
	Title string
	Score comprehend.SentimentScore
}
