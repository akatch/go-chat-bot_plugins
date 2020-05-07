package darwin

import (
	"fmt"
	"github.com/go-chat-bot/bot"
)

const darwinQuote = "%s is feeling very poorly today and very stupid and hates everyone & everything"

func darwin(command *bot.Cmd) (msg string, err error) {
	// 1% of the time: "It is a beautiful day at The Red Pony, a continual soirée, this is %s."
	target := command.User.Nick
	if len(command.Args) > 0 {
		target = command.Args[0]
	}

	msg = fmt.Sprintf(darwinQuote, target)
	return
}

func init() {
	bot.RegisterCommand(
		"darwin",
		"Sends a bright, uplifting Charles Darwin quote",
		"[zeroCool]",
		darwin)
}
