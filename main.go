package main

import (
	"github.com/samyakbardiya/trex/cmd"
	"github.com/samyakbardiya/trex/internal/util"
)

// main is the entry point for the application. It initializes logging, executes the primary command logic,
// and defers cleanup of the logging resource.
func main() {
	f := util.InitLogging()
	cmd.Execute()
	defer f.Close()
}
