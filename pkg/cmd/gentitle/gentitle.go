package gentitle

import (
	"fmt"

	"github.com/spf13/cobra"
)

type GentitleOptions struct {
	Number   int
	Keywords []string
}

func NewCmdGentitle() *cobra.Command {
	o := &GentitleOptions{}

	cmd := &cobra.Command{
		Use:   "gentitle",
		Short: "Generate titles for tech blog",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(o.Number)
			fmt.Println(o.Keywords)
		},
	}

	cmd.Flags().IntVarP(&o.Number, "number", "n", 3, "Number of titles to generate")
	cmd.Flags().StringSliceVarP(
		&o.Keywords,
		"keywords",
		"k",
		[]string{},
		"The keywords of the tech blog content. This flag takes comma-separated value as arguments and split them accordingly.",
	)

	return cmd
}
