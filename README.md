# UnifyFootballBot
The extensible, fast and reliable Discord bot which powers the UnifyFootball.ro online community. Written in Go.

# How to Run
## Prerequisites
1. Bot has to be added as an app to the Discord server.
2. The `.env` file has to be created in this repo and populated with the correct key-value pairs (`DISCORD_BOT_TOKEN=...`, `APP_ID=...`, etc.)
    - The file [globals.go](internal/globals/globals.go) contains the declarations of all the required environment variables required by the project (i.e all of them have to be in the `.env` file), alongside bits of useful documentation for what values are expected for the variables
3. Have the Go SDK installed.
4. (Optional) Have Docker installed.

## Run the application
1. Run a built full container app composition (app, DBs, etc.) with the `docker-compose up -d --remove-orphans --build` command.
2. Bring down the application by running `docker compose down`.

_Note: In case one needs to run only the Go application code without any associated infrastructure, it is possible by simply running `go build -o build ./cmd/... && ./build/bot-srv`. This however means that the DB operations will not work and thus not all bot features available, unless a local non-containerised MySQL server is used._

# How to Test
1. Run all the tests for the project by running the `go test ./internal/...` command at the root of the repo.

# CI/CD
## CI
Continuous integration is implemented via a workflow leveraging Github Actions. The workflow uses Go's build and testing CLI to run the
automated test harness against various versions of Go. The source file for the CI flow can be seen in [test.yml](.github/workflows/test.yml).

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

## CD
Continuous deployment has not been implemented yet, as this is currently a POC. But it will leverage GitHub Actions the same way the CI step does, either by pushing packages to something like Octopus, or by directly deploying to a Cloud provider.

# Credits
fhatti (https://github.com/fhatti) for coming up with the project in the first place and the feature ideas.