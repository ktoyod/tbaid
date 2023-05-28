package cmd

import (
	"github.com/ktoyod/tbaid/pkg/cmd/gentitle"
	"github.com/spf13/cobra"
)

func NewTbaidCommand() *cobra.Command {
	cmds := &cobra.Command{
		Use:   "tbaid",
		Short: "tbaid supports to write tech blog",
	}

	gentitleCmd := gentitle.NewCmdGentitle()

	cmds.AddCommand(gentitleCmd)

	return cmds
}
