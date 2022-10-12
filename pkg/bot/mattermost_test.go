package bot

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMattermost_FindAndTrimBotMention(t *testing.T) {
	/// given
	botName := "Botkube"
	testCases := []struct {
		Name               string
		Input              string
		ExpectedTrimmedMsg string
		ExpectedFound      bool
	}{
		{
			Name:               "Mention",
			Input:              "@Botkube k get pods",
			ExpectedFound:      true,
			ExpectedTrimmedMsg: " k get pods",
		},
		{
			Name:               "Lowercase",
			Input:              "@botkube k get pods",
			ExpectedFound:      true,
			ExpectedTrimmedMsg: " k get pods",
		},
		{
			Name:               "Yet another different casing",
			Input:              "@BOTKUBE k get pods",
			ExpectedFound:      true,
			ExpectedTrimmedMsg: " k get pods",
		},
		{
			Name:          "Not at the beginning",
			Input:         "Not at the beginning @Botkube k get pods",
			ExpectedFound: false,
		},
		{
			Name:          "Different mention",
			Input:         "@bootkube k get pods",
			ExpectedFound: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			botMentionRegex, err := mattermostBotMentionRegex(botName)
			require.NoError(t, err)
			b := &Mattermost{botMentionRegex: botMentionRegex}
			require.NoError(t, err)

			// when
			actualTrimmedMsg, actualFound := b.findAndTrimBotMention(tc.Input)

			// then
			assert.Equal(t, tc.ExpectedFound, actualFound)
			assert.Equal(t, tc.ExpectedTrimmedMsg, actualTrimmedMsg)
		})
	}
}
