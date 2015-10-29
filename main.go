package main

import (
	"runtime"

	"github.com/spf13/cobra"

	"github.com/pdxjohnny/ircbot/commands"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var rootCmd = &cobra.Command{Use: "ircbot"}
	rootCmd.AddCommand(commands.Commands...)
	rootCmd.Execute()
}
