package repositories

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func setupTestDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	schema := `
    CREATE TABLE words (
        id INTEGER PRIMARY KEY,
        english TEXT,
        bahasa_indonesia TEXT
    );
    CREATE TABLE word_review_items (
        id INTEGER PRIMARY KEY,
        word_id INTEGER,
        study_session_id INTEGER,
        correct BOOLEAN,
        created_at DATETIME
    );
    CREATE TABLE groups (
        id INTEGER PRIMARY KEY,
        name TEXT
    );
    CREATE TABLE words_groups (
        id INTEGER PRIMARY KEY,
        word_id INTEGER,
        group_id INTEGER
    );
    CREATE TABLE study_sessions (
        id INTEGER PRIMARY KEY,
        group_id INTEGER,
        created_at DATETIME
    );
    CREATE TABLE study_activities (
        id INTEGER PRIMARY KEY,
        created_at DATETIME
    );
    `
	_, err = db.Exec(schema)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestGetWords(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up test database: %v", err)
	}
	DB = db

	_, err = db.Exec("INSERT INTO words (id, english, bahasa_indonesia) VALUES (1, 'apple', 'apel')")
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	words, totalItems, err := GetWords(1, 10)
	if err != nil {
		t.Fatalf("Failed to get words: %v", err)
	}

	if len(words) != 1 {
		t.Errorf("Expected 1 word, got %d", len(words))
	}

	if totalItems != 1 {
		t.Errorf("Expected 1 total item, got %d", totalItems)
	}
}

func TestGetWord(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up test database: %v", err)
	}
	DB = db

	_, err = db.Exec("INSERT INTO words (id, english, bahasa_indonesia) VALUES (1, 'apple', 'apel')")
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	word, err := GetWord(1)
	if err != nil {
		t.Fatalf("Failed to get word: %v", err)
	}

	if word["english"] != "apple" {
		t.Errorf("Expected 'apple', got '%s'", word["english"])
	}
}
