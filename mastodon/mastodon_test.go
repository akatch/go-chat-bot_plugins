package mastodon

import (
	"errors"
	"fmt"
	"testing"

	"github.com/go-chat-bot/bot"
)

func TestMastodon(t *testing.T) {
	// given a message string, I should get back a response message string
	// containing one or more parsed Tweets
	fooOutput := `Toot from Micah Lee: Since Elon Musk acquired Twitter on October 27, 2022, Semiphemeral has deleted 7.7M tweets, 4.6M retweets, 13.2M likes, and 3.5M direct messages from Twitter`
	fooURL := `https://ricearoni.org/notice/APjms6WVgEKbZ2MSsi`
	barOutput := `Toot from John Watson: hey y'all i made this for #Godot in case you want to be awesome https://github.com/jotson/ridiculous_coding`
	barURL := `https://mastodon.gamedev.place/@jotson/109367069016579141`
	bazOutput := `Toot from John Watson: hey y'all i made this for #Godot in case you want to be awesome https://github.com/jotson/ridiculous_coding`
	bazURL := `https://fosstodon.org/@jotson@mastodon.gamedev.place/109367069171884095`
	quuxOutput := `Toot from Andrew Nadeau: [after leaving willy wonka’s factory]
me:
wife:
me:
wife:
me:
wife: lot of deaths for a to—
me: a LOT of deaths for a tour!`
	quuxURL := `https://mastodon.social/@AndrewNadeau/109361740736612801`
	fredOutput := `Toot from yan: nobody:

absolutely nobody:

yubikey: cccjgjgkhcbbcvchfkfhiiuunbtnvgihdfiktncvlhck`
	fredURL := `https://infosec.exchange/@bcrypt/109341144608902590`
	thudOutput := `Toot from Tom Grochowiak: Just realised that since I'm new here, I haven't yet spammed my fave (and most useless) project this year.

A genetic algorithm that attempts to give birth to a designated image.

#gamedev #procedural`
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
			input:         fooURL,
			output:        fooOutput,
			expectedError: nil,
		}, {
			input:         "wow check out this tweet " + fooURL,
			output:        fooOutput,
			expectedError: nil,
		}, {
			input:         "wow check out this tweet" + fooURL,
			output:        fooOutput,
			expectedError: nil,
		}, {
			input:         "wow check out this tweet " + fooURL + " super cool right?",
			output:        fooOutput,
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
	for i, c := range cases {
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
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			// these CANNOT run concurrently
			got, err := expandToots(&testingCmd)
			want := c.output
			if err != nil && err.Error() != c.expectedError.Error() {
				t.Error(err)
			}
			if got != want {
				t.Errorf("\ngot %+v\nwant %+v", got, want)
			}
		})
	}
}
