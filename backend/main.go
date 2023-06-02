package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func main() {
	app := pocketbase.New()

	migratecmd.MustRegister(app, app.RootCmd, &migratecmd.Options{
		Automigrate: true, // auto creates migration files when making collection changes
	})

	app.OnRecordBeforeCreateRequest().Add(func(e *core.RecordCreateEvent) error {
		// overwrite the newly submitted "posts" record status to pending
		if e.Record.Collection().Name == "incidents" {
			e.Record.Set("status", "new")
		}

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
