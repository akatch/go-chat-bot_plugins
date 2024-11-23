package bsky

import (
	"errors"
	"testing"

	"github.com/go-chat-bot/bot"
	. "github.com/smartystreets/goconvey/convey"
)

func TestBsky(t *testing.T) {
	var cases = []struct {
		input, output string
		expectedError error
	}{
		{
			input:         "this message has no links",
			output:        "",
			expectedError: nil,
		}, {
			input:         "https://bsky.app/profile/jaredlholt.bsky.social/post/3lbfojkcj4c2s",
			output:        "Jared Holt: Checked in on how the lawsuit against Minnesota for its law banning certain \"deep fake\" content around elections was going and the plaintiff's counsel is alleging that an expert witness offered by the state cited an academic article that doesn't exist and that the citation may be an AI hallucination",
			expectedError: nil,
		}, {
			input:         "wow check out this bsky https://bsky.app/profile/jaredlholt.bsky.social/post/3lbfojkcj4c2s super cool right?",
			output:        "Jared Holt: Checked in on how the lawsuit against Minnesota for its law banning certain \"deep fake\" content around elections was going and the plaintiff's counsel is alleging that an expert witness offered by the state cited an academic article that doesn't exist and that the citation may be an AI hallucination",
			expectedError: nil,
		}, {
			input:         "https://bsky.app/profile/fake.bsky.social/post/3lbfojkcj4c2s",
			output:        "",
			expectedError: errors.New("400 Bad Request"),
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
				got, err := expandPosts(&testingCmd)
				want := c.output

				So(err, ShouldResemble, c.expectedError)
				So(got, ShouldEqual, want)
			})
		}
	},
	)
}
