package scripts

import (
	"fmt"
	"os"
	"path/filepath"
	"context"
	"time"

	"github.com/briandowns/spinner"
	"github.com/TejasGhatte/go-sail/internal/helpers"
	"github.com/TejasGhatte/go-sail/internal/initializers"
	"github.com/TejasGhatte/go-sail/internal/models"
	"github.com/TejasGhatte/go-sail/internal/prompts"
)

func CreateProject(ctx context.Context, name string) error {
	framework := prompts.SelectFramework()
	database := prompts.SelectDatabase()

	var orm string
	if database != "" {
		orm = prompts.SelectORM()
	}

	fmt.Println("Generating project with the following options:")
    fmt.Printf("Framework: %s, Database: %s, ORM: %s\n", framework, database, orm)
	
	options := &models.Options{
		ProjectName: name,
		Framework:   framework,
		Database:    database,
		ORM:         orm,
	}

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Start()
	defer s.Stop()

	for i := 0; i < 2; i++ {
		select {
		case <-ctx.Done():
			return fmt.Errorf("project creation interrupted")
		default:
			s.Suffix = fmt.Sprintf(" Creating project: %s", name)
			time.Sleep(1 * time.Second)
		}
	}

	err := PopulateDirectory(options)
	if err != nil {
		return err
	}

	return nil
}

func PopulateDirectory(ctx *models.Options) error {
	if err := GitClone(ctx.ProjectName, ctx.Framework, initializers.Config.Repositories[ctx.Framework]); err != nil {
		return fmt.Errorf("error cloning repository: %v", err)
	}

	currentDir, _ := os.Getwd()
	folder := filepath.Join(currentDir, ctx.ProjectName, "initializers")

	if ctx.Database != "" && ctx.ORM != "" {
		provider, err := helpers.ProviderFactory(ctx.Database, ctx.ORM)
		if err != nil {
			return fmt.Errorf("error creating database provider: %v", err)
		}

		err = helpers.GenerateDatabaseFile(folder, provider)
		if err != nil {
			return fmt.Errorf("error generating database file: %v", err)
		}

		err = helpers.GenerateMigrationFile(folder, provider)
		if err != nil {
			return fmt.Errorf("error generating migration file: %v", err)
		}

		err = helpers.ResolveImportErr(ctx.ProjectName)
		if err != nil {
			return fmt.Errorf("error resolving imports: %v", err)
		}

	}
	return nil
}