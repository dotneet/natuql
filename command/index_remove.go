package command

import (
	"fmt"
	"github.com/dotneet/natuql/index"
	"github.com/spf13/cobra"
	"os"
)

func IndexRemoveCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "index-remove",
		Short:   "remove an index from the database.",
		Example: "natuql index-remove",
		Run: func(cmd *cobra.Command, args []string) {
			err := index.RemoveIndex()
			if err != nil {
				fmt.Fprintln(os.Stderr, "error: %v", err)
			}
			fmt.Println("index removed.")
		},
	}

}
