package acquisition

import (
	"testing"

	"github.com/go-chat-bot/bot"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAcquisition(t *testing.T) {
	Convey("acquisition", t, func() {
		bot := &bot.Cmd{
			Command: "acquisition",
		}

		Convey("List should contain 286 Rules", func() {
			So(len(rules), ShouldEqual, 286)
		})

		Convey("Should return a random Rule", func() {
			got, error := acquisition(bot)

			So(error, ShouldBeNil)
			So(got, ShouldNotBeBlank)
		})

		Convey("Should return Rule 286", func() {
			bot.Args = []string{"286"}
			got, error := acquisition(bot)

			So(error, ShouldBeNil)
			So(got, ShouldEqual, "Rule of Acquisition 286: When Morn leaves, it's all over.")
		})

		Convey("Should return the first Rule of Acuisition", func() {
			bot.Args = []string{"1"}
			got, error := acquisition(bot)

			So(error, ShouldBeNil)
			So(got, ShouldEqual, "Rule of Acquisition 1: Once you have their money, never give it back.")
		})

		Convey("Should return a error message when pass a too-large index", func() {
			bot.Args = []string{"300"}
			got, error := acquisition(bot)

			So(got, ShouldBeBlank)
			So(error.Error(), ShouldEqual, msgInvalidParam)
		})

		Convey("Should return a error message when pass a zero index", func() {
			bot.Args = []string{"0"}
			got, error := acquisition(bot)

			So(got, ShouldBeBlank)
			So(error.Error(), ShouldEqual, msgInvalidParam)
		})

		Convey("Should return a error message when pass a invalid amount of params", func() {
			bot.Args = []string{"1", "2"}
			got, error := acquisition(bot)

			So(got, ShouldBeBlank)
			So(error.Error(), ShouldEqual, msgInvalidAmountOfParams)
		})
		Convey("Empty rules should return a special message", func() {
			bot.Args = []string{"284"}
			got, error := acquisition(bot)

			So(error, ShouldBeNil)
			So(got, ShouldEqual, "Access to Rule of Acquisition 284 is restricted. Please contact your local Ferengi Commerce Authority liquidator for more information.")
		})
	})
}
