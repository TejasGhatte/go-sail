package scripts

import (
	"context"
	"fmt"

	"github.com/TejasGhatte/go-sail/internal/helpers"
)

func SecAnalyseFile(ctx context.Context, path string) error {
	extraData := map[string]interface{}{
		"path": path,
	}
	response, err := helpers.MakeRequest("security-analyse-file", extraData)
	if err != nil {
		return err
	}

	fmt.Printf("File Analysis Result: %v\n", response.Data)
	return nil
}

func SecAnalyseFolder(ctx context.Context, path string) error {
	extraData := map[string]interface{}{
		"path": path,
	}
	response, err := helpers.MakeRequest("security-analyse-file", extraData)
	if err != nil {
		return err
	}

	fmt.Printf("File Analysis Result: %v\n", response.Data)
	return nil
}

func SecAnalyseRepository(ctx context.Context) error {
	response, err := helpers.MakeRequest("security-analyse-file", nil)
	if err != nil {
		return err
	}

	fmt.Printf("File Analysis Result: %v\n", response.Data)
	return nil
}