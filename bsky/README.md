# Overview

The Bsky plugin scrapes message text for Bluesky URLs, then attempts to fetch the linked status and post it to the channel, like so:

```
08:29:19 <user> https://bsky.app/profile/jaredlholt.bsky.social/post/3lbfojkcj4c2s
08:29:21 <chatbot> Jared Holt: Checked in on how the lawsuit against Minnesota for its law banning certain \"deep fake\" content around elections was going and the plaintiff's counsel is alleging that an expert witness offered by the state cited an academic article that doesn't exist and that the citation may be an AI hallucination,
```

# Test API calls

given a post URL `https://bsky.app/profile/did:plc:x5ii2ldzprnrz4szq7kdbao6/post/3mnuspmkvxc2a`

obtain the user's handle then substitute

`at://$HANDLE.bsky.social/app.bsky.feed.post/$POST`

the post above becomes `at://brittanyepage.bsky.social/app.bsky.feed.post/3mnuspmkvxc2a`

then you can do

```
curl -sqG \
    --data-urlencode "uri=at://jaredlholt.bsky.social/app.bsky.feed.post/3lbfojkcj4c2s" \
    --data-urlencode depth=0 \
    --data-urlencode parentHeight=0 \
    https://public.api.bsky.app/xrpc/app.bsky.feed.getPostThread \
    | jq '"\(.thread.post.author.displayName): \(.thread.post.record.text) [image: \(.thread.post.embed.images[].alt)]"'
```
