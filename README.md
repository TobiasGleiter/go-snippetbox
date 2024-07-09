# Snippetbox

Snippetbox, which lets people paste and share snippets of text

(From the book of "Let's Go!" by Alex Edwards)

## Cookbook

### Command line flags

`go run ./cmd/web -help` allows you to soo a list of all available command line flags

Hint: Use environment variables and flags together:

`export SNIPPETBOX_ADDR=":9999"` and `go run ./cmd/web -addr=$SNIPPETBOX_ADDR`

### Logs

Redirect the standard out stream to an on-disk file when starting the application: `go run ./cmd/web >>/tmp/web.log`

Notes:

1. `>>` will append to the exitsting file instead of truncating it.
2. `slog.New()` is concurrency safe and can be shared
3. Be careful if multiple loggers try to write the same log file.
