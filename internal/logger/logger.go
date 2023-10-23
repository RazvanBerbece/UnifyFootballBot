package logger

import (
	"log"
	"os"
	"path/filepath"

	"github.com/bwmarrin/discordgo"
)

func LogHandlerCall(handlerName string, filename string) {

	// if filename not provided, use default
	if filename == "" {
		filename = "handler_calls.log"
	}

	// If the file and/or path doesn't exist, create, or append to the file
	pathToLogfile := filepath.Join(".", "logs/")
	err := os.MkdirAll(pathToLogfile, os.ModePerm)
	if err != nil {
		log.Fatal("Could not create filepath for Handler calls log file. Err = ", err)
	}
	file, err := os.OpenFile("logs/"+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Could not open log file for Handler calls. Err = ", err)
	}
	defer file.Close()

	log.SetOutput(file)

	log.Printf("%s", handlerName)

	// Set logging output back to stderr
	log.SetOutput(os.Stderr)

}

func LogSentMessage(handlerName string, message string) {

	// If the file and/or path doesn't exist, create, or append to the file
	pathToLogfile := filepath.Join(".", "logs/")
	err := os.MkdirAll(pathToLogfile, os.ModePerm)
	if err != nil {
		log.Fatal("Could not create filepath for sent messages log file. Err = ", err)
	}
	file, err := os.OpenFile("logs/sent_messages.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Could not open log file for sent messages. Err = ", err)
	}
	defer file.Close()

	log.SetOutput(file)

	log.Printf("Handler %s sent message \"%q\"", handlerName, message)

	log.SetOutput(os.Stderr)

}

func LogAddReaction(handlerName string, emojiId string) {

	// If the file and/or path doesn't exist, create, or append to the file
	pathToLogfile := filepath.Join(".", "logs/")
	err := os.MkdirAll(pathToLogfile, os.ModePerm)
	if err != nil {
		log.Fatal("Could not create filepath for added reactions log file. Err = ", err)
	}
	file, err := os.OpenFile("logs/added_reactions.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Could not open log file for added reactions. Err = ", err)
	}
	defer file.Close()

	log.SetOutput(file)

	log.Printf("Handler %s added reaction with ID \"%s\"", handlerName, emojiId)

	log.SetOutput(os.Stderr)

}

func LogCreateReaction(handlerName string, emojiName string) {

	// If the file and/or path doesn't exist, create, or append to the file
	pathToLogfile := filepath.Join(".", "logs/")
	err := os.MkdirAll(pathToLogfile, os.ModePerm)
	if err != nil {
		log.Fatal("Could not create filepath for created reactions log file. Err = ", err)
	}
	file, err := os.OpenFile("logs/created_reactions.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Could not open log file for created reactions. Err = ", err)
	}
	defer file.Close()

	log.SetOutput(file)

	log.Printf("Handler %s created reaction with ID \"%s\"", handlerName, emojiName)

	log.SetOutput(os.Stderr)

}

func LogTeamAssignment(handlerName string, userId string, team string) {

	// If the file and/or path doesn't exist, create, or append to the file
	pathToLogfile := filepath.Join(".", "logs/")
	err := os.MkdirAll(pathToLogfile, os.ModePerm)
	if err != nil {
		log.Fatal("Could not create filepath for team assignment log file. Err = ", err)
	}
	file, err := os.OpenFile("logs/team_assignment.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Could not open log file for team assignment. Err = ", err)
	}
	defer file.Close()

	log.SetOutput(file)

	log.Printf("Handler %s registered user with ID %s favourite team %s", handlerName, userId, team)

	log.SetOutput(os.Stderr)

}

func LogSlashCommand(command string, userId string) {

	// If the file and/or path doesn't exist, create, or append to the file
	pathToLogfile := filepath.Join(".", "logs/")
	err := os.MkdirAll(pathToLogfile, os.ModePerm)
	if err != nil {
		log.Fatal("Could not create filepath for slash commands log file. Err = ", err)
	}
	file, err := os.OpenFile("logs/slash_commands.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Could not open log file for slash commands. Err = ", err)
	}
	defer file.Close()

	log.SetOutput(file)

	log.Printf("Slash command %s triggered by user %s", command, userId)

	log.SetOutput(os.Stderr)

}

func LogCleanup(resource string) {

	// If the file and/or path doesn't exist, create, or append to the file
	pathToLogfile := filepath.Join(".", "logs/")
	err := os.MkdirAll(pathToLogfile, os.ModePerm)
	if err != nil {
		log.Fatal("Could not create filepath for cleaned up resources log file. Err = ", err)
	}
	file, err := os.OpenFile("logs/cleanup.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Could not open log file for cleaned up resources. Err = ", err)
	}
	defer file.Close()

	log.SetOutput(file)

	log.Printf("Resource %s deleted", resource)

	log.SetOutput(os.Stderr)

}

func LogRoleCreation(role *discordgo.Role) {

	// If the file and/or path doesn't exist, create, or append to the file
	pathToLogfile := filepath.Join(".", "logs/")
	err := os.MkdirAll(pathToLogfile, os.ModePerm)
	if err != nil {
		log.Fatal("Could not create filepath for the role creation log file. Err = ", err)
	}
	file, err := os.OpenFile("logs/role_creation.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Could not open log file for role creation. Err = ", err)
	}
	defer file.Close()

	log.SetOutput(file)

	log.Printf("Role with name %s; colour %d; has been created", role.Name, role.Color)

	log.SetOutput(os.Stderr)

}
