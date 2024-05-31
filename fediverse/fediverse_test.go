package fediverse

import (
    "errors"
    "testing"

    "github.com/go-chat-bot/bot"
    . "github.com/smartystreets/goconvey/convey"
)

func TestMastodon(t *testing.T) {
    barOutput := `John Watson: hey y'all i made this for #Godot <https://mastodon.gamedev.place/tags/Godot> in case you want to be awesome https://github.com/jotson/ridiculous_coding <https://github.com/jotson/ridiculous_coding>`
    barURL := `https://mastodon.gamedev.place/@jotson/109367069016579141`
    bazOutput := `John Watson: hey y'all i made this for #Godot <https://mastodon.gamedev.place/tags/Godot> in case you want to be awesome https://github.com/jotson/ridiculous_coding <https://github.com/jotson/ridiculous_coding>`
    bazURL := `https://fosstodon.org/@jotson@mastodon.gamedev.place/109367069171884095`
    quuxOutput := `Andrew Nadeau: [after leaving willy wonka’s factory] me: wife: me: wife: me: wife: lot of deaths for a to— me: a LOT of deaths for a tour!`
    quuxURL := `https://mastodon.social/@AndrewNadeau/109361740736612801`
    fredOutput := `yan: nobody:  absolutely nobody:  yubikey: cccjgjgkhcbbcvchfkfhiiuunbtnvgihdfiktncvlhck`
    fredURL := `https://infosec.exchange/@bcrypt/109341144608902590`
    thudOutput := `Tom Grochowiak: Just realised that since I'm new here, I haven't yet spammed my fave (and most useless) project this year.   A genetic algorithm that attempts to give birth to a designated image.  #gamedev <https://mastodon.gamedev.place/tags/gamedev> #procedural <https://mastodon.gamedev.place/tags/procedural> #procgen <https://mastodon.gamedev.place/tags/procgen>`
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

    //SetDefaultStackMode(StackError)
    Convey("fediverse", t, func() {
        for _, c := range cases {
            testingCmd := bot.PassiveCmd{
                Raw:     c.input,
                Channel: "test",
                User: &bot.User{
                    ID:       "test",
                    Nick:     "test",
                    RealName: "test",
                    IsBot:    true,
                },
                MessageData: &bot.Message{
                    Text:     c.input,
                    IsAction: false,
                },
            }
            Convey(c.input, func() {
                got, err := expandStatuses(&testingCmd)
                want := c.output

                So(err, ShouldResemble, c.expectedError)
                So(got, ShouldEqual, want)
            })
        }
    },
    )
}
