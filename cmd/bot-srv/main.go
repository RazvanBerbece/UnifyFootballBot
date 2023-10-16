package main

import (
	"fmt"

	"github.com/RazvanBerbece/UnifyFootballBot/internal/setup"
)

func main() {

	// Retrieve the configured runtime environment
	env := setup.Setup(".env")

	// Configure DiscordGo base
	fmt.Println(env.AppId)

}
