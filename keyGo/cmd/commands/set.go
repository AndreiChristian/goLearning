package commands

import (
	"fmt"
	"keygo/pkg/store"

	"github.com/spf13/cobra"
)

var setCommand = &cobra.Command{

	Use:   "set [key] [value]",
	Short: "Sets a key-value pair in the store",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		key := args[0]
		value := args[1]

		store := store.GetStore()
		store.Set(key, value)

		fmt.Printf("Setting key %s to value %s \n", key, value)

	},
}

func init() {

	rootCmd.AddCommand(setCommand)

}
