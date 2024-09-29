package helpers

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)
func RemoveFolders(rootFolder string, foldersToRemove []string) {
	for _, folder := range foldersToRemove {
		_ = os.RemoveAll(filepath.Join(rootFolder, folder))
	}
}

func ResolveImportErr(dir string) error {
	currentDir, _ := os.Getwd()
	folder := filepath.Join(currentDir, dir)
	if err := os.Chdir(folder); err != nil {
		return fmt.Errorf("failed to change directory: %w", err)
	}

	commands := []struct {
		name string
		args []string
	}{
		{"go", []string{"mod", "tidy"}},
		{"go", []string{"install", "golang.org/x/tools/cmd/goimports@latest"}},
		{"goimports", []string{"-w", "."}},
	}

	for _, cmd := range commands {
		command := exec.Command(cmd.name, cmd.args...)
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr

		if err := command.Run(); err != nil {
			return fmt.Errorf("failed to execute %s: %w", cmd.name, err)
		}
	}

	return nil
}