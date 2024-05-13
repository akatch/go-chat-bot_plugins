package twitter

import (
	"errors"
	"testing"

	"github.com/go-chat-bot/bot"
	. "github.com/smartystreets/goconvey/convey"
)

func TestTwitter(t *testing.T) {
	// given a message string, I should get back a response message string
	// containing one or more parsed Tweets
	jbouieOutput := `b-boy bouiebaisse: This falls into one of my favorite genres of tweets, bona fide elites whose pretenses to understanding “common people” instead reveal their cloistered, condescending view of ordinary people.`

	var cases = []struct {
		input, output string
		expectedError error
	}{
		{
			input:         "this message has no links",
			output:        "",
			expectedError: nil,
		}, {
			input:         "http://twitter.com/jbouie/status/1247273759632961537",
			output:        jbouieOutput,
			expectedError: nil,
		}, {
			input:         "https://mobile.twitter.com/jbouie/status/1247273759632961537",
			output:        jbouieOutput,
			expectedError: nil,
		}, {
			input:         "wow check out this tweet https://mobile.twitter.com/jbouie/status/1247273759632961537",
			output:        jbouieOutput,
			expectedError: nil,
		}, {
			input:         "wow check out this tweethttps://mobile.twitter.com/jbouie/status/1247273759632961537",
			output:        jbouieOutput,
			expectedError: nil,
		}, {
			input:         "wow check out this tweet https://mobile.twitter.com/jbouie/status/1247273759632961537super cool right?",
			output:        jbouieOutput,
			expectedError: nil,
		}, {
			input:         "https://twitter.com/dmackdrwns/status/1217830568848764930/photo/1",
			output:        `David Mack: It was pretty fun to try to manifest creatures plucked right from the minds of manic children.  #georgiamuseumofart`,
			expectedError: nil,
		}, {
			input:         "http://twitter.com/notARealUser/status/123456789",
			output:        "",
			expectedError: errors.New("404 Not Found"),
		}, {
			input:         "https://x.com/terriblemaps/status/1789732549415117071",
			output:        "Terrible Maps: Cow-to-person ratio in Wisconsin",
			expectedError: nil,
		},
	}

	Convey("twitter", t, func() {
		for _, c := range cases {
			testingCmd := bot.PassiveCmd{
				Raw:         c.input,
				Channel:     "test",
				User:        &bot.User{ID: "test", Nick: "test", RealName: "test", IsBot: true},
				MessageData: &bot.Message{Text: c.input, IsAction: false},
			}
			Convey(c.input, func() {
				got, err := expandTweets(&testingCmd)
				want := c.output
				So(err, ShouldResemble, c.expectedError)
				So(got, ShouldEqual, want)
			})
		}
	},
	)
}
