package logger

import (
	"bufio"
	"os"
	"regexp"
	"testing"
)

func TestLogHandlerCall(t *testing.T) {

	// Arrange
	LogHandlerCall("TestHandlerCall", "TEST_handler_calls.log")

	// Assert
	logsPattern := `\d{4}/\d{2}/\d{2} \d{2}:\d{2}:\d{2} (\w+)`
	regex := regexp.MustCompile(logsPattern)
	// Open the .log file
	file, err := os.Open("logs/" + "TEST_handler_calls.log")
	if err != nil {
		t.Error("Error:", err)
	}
	defer file.Close()

	// Read the file line by line and match the pattern
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		match := regex.FindStringSubmatch(line)
		if len(match) > 1 {
			word := match[1] // Extract the word captured by the regex (omit the timestamp)
			if word != "TestHandlerCall" {
				t.Error("A call to TestHandlerCall should be logged after calling LogHandlerCall. match =", word)
			}
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		t.Error("Error during scanning:", err)
	}

	// Cleanup test files generated in the logger package
	e := os.RemoveAll("logs/")
	if e != nil {
		t.Error("Could not cleanup after test.")
	}

}
