package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/samyakbardiya/trex/internal/ui"
	"github.com/samyakbardiya/trex/internal/util"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "trex [file]",
	Short:        "A TUI tool to work with RegEx",
	Long:         "A TUI tool to work with RegEx",
	Example:      util.CliExample,
	Args:         cobra.MaximumNArgs(1),
	Version:      "0.0.0",
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
	var (
		data     []byte
		filePath string
		err      error
	)
	if len(args) == 0 {
		log.Println("Using default text")
		data = []byte(util.DefaultText)
	} else {
		log.Println("Reading from file")
		filePath, err = util.GetFilePath(args[0])
		if err != nil {
			return err
		}
		data, err = os.ReadFile(filePath)
		if err != nil {
			return err
		}
	}
	cmd.SetContext(context.WithValue(cmd.Context(), util.KeyFileData, data))
	return nil
}

func run(cmd *cobra.Command, args []string) error {
	file, ok := cmd.Context().Value(util.KeyFileData).([]byte)
	if !ok {
		panic("Unable to read the content")
	}

	// NOTE: usage example
	//
	// expr := "Lorem"
	// re, err := util.GetAllMatchingIndex(expr, file)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// log.Println("re", re)

	p := tea.NewProgram(ui.InitialModel(ui.DefaultTime, string(file)), tea.WithAltScreen(), tea.WithMouseCellMotion())
	if _, err := p.Run(); err != nil {
		return fmt.Errorf("error while running program: %w", err)
	}

	return nil
}
