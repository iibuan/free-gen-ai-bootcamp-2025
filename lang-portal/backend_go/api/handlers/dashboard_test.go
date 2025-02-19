package handlers

import (
	"backend_go/api/repositories"
	"backend_go/api/services"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func setupTestDB_Dashboard() (*sql.DB, error) {
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

func TestDashboardHandler_GetLastStudySession(t *testing.T) {
	db, err := setupTestDB_Dashboard()
	if err != nil {
		t.Fatalf("Failed to set up test database: %v", err)
	}
	repositories.DB = db

	_, err = db.Exec("INSERT INTO groups (id, name) VALUES (1, 'Basic Vocabulary')")
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	_, err = db.Exec("INSERT INTO study_sessions (id, group_id, created_at) VALUES (1, 1, ?)", time.Now())
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	service := &services.DashboardService{}
	quickStatsService := &services.QuickStatsService{}
	handler := NewDashboardHandler(service, quickStatsService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/api/dashboard/last_study_session", handler.GetLastStudySession)

	req, _ := http.NewRequest("GET", "/api/dashboard/last_study_session", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", w.Code)
	}
}

func TestDashboardHandler_GetStudyProgress(t *testing.T) {
	db, err := setupTestDB_Dashboard()
	if err != nil {
		t.Fatalf("Failed to set up test database: %v", err)
	}
	repositories.DB = db

	_, err = db.Exec("INSERT INTO word_review_items (id, word_id, study_session_id, correct, created_at) VALUES (1, 1, 1, 1, ?)", time.Now())
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	service := &services.DashboardService{}
	quickStatsService := &services.QuickStatsService{}
	handler := NewDashboardHandler(service, quickStatsService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/api/dashboard/study_progress", handler.GetStudyProgress)

	req, _ := http.NewRequest("GET", "/api/dashboard/study_progress", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", w.Code)
	}
}

func TestDashboardHandler_GetQuickStats(t *testing.T) {
	db, err := setupTestDB_Dashboard()
	if err != nil {
		t.Fatalf("Failed to set up test database: %v", err)
	}
	repositories.DB = db

	_, err = db.Exec("INSERT INTO word_review_items (id, word_id, study_session_id, correct, created_at) VALUES (1, 1, 1, 1, ?)", time.Now())
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	service := &services.DashboardService{}
	quickStatsService := &services.QuickStatsService{}
	handler := NewDashboardHandler(service, quickStatsService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/api/dashboard/quick-stats", handler.GetQuickStats)

	req, _ := http.NewRequest("GET", "/api/dashboard/quick-stats", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", w.Code)
	}
}
