package main

import (
	"github.com/samyakbardiya/trex/cmd"
	"github.com/samyakbardiya/trex/internal/util"
)

func main() {
	f := util.InitLogging()
	cmd.Execute()
	defer f.Close()
}
