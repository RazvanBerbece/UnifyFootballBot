# UnifyFootballBot
The extensible, fast and reliable Discord bot which powers the UnifyFootball.ro online community. Written in Go.

# How to Run
## On host machine
1. Build the project by running the `go build -o build ./cmd/...` command at the root of the repo.
2. Start the bot application by running the `./build/bot-srv` command at the root of the repo.

## In Docker
1. Build the container image by running the `docker build --tag unify-bot-app .` command at the root of the repo.
2. Run the container with the `docker run unify-bot-app` command at the root of the repo.

# How to Test
1. Run all the tests for the project by running the `go test ./internal/...` command at the root of the repo.
