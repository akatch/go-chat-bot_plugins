package exchange

import (
	"errors"
	"os"
	"strings"

	"github.com/go-chat-bot/bot"
)

var (
	ALLOWED_ADMINS = strings.Split(os.Getenv("EXCHANGE_ALLOWED_ADMINS"), ",")
)

// A Participant is a user that has chosen to participate in a gift exchange Campaign
type Participant struct {
	user           *bot.User
	address, notes string
	recipientID    int
}

type Campaign struct {
	channel      string
	participants []*Participant
}

// initialize starts a new gift exchange Campaign
func initialize(command *bot.Cmd) (msg string, err error) {
	if isAllowedAdmin(command.User) {
		_, err = NewCampaign(command)
		msg = "welcome to satans"
	} else {
		msg = "no satans for u"
		err = errors.New("unauthorized")
	}
	return
}

func NewCampaign(cmd *bot.Cmd) (*Campaign, error) {
	// check for existing campaign in this channel
	// create new one if not
	// SELECT * FROM campaigns WHERE 'channel' = cmd.Channel
	c := Campaign{
		channel:      cmd.Channel,
		participants: nil,
	}
	return &c, nil
}

// isAllowedAdmin returns true if isNickservRegistered  AND user in EXCHANGE_ALLOWED_ADMINS
func isAllowedAdmin(user *bot.User) bool {
	for _, value := range ALLOWED_ADMINS {
		if value == user.Nick {
			return true
		}
	}
	return false
}

// init initalizes Commands for managing a gift exchange campaign
func init() {
	bot.RegisterCommand(
		"satans campaign init",
		"Initialize a gift exchange campaign",
		"",
		initialize)
}
