package scripts

import (
	"context"
	"fmt"

	"github.com/TejasGhatte/go-sail/internal/helpers"
)

func AnalyseFile(ctx context.Context, path string) error {
	extraData := map[string]interface{}{
		"path": path,
	}
	response, err := helpers.MakeRequest("analyse-file", extraData)
	if err != nil {
		return err
	}

	fmt.Printf("File Analysis Result: %v\n", response.Data)
	return nil
}

func AnalyseFolder(ctx context.Context, path string) error {
	extraData := map[string]interface{}{
		"path": path,
	}
	response, err := helpers.MakeRequest("analyse-folder", extraData)
	if err != nil {
		return err
	}

	fmt.Printf("Folder Analysis Result: %v\n", response.Data)
	return nil
}

func AnalyseRepository(ctx context.Context) error {
	extraData := map[string]interface{}{
		"githubUrl": "repo-url",
	}
	response, err := helpers.MakeRequest("analyse-repository", extraData)
	if err != nil {
		return err
	}

	fmt.Printf("Repository Analysis Result: %v\n", response.Data)
	return nil
}