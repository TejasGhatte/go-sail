package main

import (
	"github.com/TejasGhatte/go-sail/cmd"
	"github.com/TejasGhatte/go-sail/internal/initializers"
	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
	Use:   "go-sail",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains`,
}

func main() {
	initializers.LoadConfig("config.yml")
	rootCmd.AddCommand(cmd.CreateProjectCommand)
	cobra.CheckErr(rootCmd.Execute())
}
