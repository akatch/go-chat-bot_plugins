// Package bsky provides a plugin that scrapes messages for Bluesky links,
// then expands them into chat messages.
package bsky

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/akatch/go-chat-bot_plugins/web"
	"github.com/go-chat-bot/bot"
)

const (
	apiUrl = "https://public.api.bsky.app/xrpc/app.bsky.feed.getPostThread"
)

type Post struct {
	handle, id, name, text string
}

// findPosts checks a given message string for strings that look like Bluesky links,
// then attempts to extract the server and post ID from the link.
// It returns an array of Posts.
func findPosts(message string) ([]Post, error) {
	re := regexp.MustCompile(`http(?:s)?://bsky\.app/profile/?(.*)/post/?([A-z0-9]*)`)
	result := re.FindAllStringSubmatch(message, -1)
	var (
		posts []Post
		err   error
	)

	// parse server URL and post ID from `result`
	// server URL is the right-most non-empty DNS name match
	for i := range result {
		posts = append(posts, Post{handle: result[i][1], id: result[i][2]})
	}

	return posts, err
}

// fetchPosts takes an array of Post IDs and retrieves the corresponding
// Posts.
// It returns an array of Posts.
func fetchPosts(s []Post) ([]Post, error) {
	var posts []Post
	for _, id := range s {
		p, err := fetchPost(&id)
		if err != nil {
			return nil, err
		}
		posts = append(posts, *p)
	}
	return posts, nil
}

// fetchPost takes a single Post ID and returns the
// corresponding Post.
func fetchPost(post *Post) (*Post, error) {
	response := map[string]interface{}{}
	err := web.GetJSON(fmt.Sprintf(`%s?uri=at://%s/app.bsky.feed.post/%s&depth=0&parentHeight=0`, apiUrl, post.handle, post.id), &response)

	if err != nil {
		return post, err
	}

	post = &Post{
		handle: post.handle,
		id:     post.id,
		name:   fmt.Sprint(response["thread"].(map[string]interface{})["post"].(map[string]interface{})["author"].(map[string]interface{})["displayName"]),
		text:   fmt.Sprint(response["thread"].(map[string]interface{})["post"].(map[string]interface{})["record"].(map[string]interface{})["text"])}

	return post, err
}

// formatPosts takes an array of Posts and formats them in preparation for
// sending as a chat message.
// It returns an array of nicely formatted strings.
func formatPosts(posts []Post) []string {
	formatString := "%s: %s"
	newlines := regexp.MustCompile(`\r?\n`)
	var messages []string
	for _, post := range posts {
		// TODO include alt text thread.post.embed.images[].alt
		newMessage := fmt.Sprintf(formatString, post.name, newlines.ReplaceAllString(post.text, " "))
		messages = append(messages, newMessage)
	}
	return messages
}

// expandPosts receives a bot.PassiveCmd and performs the full parse-and-fetch
// pipeline. It sets up a client, finds Post IDs in the message text, fetches
// the posts, and formats them. If multiple Post IDs were found in the message,
// all formatted Posts will be joined into a single message.
// It returns a single string suitable for sending as a chat message.
func expandPosts(cmd *bot.PassiveCmd) (string, error) {
	var message string
	messageText := cmd.MessageData.Text

	ids, err := findPosts(messageText)
	if err != nil {
		return message, err
	}

	posts, err := fetchPosts(ids)
	if err != nil {
		return message, err
	}

	formattedPosts := formatPosts(posts)
	if formattedPosts != nil {
		message = strings.Join(formattedPosts, "\n")
	}
	return message, err
}

// init initalizes a PassiveCommand for expanding Toots.
func init() {
	bot.RegisterPassiveCommand(
		"bsky",
		expandPosts)
}
