package exchange

import (

	"os"
	"errors"
	"strings"
	"github.com/go-chat-bot/bot"
)

var (
	ALLOWED_ADMINS = strings.Split(os.Getenv("EXCHANGE_ALLOWED_ADMINS"),",")
)

// initialize starts a new gift exchange Campaign
func initialize(command *bot.Cmd) (msg string, err error) {
	if (isAllowedAdmin(command.User)) {
		// check for existing campaign in this channel
		// create new one if not
		msg = "welcome to satans"
		err = nil
	} else {
		msg = "no satans for u"
		err = errors.New("unauthorized")
	}
	return
}

// isAllowedAdmin returns true if isNickservRegistered  AND user in EXCHANGE_ALLOWED_ADMINS
func isAllowedAdmin(user *bot.User) bool {
	for _, value := range ALLOWED_ADMINS {
		if (value == user.Nick) {
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
