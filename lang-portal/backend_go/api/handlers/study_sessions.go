package handlers

import (
	"backend_go/api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StudySessionsHandler struct {
	Service *services.StudySessionService
}

func NewStudySessionsHandler(service *services.StudySessionService) *StudySessionsHandler {
	return &StudySessionsHandler{Service: service}
}

func (h *StudySessionsHandler) GetStudySessions(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "100"))

	sessions, totalItems, err := h.Service.GetStudySessions(page, perPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items": sessions,
		"pagination": gin.H{
			"page":        page,
			"per_page":    perPage,
			"total_pages": (totalItems + perPage - 1) / perPage,
			"total_items": totalItems,
		},
	})
}

func (h *StudySessionsHandler) GetStudySession(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	session, err := h.Service.GetStudySession(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, session)
}

func (h *StudySessionsHandler) GetStudySessionWords(c *gin.Context) {
	studySessionID, _ := strconv.Atoi(c.Param("id"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "100"))

	words, totalItems, err := h.Service.GetStudySessionWords(studySessionID, page, perPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items": words,
		"pagination": gin.H{
			"page":        page,
			"per_page":    perPage,
			"total_pages": (totalItems + perPage - 1) / perPage,
			"total_items": totalItems,
		},
	})
}

func (h *StudySessionsHandler) CreateWordReviewItem(c *gin.Context) {
	studySessionID, _ := strconv.Atoi(c.Param("id"))
	wordID, _ := strconv.Atoi(c.Param("word_id"))

	var params struct {
		Correct bool `json:"correct"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reviewItem, err := h.Service.CreateWordReviewItem(studySessionID, wordID, params.Correct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reviewItem)
}
