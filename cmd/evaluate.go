package cmd

// import (
// 	"context"
// 	"fmt"

// 	"github.com/TejasGhatte/go-sail/internal/scripts"
// 	"github.com/TejasGhatte/go-sail/internal/signals"
// 	"github.com/spf13/cobra"
// )

// var EvaluateCommand *cobra.Command

// func init() {
// 	EvaluateCommand = &cobra.Command{
// 		Use: "evaluate",
// 		Short: "Evaluate your oroject",
// 		Run: func(cmd *cobra.Command, args []string) {
// 			ctx := context.Background()
// 			ctx = signals.HandleCancellation(ctx)
// 			err := scripts.Evaluate(ctx)
// 			if err != nil {
// 				fmt.Printf("Error signing up: %v\n", err)
// 			}
// 		},
// 	}
// }