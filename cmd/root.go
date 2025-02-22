package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/samyakbardiya/trex/internal/ui"
	"github.com/samyakbardiya/trex/internal/util"
	"github.com/spf13/cobra"
)

const version = "0.0.0"

var rootCmd = &cobra.Command{
	Use:          "trex [file]",
	Short:        "A TUI tool to work with RegEx",
	Long:         "A TUI tool to work with RegEx",
	Example:      util.CliExample,
	Args:         cobra.MaximumNArgs(1),
	Version:      version,
	PreRunE:      preRun,
	RunE:         run,
	SilenceUsage: true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func preRun(cmd *cobra.Command, args []string) error {
	data, err := loadInputData(args)
	if err != nil {
		return fmt.Errorf("failed to load input data: %w", err)
	}

	log.Printf("content: %q", data)
	ctx := context.WithValue(cmd.Context(), util.KeyFileData, data)
	cmd.SetContext(ctx)
	return nil
}

func run(cmd *cobra.Command, args []string) error {
	data, ok := cmd.Context().Value(util.KeyFileData).([]byte)
	if !ok {
		return fmt.Errorf("unable to read content")
	}

	p := tea.NewProgram(
		ui.New(string(data)),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)
	if _, err := p.Run(); err != nil {
		return fmt.Errorf("error while running program: %w", err)
	}

	return nil
}

func loadInputData(args []string) ([]byte, error) {
	var text string

	if len(args) == 0 {
		log.Println("Using default text")
		text = util.DefaultText
	} else {
		log.Println("Reading from file")
		filePath, err := util.GetFilePath(args[0])
		if err != nil {
			return nil, err
		}

		data, err := os.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to read file: %w", err)
		}

		text = string(data)
	}

	text = strings.TrimSpace(text)
	text = strings.ReplaceAll(text, "\r\n", "\n")

	return []byte(text), nil
}
