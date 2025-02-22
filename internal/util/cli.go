package util

import (
	"fmt"
	"io"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

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
	log.Println("########################################")
	return f
}
