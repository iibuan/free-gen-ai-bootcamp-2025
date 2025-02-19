package routes

import (
	"backend_go/api/handlers"
	"backend_go/api/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	dashboardService := &services.DashboardService{}
	quickStatsService := &services.QuickStatsService{}
	dashboardHandler := handlers.NewDashboardHandler(dashboardService, quickStatsService)

	studyActivityService := &services.StudyActivityService{}
	studyActivityHandler := handlers.NewStudyActivityHandler(studyActivityService)

	studySessionsService := &services.StudySessionService{}
	studySessionsHandler := handlers.NewStudySessionsHandler(studySessionsService)

	wordService := &services.WordService{}
	wordsHandler := handlers.NewWordsHandler(wordService)

	groupService := &services.GroupService{}
	groupsHandler := handlers.NewGroupsHandler(groupService)

	resetService := &services.ResetService{}
	resetHandler := handlers.NewResetHandler(resetService)

	r.GET("/api/dashboard/last_study_session", dashboardHandler.GetLastStudySession)
	r.GET("/api/dashboard/study_progress", dashboardHandler.GetStudyProgress)
	r.GET("/api/dashboard/quick-stats", dashboardHandler.GetQuickStats)
	r.GET("/api/study_activities", studyActivityHandler.GetStudyActivities)
	r.GET("/api/study_activities/:id", studyActivityHandler.GetStudyActivity)
	r.POST("/api/study_activities", studyActivityHandler.CreateStudyActivity)
	r.GET("/api/study_sessions", studySessionsHandler.GetStudySessions)
	r.GET("/api/study_sessions/:id", studySessionsHandler.GetStudySession)
	r.GET("/api/study_sessions/:id/words", studySessionsHandler.GetStudySessionWords)
	r.POST("/api/study_sessions/:id/words/:word_id/review", studySessionsHandler.CreateWordReviewItem)
	r.GET("/api/words", wordsHandler.GetWords)
	r.GET("/api/words/:id", wordsHandler.GetWord)
	r.GET("/api/groups", groupsHandler.GetGroups)
	r.GET("/api/groups/:id", groupsHandler.GetGroup)
	r.GET("/api/groups/:id/words", groupsHandler.GetGroupWords)
	r.GET("/api/groups/:id/study_sessions", groupsHandler.GetGroupStudySessions)
	r.POST("/api/reset_history", resetHandler.ResetHistory)
	r.POST("/api/full_reset", resetHandler.FullReset)
}
