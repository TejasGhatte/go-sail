package scripts

import (
	"context"
	"fmt"

	"github.com/TejasGhatte/go-sail/internal/helpers"
)

func DescribeFile(ctx context.Context, path string) error {
	extraData := map[string]interface{}{
		"path": path,
	}
	response, err := helpers.MakeRequest("describe-file", extraData)
	if err != nil {
		return err
	}

	fmt.Printf("File Analysis Result: %v\n", response.Data)
	return nil
}

func DescribeFolder(ctx context.Context, path string) error {
	extraData := map[string]interface{}{
		"path": path,
	}
	response, err := helpers.MakeRequest("describe-folder", extraData)
	if err != nil {
		return err
	}

	fmt.Printf("File Analysis Result: %v\n", response.Data)
	return nil
}

func DescribeRepository(ctx context.Context) error {
	response, err := helpers.MakeRequest("describe-repository", nil)
	if err != nil {
		return err
	}

	fmt.Printf("File Analysis Result: %v\n", response.Data)
	return nil
}