package dbase

import (
	"database/sql"
	"fmt"
	"gymlink/internal/utils"
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(db *sql.DB, isUp bool) error {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		fmt.Println("error mysql.WithInstance()")
		return err
	}
	rootPath, err := utils.GetProjectRoot()
	if err != nil {
		fmt.Println(err)
	}
	rawMigrationFilesPath := filepath.Join(rootPath, "internal", "dbase", "migration")
	migrationFilesPath := filepath.Join("file://", rawMigrationFilesPath)

	m, err := migrate.NewWithDatabaseInstance(migrationFilesPath, "mysql", driver)
	if err != nil {
		fmt.Println("migration file path:", migrationFilesPath)
		fmt.Println(err)
		os.Exit(1)
	}

	if isUp {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal("migrate.Up failed ", err)
		}
	} else {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal("migrate.Down failed ", err)
		}
	}
	log.Println("Migration successful ✅")
	return nil
}

func MigrateUp(db *sql.DB) error {
	return Migrate(db, true)
}

func MigrateDown(db *sql.DB) error {
	return Migrate(db, false)
}
