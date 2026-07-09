# go-ping-reddit

A minimal Reddit + Go proof of concept. Connects to Reddit with a script app, fetches recent posts from a subreddit, and prints them to stdout.

## Prerequisites

- Go 1.22+
- A Reddit account
- A Reddit **script** app registered at [reddit.com/prefs/apps](https://www.reddit.com/prefs/apps)

When creating the app, choose **script** as the type and use `http://localhost:8080` as the redirect URI (required by Reddit even for script apps).

## Development setup

```bash
make setup   # download deps, build binary, create reddit-account.agent template
make demo    # offline demo — no Reddit credentials required
```

`make setup` creates `bin/go-ping-reddit` and copies `reddit-account.agent.example` to `reddit-account.agent` if it does not exist yet.

## Demo (no credentials)

Run the offline demo to see sample output without calling the Reddit API:

```bash
make demo
# or
go run . -demo
go run . -demo -sub programming -limit 3
```

## Live mode (credentials required)

1. Copy the credentials template if you have not run `make setup`:

   ```bash
   cp reddit-account.agent.example reddit-account.agent
   ```

2. Edit `reddit-account.agent` with your Reddit app credentials and account details.

   The `user_agent` must follow Reddit's format:

   ```
   <platform>:<app ID>:<version> (by /u/<reddit username>)
   ```

   If your account uses 2FA, append the TOTP code to your password with a colon: `password:123456`.

3. Run against the live API:

   ```bash
   make run
   # or
   go run .
   ```

   By default this prints the 5 most recent posts from `/r/golang`.

## Options

```bash
go run . -sub programming -limit 10
go run . -agent /path/to/reddit-account.agent
```

| Flag | Default | Description |
|------|---------|-------------|
| `-sub` | `golang` | Subreddit name (without `/r/`) |
| `-limit` | `5` | Number of posts to print |
| `-agent` | `reddit-account.agent` | Path to credentials file |
| `-demo` | `false` | Print sample posts without calling Reddit |

## Build

```bash
go build -o go-ping-reddit .
./go-ping-reddit -sub golang
```

## How it works

This PoC uses [graw](https://github.com/turnage/graw) with Reddit's script-app (password grant) authentication. Credentials live in a local agent file and are not committed to git.

## Next steps

- Subscribe to subreddit events with graw's bot framework
- Reply to posts or comments
- Deploy as a background worker (e.g. on Render)
