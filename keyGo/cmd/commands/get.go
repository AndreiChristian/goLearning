package commands

import (
	"fmt"
	"keygo/pkg/store"

	"github.com/spf13/cobra"
)

var getCommand = &cobra.Command{

	Use:   "GET [key]",
	Short: "Gets a value from the store by key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		key := args[0]

		store := store.GetStore()

		value, ok := store.Get(key)

		if ok {
			fmt.Printf("The value of key '%s' is '%s'\n", key, value)
		} else {
			fmt.Printf("No value found for key '%s'\n", key)
		}

	},
}

func init() {
	rootCmd.AddCommand(getCommand)
}
