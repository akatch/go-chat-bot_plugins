// Package fediverse provides a plugin that scrapes messages for Mastodon links,
// then expands them into chat messages.
package fediverse

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/go-chat-bot/bot"
	"github.com/jaytaylor/html2text"
	"github.com/mattn/go-mastodon"
)

type Status struct {
	server, id string
}

// findStatuses checks a given message string for strings that look like Mastodon links,
// then attempts to extract the server and status ID from the link.
// It returns an array of Statuses.
func findStatuses(message string) ([]Status, error) {
	re := regexp.MustCompile(`(http(?:s)?://[A-z0-9\.\-\/]+[^/])/(?:@.*@(.*)?)?(?:.*)/(statuses/|notice/)?([A-z0-9]*)`)
	result := re.FindAllStringSubmatch(message, -1)
	var (
		statuses []Status
		err      error
	)

	// parse server URL and status ID from `result`
	// server URL is the right-most non-empty DNS name match
	for i := range result {
		last := len(result[i]) - 1
		third := len(result[i]) - 2
		svr := result[i][third]
		if svr == "" {
			// if penultimate capture is empty, the first capture will be the server URL
			svr = result[i][1]
		}
		idStr := result[i][last]
		statuses = append(statuses, Status{server: svr, id: idStr})
	}

	return statuses, err
}

// fetchStatuses takes an array of Status IDs and retrieves the corresponding
// Statuses.
// It returns an array of Statuses.
func fetchStatuses(s []Status) ([]mastodon.Status, error) {
	var statuses []mastodon.Status
	for _, id := range s {
		st, err := fetchStatus(id)
		if err != nil {
			return nil, err
		}
		statuses = append(statuses, *st)
	}
	return statuses, nil
}

// fetchStatus takes a single Status ID and returns the
// corresponding Status.
func fetchStatus(s Status) (*mastodon.Status, error) {
	// credentials are not needed for fediverse
	client := mastodon.NewClient(&mastodon.Config{Server: s.server})
	status, err := client.GetStatus(context.Background(), mastodon.ID(s.id))

	return status, err
}

// formatStatuses takes an array of Statuses and formats them in preparation for
// sending as a chat message.
// It returns an array of nicely formatted strings.
func formatStatuses(s []mastodon.Status) []string {
	formatString := "%s: %s"
	newlines := regexp.MustCompile(`\r?\n`)
	var messages []string
	for _, st := range s {
		// TODO get link title, eg: user: look at this cool thing https://thing.cool (Link title: A Cool Thing)
		// TODO get alt text
		username := st.Account.DisplayName
		pt, err := html2text.FromString(st.Content, html2text.Options{TextOnly: true, OmitLinks: true})
		if err != nil {
			return nil
		}
		text := newlines.ReplaceAllString(pt, " ")
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

	ids, err := findStatuses(messageText)
	if err != nil {
		return message, err
	}

	statuses, err := fetchStatuses(ids)
	if err != nil {
		return message, err
	}

	formattedStatuses := formatStatuses(statuses)
	if formattedStatuses != nil {
		message = strings.Join(formattedStatuses, "\n")
	}
	return message, err
}

// init initalizes a PassiveCommand for expanding Toots.
func init() {
	bot.RegisterPassiveCommand(
		"fediverse",
		expandStatuses)
}
