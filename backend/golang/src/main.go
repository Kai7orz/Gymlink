package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func GetProjectRoot() (string, error) {
	_, exPath, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(exPath)

	for {
		parent := filepath.Dir(currentDir)
		if currentDir == parent {
			return "", errors.New("cannot get project root path: not found go.mod")
		}

		if _, err := os.Stat(filepath.Join(currentDir, "go.mod")); err == nil {
			return currentDir, nil
		}

		currentDir = parent
	}
}

func open(path string, count uint) *sql.DB {
	db, err := sql.Open("mysql", path)
	if err != nil {
		log.Fatal("open error: ", err)
	}

	if err = db.Ping(); err != nil {
		time.Sleep(time.Second * 2)
		count--
		fmt.Printf("retry... count:%v\n", count)
		return open(path, count)
	}
	fmt.Println("db connected!!")
	return db
}

func connectDB() *sql.DB {
	var path string = fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8&parseTime=true",
		os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DATABASE"))

	return open(path, 100)
}

func main() {
	db := connectDB()
	defer db.Close()

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		fmt.Println("error mysql.WithInstance()")
		os.Exit(1)
	}

	rootPath, err := GetProjectRoot()
	if err != nil {
		fmt.Println(err)
	}
	rawMigrationFilesPath := filepath.Join(rootPath, "internal", "db", "migration")
	migrationFilesPath := filepath.Join("file://", rawMigrationFilesPath)

	m, err := migrate.NewWithDatabaseInstance(migrationFilesPath, "mysql", driver)
	if err != nil {
		fmt.Println("migration file path:", migrationFilesPath)
		fmt.Println(err)
		os.Exit(1)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("migrate.Up failed ", err)
	}
	log.Println("Migration successful ✅")
}
