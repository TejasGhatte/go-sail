package scripts

import (
	"context"
	"fmt"

	"github.com/TejasGhatte/go-sail/internal/prompts"
	"github.com/TejasGhatte/go-sail/internal/signals"
)

func Signup(ctx context.Context) error {
	ctx = signals.HandleCancellation(ctx)

	username, email, password, err := prompts.PromptUserSignupDetails(ctx)
	if err != nil {
		return err
	}

	// err := sendSignupRequest(ctx, username, email, password)
	// if err != nil {
	// 	return err
	// }
	fmt.Println(username)
	fmt.Println(email)
	fmt.Println(password)
	fmt.Println("Signup successful!")
	return nil
}