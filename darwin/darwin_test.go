package darwin

import (
	"fmt"
	"testing"

	"github.com/go-chat-bot/bot"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDarwin(t *testing.T) {
	var cases = []struct {
		text string
		args []string
	}{
		{fmt.Sprintf(darwinQuote, "test"), nil},
		{fmt.Sprintf(darwinQuote, "someothernick"), []string{"someothernick"}},
	}
	for i, c := range cases {
		Convey(fmt.Sprintf("Case %d", i), t, func() {
			bot := &bot.Cmd{
				Args:    c.args,
				Command: "darwin",
				User: &bot.User{
					Nick:     "test",
					RealName: "test",
				},
			}
			got, err := darwin(bot)
			want := c.text

			So(err, ShouldResemble, nil)
			So(got, ShouldEqual, want)
		})

	}
}
