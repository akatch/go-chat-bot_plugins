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
			input: "https://bsky.app/profile/jaredlholt.bsky.social/post/3lbfojkcj4c2s",
			output: `Jared Holt: Checked in on how the lawsuit against Minnesota for its law banning certain "deep fake" content around elections was going and the plaintiff's counsel is alleging that an expert witness offered by the state cited an academic article that doesn't exist and that the citation may be an AI hallucination [image alt text: legal filing alleges expert testimony in case cited a nonexistent article]
`,
			expectedError: nil,
		}, {
			input: "wow check out this bsky https://bsky.app/profile/jaredlholt.bsky.social/post/3lbfojkcj4c2s super cool right?",
			output: `Jared Holt: Checked in on how the lawsuit against Minnesota for its law banning certain "deep fake" content around elections was going and the plaintiff's counsel is alleging that an expert witness offered by the state cited an academic article that doesn't exist and that the citation may be an AI hallucination [image alt text: legal filing alleges expert testimony in case cited a nonexistent article]
`,
			expectedError: nil,
		}, {
			input:         "https://bsky.app/profile/fake.bsky.social/post/3lbfojkcj4c2s",
			output:        "",
			expectedError: errors.New("400 Bad Request"),
		}, {
			input: "https://bsky.app/profile/andrewtorrez.bsky.social/post/3lbk7mkb5f22x",
			output: `P. Andrew Torrez: 1/ This morning, Judge Liman ordered lawyers for the Freeman/Moss plaintiffs trying to collect on their $148MM defamation judgment against Rudy Giuliani to answer TODAY whether their pending motion to compel discovery from "Standard USA, LLC" is moot.  www.courtlistener.com/docket/69015... [image alt text: ORDER taking under advisement 137 Motion to Compel. Plaintiffs to indicate by letter no later than November 22, 2024, why this motion should not be denied as moot in light of the filing at Dkt. No. 139. (HEREBY ORDERED by Judge Lewis J. Liman)(Text Only Order) (Liman, Lewis) (Entered: 11/22/2024)]
`,
			expectedError: nil,
		}, {
			// has an image but no alt text
			input:         "https://bsky.app/profile/did:plc:b2nvxgipvmqraxogjxaaugzf/post/3mptbr3n6xc2a",
			output:        `David Waldman: See, the plan is, you put 1,000 foreign women who all got pregnant at the same time on a hospital ship, and you do that 1,000 times, and then you wait 18 years, and then you move all the babies to Wyoming, and you win a House seat. `,
			expectedError: nil,
		}, /*j  {
			input:         "what https://bsky.app/profile/did:plc:tpcydhxibth3qbp2o3g4u37n/post/3mpzbj2xcys2w",
			output:        "asdf",
			expectedError: nil,
		}, */
	}

	//SetDefaultStackMode(StackError)
	Convey("bsky", t, func() {
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
