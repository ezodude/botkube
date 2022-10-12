package bot

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTeams_TrimBotMention(t *testing.T) {
	/// given
	botName := "Botkube"
	testCases := []struct {
		Name               string
		Input              string
		ExpectedTrimmedMsg string
	}{
		{
			Name:               "Mention",
			Input:              "<at>Botkube</at> k get pods",
			ExpectedTrimmedMsg: " k get pods",
		},
		{
			Name:               "Not at the beginning",
			Input:              "Not at the beginning <at>Botkube</at> k get pods",
			ExpectedTrimmedMsg: "Not at the beginning <at>Botkube</at> k get pods",
		},
		{
			Name:               "Different mention",
			Input:              "<at>bootkube</at> k get pods",
			ExpectedTrimmedMsg: "<at>bootkube</at> k get pods",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			botMentionRegex, err := teamsBotMentionRegex(botName)
			require.NoError(t, err)
			b := &Teams{botMentionRegex: botMentionRegex}
			require.NoError(t, err)

			// when
			actualTrimmedMsg := b.trimBotMention(tc.Input)

			// then
			assert.Equal(t, tc.ExpectedTrimmedMsg, actualTrimmedMsg)
		})
	}
}
