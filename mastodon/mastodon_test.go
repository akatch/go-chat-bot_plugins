package mastodon

import (
	"errors"
	"testing"

	"github.com/go-chat-bot/bot"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMastodon(t *testing.T) {
	// Mastodon injects weird span tags into Statuses which render as spaces. Dunno why.
	barOutput := `Toot from John Watson: hey y'all i made this for # Godot in case you want to be awesome https:// github.com/jotson/ridiculous_c oding`
	barURL := `https://mastodon.gamedev.place/@jotson/109367069016579141`
	bazOutput := `Toot from John Watson: hey y'all i made this for # Godot in case you want to be awesome https:// github.com/jotson/ridiculous_c oding`
	bazURL := `https://fosstodon.org/@jotson@mastodon.gamedev.place/109367069171884095`
	quuxOutput := `Toot from Andrew Nadeau: [after leaving willy wonka’s factory] me: wife: me: wife: me: wife: lot of deaths for a to— me: a LOT of deaths for a tour!`
	quuxURL := `https://mastodon.social/@AndrewNadeau/109361740736612801`
	fredOutput := `Toot from yan: nobody:  absolutely nobody:  yubikey: cccjgjgkhcbbcvchfkfhiiuunbtnvgihdfiktncvlhck`
	fredURL := `https://infosec.exchange/@bcrypt/109341144608902590`
	thudOutput := `Toot from Tom Grochowiak: Just realised that since I'm new here, I haven't yet spammed my fave (and most useless) project this year.  A genetic algorithm that attempts to give birth to a designated image.  # gamedev # procedural # procgen`
	thudURL := `https://mastodon.gamedev.place/@tomgrochowiak/109365654828117404`

	var cases = []struct {
		input, output string
		expectedError error
	}{
		{
			input:         "this message has no links",
			output:        "",
			expectedError: nil,
		}, {
			input:         "wow check out this tweet" + barURL,
			output:        barOutput,
			expectedError: nil,
		}, {
			input:         "wow check out this tweet " + barURL + " super cool right?",
			output:        barOutput,
			expectedError: nil,
		}, {
			input:         bazURL,
			output:        bazOutput,
			expectedError: nil,
		}, {
			input:         "https://mastodon.social/statuses/123456789",
			output:        "",
			expectedError: errors.New("bad request: 404 Not Found: Record not found"),
		}, {
			input:         barURL + " lol bye",
			output:        barOutput,
			expectedError: nil,
		}, {
			input:         quuxURL,
			output:        quuxOutput,
			expectedError: nil,
		}, {
			input:         fredURL,
			output:        fredOutput,
			expectedError: nil,
		}, {
			input:         thudURL,
			output:        thudOutput,
			expectedError: nil,
		},
	}

	Convey("mastodon", t, func() {
	for _, c := range cases {
		testingUser := bot.User{
			ID:       "test",
			Nick:     "test",
			RealName: "test",
			IsBot:    true,
		}
		testingMessage := bot.Message{
			Text:     c.input,
			IsAction: false,
		}
		testingCmd := bot.PassiveCmd{
			Raw:         c.input,
			Channel:     "test",
			User:        &testingUser,
			MessageData: &testingMessage,
		}
		Convey(c.input, func() {
			got, err := expandToots(&testingCmd)
			want := c.output

			So(err, ShouldEqual, c.expectedError)
			So(got, ShouldEqual, want)
		})
	}
},
)
}
