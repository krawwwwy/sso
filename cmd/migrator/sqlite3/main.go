package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/golang-migrate/migrate/v4"                    // либа для миграций
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3" // драйвер для миграций sqlite3
	_ "github.com/golang-migrate/migrate/v4/source/file"      // драйвер для получения миграций из файлов
)

func main() {
	var storagePath, migrationsPath, migrationsTable string
	var down bool

	flag.StringVar(&storagePath, "storage-path", "", "path to storage")
	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations")
	flag.StringVar(&migrationsTable, "migrations-table", "migrations", "name of migrations")
	flag.BoolVar(&down, "down", false, "run down migrations")
	flag.Parse()

	for _, value := range []string{storagePath, migrationsPath, migrationsTable} {
		if value == "" {
			panic(fmt.Sprintf("%s is required", value))
		}
	}

	m, err := migrate.New(
		"file://"+migrationsPath,
		fmt.Sprintf("sqlite3://%s?x-migrations-table=%s", storagePath, migrationsTable),
	)
	if err != nil {
		panic(err)
	}

	var migrationErr error
	if down {
		migrationErr = m.Down()
	} else {
		migrationErr = m.Up()
	}

	if migrationErr != nil {
		if errors.Is(migrationErr, migrate.ErrNoChange) {
			if down {
				fmt.Println("no migrations to rollback")
			} else {
				fmt.Println("no migrations to apply")
			}
			return
		}
		panic(migrationErr)
	}

	if down {
		fmt.Println("migrations rolled back successfully")
	} else {
		fmt.Println("migrations applied successfully")
	}
}
