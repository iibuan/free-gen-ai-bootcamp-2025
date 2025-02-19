package handlers

import (
	"backend_go/api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	Service           *services.DashboardService
	QuickStatsService *services.QuickStatsService
}

func NewDashboardHandler(service *services.DashboardService, quickStatsService *services.QuickStatsService) *DashboardHandler {
	return &DashboardHandler{Service: service, QuickStatsService: quickStatsService}
}

func (h *DashboardHandler) GetLastStudySession(c *gin.Context) {
	result, err := h.Service.GetLastStudySession()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *DashboardHandler) GetStudyProgress(c *gin.Context) {
	result, err := h.Service.GetStudyProgress()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *DashboardHandler) GetQuickStats(c *gin.Context) {
	result, err := h.QuickStatsService.GetQuickStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
