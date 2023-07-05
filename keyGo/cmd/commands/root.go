package commands

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "KeyGo",
	Short: "KeyGo is an in-memory database written in go",
	Long:  `A cool project written in Go`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
