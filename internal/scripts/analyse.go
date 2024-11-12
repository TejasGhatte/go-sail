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
	response, err := helpers.MakeAnalysisReq("analyse-file", extraData)
	if err != nil {
		return err
	}

	fmt.Printf("File Analysis Result: %v\n", response.Analysis)
	return nil
}

func AnalyseFolder(ctx context.Context, path string) error {
	extraData := map[string]interface{}{
		"path": path,
	}
	response, err := helpers.MakeAnalysisReq("analyse-folder", extraData)
	if err != nil {
		return err
	}

	fmt.Printf("Folder Analysis Result: %v\n", response.Analysis)
	return nil
}

func AnalyseRepository(ctx context.Context) error {
	extraData := map[string]interface{}{
		"githubUrl": "repo-url",
	}
	response, err := helpers.MakeAnalysisReq("analyse-repository", extraData)
	if err != nil {
		return err
	}

	fmt.Printf("Repository Analysis Result: %v\n", response.Analysis)
	return nil
}