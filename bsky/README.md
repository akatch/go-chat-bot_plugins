# Overview

The Bsky plugin scrapes message text for Bluesky URLs, then attempts to fetch the linked status and post it to the channel, like so:

```
08:29:19 <user> https://bsky.app/profile/jaredlholt.bsky.social/post/3lbfojkcj4c2s
08:29:21 <chatbot> Jared Holt: Checked in on how the lawsuit against Minnesota for its law banning certain \"deep fake\" content around elections was going and the plaintiff's counsel is alleging that an expert witness offered by the state cited an academic article that doesn't exist and that the citation may be an AI hallucination,
```
