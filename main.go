package main

import (
	"github.com/TejasGhatte/go-sail/cmd"
	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
	Use:   "go-sail",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains`,
}

func main() {
	rootCmd.AddCommand(cmd.CreateProjectCommand)
	cobra.CheckErr(rootCmd.Execute())
}
