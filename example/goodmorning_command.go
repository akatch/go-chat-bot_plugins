package example

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-chat-bot/bot"
)

func goodMorning() (msg string, err error) {
	msg = fmt.Sprintf("Good morning, channel!")
	return
}

func init() {
	// A comma separated list of channel ids from environment
	channels := strings.Split(os.Getenv("CHANNEL_IDS"), ",")

	if len(channels) > 0 {
		// Greets channel at 8am every week day
		config := &bot.PeriodicConfig{CronSpec: "0 0 08 * * mon-fri", Channels: channels}

		bot.RegisterPeriodicCommand(
			config,
			goodMorning)
	}
}
