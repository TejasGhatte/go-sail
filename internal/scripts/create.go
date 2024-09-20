package scripts

import (
	"fmt"
	"github.com/TejasGhatte/go-sail/internal/prompts"
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
}