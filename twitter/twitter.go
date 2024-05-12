// Package twitter provides a plugin that scrapes messages for Twitter links,
// then expands them into chat messages.
package twitter

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-chat-bot/bot"
)

const (
	apiUrl = "https://api.fxtwitter.com"
)

type Status struct {
	Name string `json:"tweet.author.name"`
	Text string `json:"tweet.text"`
}

// findTweetIDs checks a given message string for strings that look like Twitter links,
// then attempts to extract the Tweet ID from the link.
// It returns an array of Tweet IDs.
func findTweetIDs(message string) ([]int64, error) {
	re := regexp.MustCompile(`http(?:s)?://(?:mobile.)?twitter.com/(?:.*)/status/([0-9]*)`)
	// FIXME this is only returning the LAST match, should return ALL matches
	result := re.FindAllStringSubmatch(message, -1)
	var (
		tweetIDs []int64
		id       int64
		err      error
	)

	for i := range result {
		last := len(result[i]) - 1
		idStr := result[i][last]
		id, err = strconv.ParseInt(idStr, 10, 64)
		tweetIDs = append(tweetIDs, id)
	}
	return tweetIDs, err
}

// fetchTweets takes an array of Tweet IDs and retrieves the corresponding
// Statuses.
// It returns an array of twitter.Tweets.
func fetchTweets(tweetIDs []int64) ([]Status, error) {
	var tweets []Status
	for _, tweetID := range tweetIDs {
		tweet, err := fetchTweet(tweetID)
		if err != nil {
			return nil, err
		}
		tweets = append(tweets, tweet)
	}
	return tweets, nil
}

// fetchTweet takes a twitter.Client and a single Tweet ID and fetches the
// corresponding Status.
// It returns a twitter.Tweet.
func fetchTweet(tweetID int64) (Status, error) {
	var err error
	var status Status
	// TODO get alt text
	// tweet.media.photos[].altText

	resp, err := http.Get(fmt.Sprintf("%s/status/%d", apiUrl, tweetID))

	// If we return nil instead of tweet, a panic happens
	if resp.StatusCode/200 != 1 {
		return status, errors.New(resp.Status)
	} else if err != nil {
		return status, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return Status{}, err
	}

	decodedResponse := map[string]interface{}{}
	err = json.Unmarshal(body, &decodedResponse)

	if err != nil {
		return Status{}, err
	}

	status = Status{
		Name: fmt.Sprint(decodedResponse["tweet"].(map[string]interface{})["author"].(map[string]interface{})["name"]),
		Text: fmt.Sprint(decodedResponse["tweet"].(map[string]interface{})["text"])}

	return status, nil
}

// formatTweets takes an array of twitter.Tweets and formats them in preparation for
// sending as a chat message.
// It returns an array of nicely formatted strings.
func formatTweets(tweets []Status) []string {
	formatString := "%s: %s"
	newlines := regexp.MustCompile(`\r?\n`)
	var messages []string
	for _, tweet := range tweets {
		// TODO get link title, eg: User: look at this cool thing https://thing.cool (Link title: A Cool Thing)
		// tweet.Entities.Urls contains []URLEntity
		// fetch title from urlEntity.URL
		// urls plugin already correctly handles t.co links
		username := tweet.Name
		text := newlines.ReplaceAllString(tweet.Text, " ")
		newMessage := fmt.Sprintf(formatString, username, text)
		messages = append(messages, newMessage)
	}
	return messages
}

// expandTweets receives a bot.PassiveCmd and performs the full parse-and-fetch
// pipeline. It sets up a client, finds Tweet IDs in the message text, fetches
// the tweets, and formats them. If multiple Tweet IDs were found in the message,
// all formatted Tweets will be joined into a single message.
// It returns a single string suitable for sending as a chat message.
func expandTweets(cmd *bot.PassiveCmd) (string, error) {
	var message string
	messageText := cmd.MessageData.Text

	tweetIDs, err := findTweetIDs(messageText)
	if err != nil {
		return message, err
	}

	tweets, err := fetchTweets(tweetIDs)
	if err != nil {
		return message, err
	}

	formattedTweets := formatTweets(tweets)
	if formattedTweets != nil {
		message = strings.Join(formattedTweets, "\n")
	}
	return message, err
}

// init initalizes a PassiveCommand for expanding Tweets.
func init() {
	bot.RegisterPassiveCommand(
		"twitter",
		expandTweets)
}
