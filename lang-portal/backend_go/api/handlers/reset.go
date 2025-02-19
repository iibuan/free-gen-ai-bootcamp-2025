package handlers

import (
	"backend_go/api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResetHandler struct {
	Service *services.ResetService
}

func NewResetHandler(service *services.ResetService) *ResetHandler {
	return &ResetHandler{Service: service}
}

func (h *ResetHandler) ResetHistory(c *gin.Context) {
	err := h.Service.ResetHistory()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "History reset successfully"})
}

func (h *ResetHandler) FullReset(c *gin.Context) {
	err := h.Service.FullReset()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Full reset successfully"})
}
