package main

import (
	migrations "github.com/ShkrutDenis/go-migrations"
	run "github.com/ShkrutDenis/go-migrations/store"
	"github.com/jabardigitalservice/portal-jabar-api/database/migrations/list"
)

func main() {
	migrations.Run(getMigrationsList())
}

func getMigrationsList() []run.Migratable {
	return []run.Migratable{
		&list.CreateNewsTable{},
	}
}
