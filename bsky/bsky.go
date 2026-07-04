// Package bsky provides a plugin that scrapes messages for Bluesky links,
// then expands them into chat messages.
package bsky

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/akatch/go-chat-bot_plugins/web"
	"github.com/go-chat-bot/bot"
)

const (
	apiUrl = "https://public.api.bsky.app/xrpc/app.bsky.feed.getPostThread"
)

type Response struct {
	Thread *Thread `json:"thread"`
}

type Thread struct {
	Post *Post `json:"post"`
}

type Post struct {
	ID     string  `json:"cid"`
	Author *Author `json:"author"`
	Record *Record `json:"record"`
}

type Author struct {
	Handle string `json:"handle"`
	Name   string `json:"displayName"`
}

type Record struct {
	Text  string `json:"text"`
	Embed *Embed `json:"embed"`
}

type Embed struct {
	Images []*Image `json:"images"`
}

type Image struct {
	Alt string `json:"alt"`
}

// findPosts checks a given message string for strings that look like Bluesky links,
// then attempts to extract the server and post ID from the link.
// It returns an array of Posts.
func findPosts(message string) ([]Response, error) {
	re := regexp.MustCompile(`http(?:s)?://bsky\.app/profile/?(.*)/post/?([A-z0-9]*)`)
	result := re.FindAllStringSubmatch(message, -1)
	var (
		posts []Response
		err   error
	)

	// parse server URL and post ID from `result`
	// server URL is the right-most non-empty DNS name match
	for i := range result {
		posts = append(
			posts,
			Response{
				Thread: &Thread{
					Post: &Post{
						ID: result[i][2],
						Author: &Author{
							Handle: result[i][1],
						},
					}}})
	}

	return posts, err
}

// fetchPosts takes an array of Post IDs and retrieves the corresponding
// Posts.
// It returns an array of Posts.
func fetchPosts(s []Response) ([]Response, error) {
	var posts []Response
	for _, id := range s {
		p, err := fetchPost(&id)
		if err != nil {
			return nil, err
		}
		posts = append(posts, *p)
	}
	return posts, nil
}

func GetAPIURL(post *Response) string {
	return fmt.Sprintf(
		`%s?uri=at://%s/app.bsky.feed.post/%s&depth=0&parentHeight=0`,
		apiUrl,
		post.Thread.Post.Author.Handle,
		post.Thread.Post.ID)
}

// fetchPost takes a single Post ID and returns the
// corresponding Post.
func fetchPost(post *Response) (*Response, error) {
	resp, err := web.GetBody(GetAPIURL(post))
	if err != nil {
		return nil, err
	}

	var r Response
	err = json.Unmarshal(resp, &r)

	if err != nil {
		return nil, err
	}

	return &r, nil

}

// JSON returns a JSON string representation of the Response
func (post *Response) JSON() string {
	str, err := json.MarshalIndent(post, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(str)
}

// String returns a concise string representation of the Response, including image alt text if present
func (post *Response) String() string {
	newlines := regexp.MustCompile(`\r?\n`)
	name := post.Thread.Post.Author.Name
	text := newlines.ReplaceAllString(post.Thread.Post.Record.Text, " ")

	if len(post.Thread.Post.Record.Embed.Images) > 0 {
		var alts string
		for _, t := range post.Thread.Post.Record.Embed.Images {
			alts += fmt.Sprintf("[image alt text: %s] ", t.Alt)
		}
		return fmt.Sprintf("%s: %s %s", name, text, alts)
	} else {
		return fmt.Sprintf("%s: %s", name, text)
	}

}

// formatPosts takes an array of Posts and formats them in preparation for
// sending as a chat message.
// It returns an array of nicely formatted strings.
func formatPosts(posts []Response) []string {
	var messages []string
	for _, post := range posts {
		newMessage := post.String()
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
