//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// Build the project
func Build() error {
	fmt.Println("Building the project...")
	cmd := exec.Command("go", "build", "-o", "bin/server", "./cmd/server")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Run the project
func Run() error {
	fmt.Println("Running the project...")
	cmd := exec.Command("go", "run", "./cmd/server/main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Setup the database
func SetupDB() error {
	fmt.Println("Setting up the database...")
	dbPath := "words.db"
	if _, err := os.Stat(dbPath); err == nil {
		fmt.Println("Database already exists. Removing it...")
		os.Remove(dbPath)
	}

	fmt.Println("Creating new database...")
	file, err := os.Create(dbPath)
	if err != nil {
		return err
	}
	file.Close()

	fmt.Println("Running migrations...")
	return Migrate()
}

// Run database migrations
func Migrate() error {
	fmt.Println("Running migrations...")
	migrationsPath := "./db/migrations"
	files, err := filepath.Glob(filepath.Join(migrationsPath, "*.sql"))
	if err != nil {
		return err
	}

	for _, file := range files {
		fmt.Printf("Running migration: %s\n", file)
		cmd := exec.Command("sqlite3", "words.db", ".read "+file)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}

// Seed the database
func Seed() error {
	fmt.Println("Seeding the database...")
	seedFile := "./db/seeds/basic_vocabulary.json"
	cmd := exec.Command("sqlite3", "words.db", ".read "+seedFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
