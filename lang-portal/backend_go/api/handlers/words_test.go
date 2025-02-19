package handlers

import (
	"backend_go/api/repositories"
	"backend_go/api/services"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
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

func TestWordsHandler_GetWords(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up test database: %v", err)
	}
	repositories.DB = db

	_, err = db.Exec("INSERT INTO words (id, english, bahasa_indonesia) VALUES (1, 'apple', 'apel')")
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	service := &services.WordService{}
	handler := NewWordsHandler(service)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/api/words", handler.GetWords)

	req, _ := http.NewRequest("GET", "/api/words", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", w.Code)
	}
}

func TestWordsHandler_GetWord(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up test database: %v", err)
	}
	repositories.DB = db

	_, err = db.Exec("INSERT INTO words (id, english, bahasa_indonesia) VALUES (1, 'apple', 'apel')")
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	service := &services.WordService{}
	handler := NewWordsHandler(service)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/api/words/:id", handler.GetWord)

	req, _ := http.NewRequest("GET", "/api/words/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", w.Code)
	}
}
