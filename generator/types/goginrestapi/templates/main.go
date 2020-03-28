package templates

// Constants

const (
	MainGo = `package main

import (
	"log"

	app "{{ .Package.Name }}/internal/app"
)

// @title {{ .FriendlyName }}
// @version 1.0
// @description {{ .Description }}

// @contact.name API Support

// @license.name {{ .LicenseType }}

// @host localhost:8080
// @BasePath /
func main() {
	application, err := app.NewAppFromEnv()

	if err != nil {
		log.Fatal(err)
	}

	err = application.Run()

	if err != nil {
		log.Fatal(err)
	}
}

`
)
