package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	var err error

	cfg := config.NewConfig()

	if len(os.Args) <= 2 {
		logrus.Error("Usage:", os.Args[1], "command", "argument")
		return errors.New("invalid command")
	}

	switch os.Args[1] {
	case "migrate":
		dsn := cfg.DB.DSN + "&multiStatements=true"
		err = Migrate(dsn, os.Args[2])
	case "seed":
		err = errors.New("to be develop")
	default:
		err = errors.New("must specify a command")
	}

	if err != nil {
		return err
	}

	return nil
}

// Migrate to run database migration up or down
func Migrate(dsn string, command string) error {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logrus.Error(err)
		return err
	}
	path, err := os.Getwd()
	if err != nil {
		logrus.Error(err)
		return err
	}

	migrationPath := fmt.Sprintf("file://%s/migrations", path)
	logrus.Infof("migrationPath : %s", migrationPath)

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		logrus.Error(err)
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		migrationPath,
		"mysql",
		driver,
	)
	if err != nil {
		logrus.Error(err)
		return err
	}
	if command == "up" {
		logrus.Info("Migrate up")
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			logrus.Errorf("An error occurred while syncing the database.. %v", err)
			return err
		}
	}

	if command == "down" {
		logrus.Info("Migrate down")
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			logrus.Errorf("An error occurred while syncing the database.. %v", err)
			return err
		}
	}

	if err != nil {
		logrus.Error(err)
		return err
	}

	logrus.Info("Migrate complete")
	return nil
}
