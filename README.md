# LeRT (**Le**mmy **R**eddit **Translation** Layer)

This project aims to translate the Lemmy API to the Reddit API, such that existing Reddit clients can be easily made compatible with Lemmy.

It is the same effor as [Tafkars](https://github.com/derivator/tafkars/tree/main) but written in Go (as I'm unfamiliar with Rust) and much less featureful (currently).

## Contributing

If you want to work on mappings, please raise an issue with the feature of the Reddit API you want to emulate before starting, this is to avoid duplicate work.

The biggest help right now, would be building up a test suite of to ensure compatibility and reproducibility.

## Status

This is a very early iteration of the software, it's heavily reliant on packages for the Reddit API and the Lemmy API. It is currently intended to support **read only** access.

## Intentions

The intended functionality for the translation layer is to expose an OAuth endpoint that authenticates the user to their Lemmy homeserver and then returns the JWT from Lemmy as the Access Token.

I'm unsure if this works, if there's someone more familiar with how clients are consuming the API (ie: if they're reading specific contents of the Reddit token_v2 or relying on the older Bearer token) please let me know and I'll work an alternative (worst case, Lemmy authorization will have to be stored locally in the translation layer)

There will be things that don't translate perfectly, but, the aim is to at least reach support of:

- [ ] User specific frontpage
- [ ] Single posts and comments
- [ ] Specific "subreddit" browsing and filtering (at least Hot and Top by time)


## Mappings

There are likely aspects I'm missing here, please raise a git issue for additional items that need more complex translations

### Subreddit names

This mapping will likely break some assumptions that Reddit clients make about Subreddit names returned. As we'll use the <Community>@<Instance> format, clients recieving the "@" or names exceeding traditional subreddit length may break.

A future feature may be a "compatability mode" which removes the "@<Instance>" but restricts the clients to only browsing local communities (this could be useful for Bots etc that consume the API, instead of clients, as you'd want to restrict them to local instances regardless) 


### Fullnames

Reddit uses a system called "Full names" for comments, posts, etc which use a prefix (t1_ through to t6_) followed by 6 chars of base36.

To translate this, we'll map the prefix appropriately and then convert the integer representation from Lemmy to base36, and left-pad it be 6 chars with "0".

