package handlers

import (
	"backend_go/api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GroupsHandler struct {
	Service *services.GroupService
}

func NewGroupsHandler(service *services.GroupService) *GroupsHandler {
	return &GroupsHandler{Service: service}
}

func (h *GroupsHandler) GetGroups(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "100"))

	groups, totalItems, err := h.Service.GetGroups(page, perPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items": groups,
		"pagination": gin.H{
			"page":        page,
			"per_page":    perPage,
			"total_pages": (totalItems + perPage - 1) / perPage,
			"total_items": totalItems,
		},
	})
}

func (h *GroupsHandler) GetGroup(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	group, err := h.Service.GetGroup(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, group)
}

func (h *GroupsHandler) GetGroupWords(c *gin.Context) {
	groupID, _ := strconv.Atoi(c.Param("id"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "100"))

	words, totalItems, err := h.Service.GetGroupWords(groupID, page, perPage)
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

func (h *GroupsHandler) GetGroupStudySessions(c *gin.Context) {
	groupID, _ := strconv.Atoi(c.Param("id"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "100"))

	sessions, totalItems, err := h.Service.GetGroupStudySessions(groupID, page, perPage)
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
