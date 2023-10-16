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

# CI/CD
## CI
Continuous integration is implemented via a workflow leveraging Github Actions. The workflow uses Go's build and testing CLI to run the
automated test harness against various versions of Go. The source file for the CI flow can be seen in [test.yml](.github/workflows/test.yml)

### Run CI Locally
I took the opportunity to try out Act (https://github.com/nektos/act) for this project and to my surprise it was really easy to setup and use. In order to be able to run the workflows locally, one must have Act installed on the host machine (check the guide here https://github.com/nektos/act#installation). 

#### Steps to run the workflows using Act:
1. In dry-run mode (don't create any containers and will validate workflow file syntax and correctness)
```shell
act -n -s GH_TOKEN=github_token_here -s DISCORD_BOT_TOKEN=discord_token_here -s APP_ID=app_id_here
```

2. Or not in dry-run mode (just drop the `-n` flag)
```shell
act -s GH_TOKEN=github_token_here -s DISCORD_BOT_TOKEN=discord_token_here -s APP_ID=app_id_here
```

# Credits
fhatti (https://github.com/fhatti) for coming up with the feature ideas.