package prompts

import (
    "github.com/AlecAivazis/survey/v2"
    "fmt"
)

var frameworks = []string{"Go Fiber", "Gin", "Echo", "Chi"}

var databases = []string{"PostgreSQL", "MySQL", "MongoDB", "None"}

var orms = []string{"GORM", "SQLX", "None"}

func SelectFramework() string {
    var framework string
    prompt := &survey.Select{
        Message: "Choose a Go framework:",
        Options: frameworks,
    }
    err := survey.AskOne(prompt, &framework)
    if err != nil {
        fmt.Println("Error selecting framework:", err)
    }
    return framework
}

func SelectDatabase() string {
    var database string
    prompt := &survey.Select{
        Message: "Choose a database (or None):",
        Options: databases,
		Default: "None",
    }
    err := survey.AskOne(prompt, &database)
    if err != nil {
        fmt.Println("Error selecting database:", err)
    }
    if database == "None" {
        return ""
    }
    return database
}

func SelectORM() string {
    var orm string
    prompt := &survey.Select{
        Message: "Choose an ORM (or None):",
        Options: orms,
		Default: "None",
    }
    err := survey.AskOne(prompt, &orm)
    if err != nil {
        fmt.Println("Error selecting ORM:", err)
    }
    if orm == "None" {
        return ""
    }
    return orm
}

func SelectLogging() bool {
    var logging bool
    prompt := &survey.Confirm{
        Message: "Do you want to include logging?",
        Default: true,
    }
    err := survey.AskOne(prompt, &logging)
    if err != nil {
        fmt.Println("Error selecting logging option:", err)
    }
    return logging
}

func SelectCaching() bool {
    var caching bool
    prompt := &survey.Confirm{
        Message: "Do you want to include caching?",
        Default: false,
    }
    err := survey.AskOne(prompt, &caching)
    if err != nil {
        fmt.Println("Error selecting caching option:", err)
    }
    return caching
}