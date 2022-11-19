// Package mastodon provides a plugin that scrapes messages for Mastodon links,
// then expands them into chat messages.
package mastodon

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/go-chat-bot/bot"
	"github.com/mattn/go-mastodon"
	"jaytaylor.com/html2text"
)

type Status struct {
	server, id string
}

// findStatuses checks a given message string for strings that look like Twitter links,
// then attempts to extract the Tweet ID from the link.
// It returns an array of Tweet IDs.
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

// fetchTweets takes an array of Tweet IDs and retrieves the corresponding
// Statuses.
// It returns an array of twitter.Tweets.
func fetchTweets(tweetIDs []Status) ([]mastodon.Status, error) {
	var tweets []mastodon.Status
	for _, tweetID := range tweetIDs {
		tweet, err := fetchTweet(tweetID)
		if err != nil {
			return nil, err
		}
		tweets = append(tweets, *tweet)
	}
	return tweets, nil
}

// fetchTweet takes a twitter.Client and a single Tweet ID and fetches the
// corresponding Status.
// It returns a twitter.Tweet.
func fetchTweet(tweetID Status) (*mastodon.Status, error) {
	// credentials are not needed for mastodon
	client := mastodon.NewClient(&mastodon.Config{Server: tweetID.server})
	tweet, err := client.GetStatus(context.Background(), mastodon.ID(tweetID.id))

	return tweet, err
}

// formatTweets takes an array of twitter.Tweets and formats them in preparation for
// sending as a chat message.
// It returns an array of nicely formatted strings.
func formatTweets(tweets []mastodon.Status) []string {
	formatString := "Toot from %s: %s"
	//newlines := regexp.MustCompile(`\r?\n`)
	var messages []string
	for _, tweet := range tweets {
		// TODO get link title, eg: Tweet from @user: look at this cool thing https://thing.cool (Link title: A Cool Thing)
		// TODO get alt text
		username := tweet.Account.DisplayName
		pt, err := html2text.FromString(tweet.Content, html2text.Options{TextOnly: true, OmitLinks: true})
		if err != nil {
			return nil
		}
		//text := newlines.ReplaceAllString(pt, " ")
		newMessage := fmt.Sprintf(formatString, username, pt)
		messages = append(messages, newMessage)
	}
	return messages
}

// expandTweets receives a bot.PassiveCmd and performs the full parse-and-fetch
// pipeline. It sets up a client, finds Tweet IDs in the message text, fetches
// the tweets, and formats them. If multiple Tweet IDs were found in the message,
// all formatted Tweets will be joined into a single message.
// It returns a single string suitable for sending as a chat message.
func expandToots(cmd *bot.PassiveCmd) (string, error) {
	var message string
	messageText := cmd.MessageData.Text

	tweetIDs, err := findStatuses(messageText)
	if err != nil {
		return message, err
	}

	statuses, err := fetchTweets(tweetIDs)
	if err != nil {
		return message, err
	}

	formattedStatuses := formatTweets(statuses)
	if formattedStatuses != nil {
		message = strings.Join(formattedStatuses, "\n")
	}
	return message, err
}

// init initalizes a PassiveCommand for expanding Toots.
func init() {
	bot.RegisterPassiveCommand(
		"mastodon",
		expandToots)
}
