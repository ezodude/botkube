//go:build integration

package e2e

import (
	"regexp"
	"testing"

	"github.com/kubeshop/botkube/pkg/bot/interactive"
	"github.com/sanity-io/litter"
)

const recentMessagesLimit = 5

// structDumper provides an option to print the struct in more readable way.
var structDumper = litter.Options{
	HidePrivateFields: true,
	HideZeroValues:    true,
	StripPackageNames: false,
	FieldExclusions:   regexp.MustCompile(`^(XXX_.*)$`), // XXX_ is a prefix of fields generated by protoc-gen-go
	Separator:         " ",
}

type MessageAssertion func(content string) bool
type AttachmentAssertion func(title, color, msg string) bool

type Channel interface {
	ID() string
	Name() string
	Identifier() string
}

// DriverType to instrument
type DriverType string

const (
	// CreateEvent when resource is created
	SlackBot DriverType = "slack"
	// UpdateEvent when resource is updated
	DiscordBot DriverType = "discord"
)

type BotDriver interface {
	Type() DriverType
	InitUsers(t *testing.T)
	InitChannels(t *testing.T) []func()
	PostInitialMessage(t *testing.T, channel string)
	PostMessageToBot(t *testing.T, channel, command string)
	InviteBotToChannel(t *testing.T, channel string)
	WaitForMessagePostedRecentlyEqual(userID, channelID, expectedMsg string) error
	WaitForLastMessageContains(userID, channel, expectedMsgSubstring string) error
	WaitForLastMessageEqual(userID, channel, expectedMsg string) error
	WaitForMessagePosted(userID, channel string, limitMessages int, assertFn MessageAssertion) error
	WaitForInteractiveMessagePosted(userID, channelID string, limitMessages int, assertFn MessageAssertion) error
	WaitForMessagePostedWithAttachment(userID, channel string, assertFn AttachmentAssertion) error
	WaitForMessagesPostedOnChannelsWithAttachment(userID string, channelIDs []string, assertFn AttachmentAssertion) error
	Channel() Channel
	SecondChannel() Channel
	BotUserID() string
	TesterUserID() string
	WaitForInteractiveMessagePostedRecentlyEqual(userID string, channelID string, message interactive.Message) error
}
