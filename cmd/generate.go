package cmd

import (
	"github.com/spf13/cobra"
	"github.com/TejasGhatte/go-sail/internal/scripts"
)

var CreateProjectCommand *cobra.Command
func init() {
	CreateProjectCommand = &cobra.Command{
		Use: "create [project-name]",
		Short: "Creates a new go project",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			project_name := args[0]
			scripts.CreateProject(project_name)
		},
	}
}