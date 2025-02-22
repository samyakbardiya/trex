package util

import (
	"fmt"
	"io"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// InitLogging initializes logging based on the DEBUG environment variable.
// If the DEBUG variable is unset, it disables logging by directing output to io.Discard and returns nil.
// Otherwise, it attempts to create a "debug.log" file using tea.LogToFile.
// On failure to create the log file, it prints a fatal error and terminates the program.
// If successful, it logs a startup message and returns the file handle.
func InitLogging() *os.File {
	if len(os.Getenv("DEBUG")) == 0 {
		log.SetOutput(io.Discard)
		return nil
	}

	f, err := tea.LogToFile("debug.log", "")
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	log.Println("logging...")
	return f
}
