// Package twitter provides a plugin that scrapes messages for Twitter links,
// then expands them into chat messages.
package twitter

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/akatch/go-chat-bot_plugins/web"
	"github.com/go-chat-bot/bot"
)

const (
	apiUrl = "https://api.fxtwitter.com"
)

type Status struct {
	Name string `json:"tweet.author.name"`
	Text string `json:"tweet.text"`
}

// findStatusIDs checks a given message string for strings that look like Twitter links,
// then attempts to extract the Status ID from the link.
// It returns an array of Status IDs.
func findStatusIDs(message string) ([]int64, error) {
	re := regexp.MustCompile(`http(?:s)?://((?:mobile.)?twitter|x).com/(?:.*)/status/([0-9]*)`)
	// FIXME this is only returning the LAST match, should return ALL matches
	result := re.FindAllStringSubmatch(message, -1)
	var (
		statusIDs []int64
		id        int64
		err       error
	)

	for i := range result {
		last := len(result[i]) - 1
		idStr := result[i][last]
		id, err = strconv.ParseInt(idStr, 10, 64)
		statusIDs = append(statusIDs, id)
	}
	return statusIDs, err
}

// fetchStatuses takes an array of Status IDs and retrieves the corresponding
// Statuses.
// It returns an array of twitter.Statuses.
func fetchStatuses(statusIDs []int64) ([]Status, error) {
	var statuses []Status
	for _, statusID := range statusIDs {
		status, err := fetchStatus(statusID)
		if err != nil {
			return nil, err
		}
		statuses = append(statuses, status)
	}
	return statuses, nil
}

// fetchStatus takes a twitter.Client and a single Status ID and fetches the
// corresponding Status.
// It returns a twitter.Status.
func fetchStatus(id int64) (Status, error) {
	// TODO get alt text
	// tweet.media.photos[].altText
	response := map[string]interface{}{}
	err := web.GetJSON(fmt.Sprintf("%s/status/%d", apiUrl, id), &response)

	if err != nil {
		return Status{}, err
	}

	status := Status{
		Name: fmt.Sprint(response["tweet"].(map[string]interface{})["author"].(map[string]interface{})["name"]),
		Text: fmt.Sprint(response["tweet"].(map[string]interface{})["text"])}

	return status, nil
}

// formatStatuses takes an array of twitter.Statuses and formats them in preparation for
// sending as a chat message.
// It returns an array of nicely formatted strings.
func formatStatuses(statuses []Status) []string {
	formatString := "%s: %s"
	newlines := regexp.MustCompile(`\r?\n`)
	var messages []string
	for _, status := range statuses {
		// TODO get link title, eg: User: look at this cool thing https://thing.cool (Link title: A Cool Thing)
		// urls plugin already correctly handles t.co links
		username := status.Name
		text := newlines.ReplaceAllString(status.Text, " ")
		newMessage := fmt.Sprintf(formatString, username, text)
		messages = append(messages, newMessage)
	}
	return messages
}

// expandStatuses receives a bot.PassiveCmd and performs the full parse-and-fetch
// pipeline. It sets up a client, finds Status IDs in the message text, fetches
// the statuses, and formats them. If multiple Status IDs were found in the message,
// all formatted Statuses will be joined into a single message.
// It returns a single string suitable for sending as a chat message.
func expandStatuses(cmd *bot.PassiveCmd) (string, error) {
	var message string
	messageText := cmd.MessageData.Text

	statusIDs, err := findStatusIDs(messageText)
	if err != nil {
		return message, err
	}

	statuses, err := fetchStatuses(statusIDs)
	if err != nil {
		return message, err
	}

	formattedStatuses := formatStatuses(statuses)
	if formattedStatuses != nil {
		message = strings.Join(formattedStatuses, "\n")
	}
	return message, err
}

// init initalizes a PassiveCommand for expanding Statuses.
func init() {
	bot.RegisterPassiveCommand(
		"twitter",
		expandStatuses)
}
