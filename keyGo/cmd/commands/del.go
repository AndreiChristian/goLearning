package commands

import (
	"fmt"

	"keygo/pkg/store"

	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:   "del [key]",
	Short: "Deletes a key-value pair in the store",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]

		store := store.GetStore()
		ok := store.Del(key)

		if ok {
			fmt.Printf("Key '%s' is deleted\n", key)
		} else {
			fmt.Printf("No key '%s' found to delete\n", key)
		}
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}
