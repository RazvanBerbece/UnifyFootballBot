package main

import (
	"fmt"

	"github.com/RazvanBerbece/UnifyFootballBot/internal/setup"
)

func main() {

	// Setup Runtime Environment
	env := setup.Setup(".env")

	// Configure DiscordGo base
	fmt.Println(env.AppId)

}
