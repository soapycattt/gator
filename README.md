# Introduction

Fun projects to help me getting started at Golang. gator is a REPL program to scrape RSS feeds to local and display to the terminal interface

# Installation

### Manually

Clone this repo and run the following command in the cloned dir

```shell
  ./gator
```

### With go package manager

Or install with go

```shell
  go install github.com/soapycattt/gator
```

# Commands

- `register`: Register a user
- `login`: Login into a user
- `remove`: Remove a user
- `users` / `list`: List all the registered users
- `addfeed`: Add a new feed url
- `feeds`: Show all added feeds
- `follow`: Follow a feed
- `unfollow`: Unfollow a feed
- `following`: List all user-feed pairs
- `browse`: List posts from feed that the user has followed
