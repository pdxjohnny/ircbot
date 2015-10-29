package commands

import (
	"github.com/spf13/cobra"

	"github.com/pdxjohnny/ircbot/echo"
)

var Commands = []*cobra.Command{
	&cobra.Command{
		Use:   "echo",
		Short: "Echo text in channel",
		Run: func(cmd *cobra.Command, args []string) {
			ConfigBindFlags(cmd)
			echo.Run()
		},
	},
}

func init() {
	ConfigDefaults(Commands...)
}
