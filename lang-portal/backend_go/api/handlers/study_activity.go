package handlers

import (
	"backend_go/api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StudyActivityHandler struct {
	Service *services.StudyActivityService
}

func NewStudyActivityHandler(service *services.StudyActivityService) *StudyActivityHandler {
	return &StudyActivityHandler{Service: service}
}

func (h *StudyActivityHandler) GetStudyActivities(c *gin.Context) {
	activities, err := h.Service.GetStudyActivities()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, activities)
}

func (h *StudyActivityHandler) GetStudyActivity(c *gin.Context) {
	id := c.Param("id")
	activity, err := h.Service.GetStudyActivity(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, activity)
}

func (h *StudyActivityHandler) CreateStudyActivity(c *gin.Context) {
	var params struct {
		StudySessionID int `json:"study_session_id"`
		GroupID        int `json:"group_id"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	activity, err := h.Service.CreateStudyActivity(params.StudySessionID, params.GroupID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, activity)
}
