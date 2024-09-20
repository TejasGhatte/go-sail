package scripts

import (
	"fmt"

	"github.com/TejasGhatte/go-sail/internal/models"
	"github.com/TejasGhatte/go-sail/internal/prompts"
	"github.com/TejasGhatte/go-sail/internal/configs"
)

func CreateProject(name string) {
	framework := prompts.SelectFramework()
	database := prompts.SelectDatabase()

	var orm string
	if database != "" {
		orm = prompts.SelectORM()
	}
	caching := prompts.SelectCaching()
	logging := prompts.SelectLogging()

	fmt.Println("Generating project with the following options:")
    fmt.Printf("Framework: %s, Database: %s, ORM: %s, Logging: %t, Caching: %t\n", framework, database, orm, logging, caching)
	
	ctx := models.Options{
		ProjectName: name,
		Framework:   framework,
		Database:    database,
		ORM:         orm,
		Caching:     caching,
		Logging:     logging,
	}

	err := PopulateDirectory(ctx)
	if err != nil {
		fmt.Println("Error creating project:", err)
	}
}

func PopulateDirectory(ctx models.Options) error {

	if err := GitClone(ctx.ProjectName, ctx.Framework, configs.FIBER_URL); err != nil {
		return fmt.Errorf("error cloning repository: %v", err)
	}
	return nil
}