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

	flag.StringVar(&storagePath, "storage-path", "", "path to storage")
	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations")
	flag.StringVar(&migrationsTable, "migrations-table", "migrations", "name of migrations")
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

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrates to apply")

			return
		}

		panic(err)
	}

	fmt.Println("migrations applied successfully")
}
