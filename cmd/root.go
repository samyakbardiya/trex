package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

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
	if len(args) == 0 {
		return nil
	}

	filePath, err := util.ValidateFilepath(args[0])
	if err != nil {
		return err
	}
	ctx := context.WithValue(cmd.Context(), util.KeyFilePath, filePath)
	cmd.SetContext(ctx)

	return nil
}

func run(cmd *cobra.Command, args []string) error {
	var data []byte
	var err error

	if filePath, ok := cmd.Context().Value(util.KeyFilePath).(string); ok {
		data, err = os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("failed to read file %q: %w", filePath, err)
		}
	} else {
		data = []byte(util.DefaultText)
	}
	log.Printf("Processing %d bytes of data\n", len(data))
	log.Printf("Proccessed file:\n%s\n", string(data))

	return nil
}
