package exchange

import (
	"testing"
	"errors"

	"github.com/go-chat-bot/bot"
    . "github.com/smartystreets/goconvey/convey"
)

func testingCommandGenerator(nick, channel, message string) bot.Cmd {
	return bot.Cmd{
		Raw:     message,
		Channel: channel,
		User: &bot.User{
			ID:       nick,
			Nick:     nick,
			RealName: nick,
			IsBot:    true,
		},
		MessageData: &bot.Message{
			Text:     message,
			IsAction: false,
		},
	}
}

func TestInitialize(t *testing.T) {

	Convey("As an authorized user, initialize a campaign", t, func() {
		testingCmd := testingCommandGenerator(
			"testadmin",
			"test",
			"!satans campaign init")

			got, err := initialize(&testingCmd)
			want := "welcome to satans"

			So(err, ShouldResemble, nil)
			So(got, ShouldEqual, want)
		},
	)

	Convey("As an unauthorized user, attempt to initialize a campaign", t, func() {
		testingCmd := testingCommandGenerator(
			"badguy",
			"test",
			"!satans campaign init")		// register 3 users

			got, err := initialize(&testingCmd)
			want := "no satans for u"

			So(err, ShouldResemble, errors.New("unauthorized"))
			So(got, ShouldEqual, want)
		},
	)
}
