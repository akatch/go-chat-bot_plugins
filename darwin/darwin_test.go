package darwin

import (
	"fmt"
	"github.com/go-chat-bot/bot"
	"testing"
)

func TestDarwin(t *testing.T) {
	var cases = []struct {
		text string
		args []string
	}{
		{fmt.Sprintf(darwinQuote, "nick"), nil},
		{fmt.Sprintf(darwinQuote, "someothernick"), []string{"someothernick"}},
	}
	for _, c := range cases {
		t.Run(c.text, func(t *testing.T) {
			bot := &bot.Cmd{
				Args:    c.args,
				Command: "darwin",
				User: &bot.User{
					Nick:     "nick",
					RealName: "Real Name",
				},
			}
			got, err := darwin(bot)
			if err != nil {
				t.Error(err)
			}

			if c.text != got {
				t.Errorf("got %s; want %s", got, c.text)
			}
		})

	}
}
