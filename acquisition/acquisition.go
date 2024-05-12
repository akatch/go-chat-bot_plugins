package acquisition

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/go-chat-bot/bot"
)

const (
	msgInvalidAmountOfParams = "Invalid amount of parameters"
	msgInvalidParam          = "Invalid parameter"
)

type ruleOfAcquisition struct {
	Number int    // The number Quark would use to refer to a Rule
	Rule   string // The text of the Rule
}

// String provides a string representation of the Rule, or a message
// if the rule is unavailable.
func (r *ruleOfAcquisition) String() string {
	// If the rule is empty, return an alternative message
	if r.Rule == "" {
		return fmt.Sprintf("Access to Rule of Acquisition %d is restricted. Please contact your local Ferengi Commerce Authority liquidator for more information.", r.Number)
	}
	return fmt.Sprintf("Rule of Acquisition %d: %s", r.Number, r.Rule)
}

// getRule will check the validity of the provided index and, if valid,
// will return the corresponding Rule of Acquisition. If the index is not
// valid, getRule will return an error
func getRule(index int) (ruleOfAcquisition, error) {
	if index > len(rules) || index <= 0 {
		return ruleOfAcquisition{}, errors.New(msgInvalidParam)
	}
	return ruleOfAcquisition{index, rules[index-1]}, nil
}

// acquisition invokes an instance of the command to return either a random
// or indexed Rule of Acquisition.
func acquisition(command *bot.Cmd) (string, error) {
	var index int
	var err error
	switch len(command.Args) {
	case 0:
		index = randRule()
	case 1:
		index, err = strconv.Atoi(command.Args[0])
		if err != nil {
			return "", err
		}
	default:
		return "", errors.New(msgInvalidAmountOfParams)
	}

	r, err := getRule(index)
	if err != nil {
		return "", err
	}

	return r.String(), nil
}

// randRule returns a random integer between 1 and 286, inclusive
func randRule() int {
	return rand.Intn(len(rules))
}

// init initializes the command
func init() {
	bot.RegisterCommand(
		"acquisition",
		"display the Ferengi Rules of Acquisition",
		"",
		acquisition)
}
