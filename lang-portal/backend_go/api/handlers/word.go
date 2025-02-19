package handlers

import (
	"backend_go/api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WordsHandler struct {
	Service *services.WordService
}

func NewWordsHandler(service *services.WordService) *WordsHandler {
	return &WordsHandler{Service: service}
}

func (h *WordsHandler) GetWords(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "100"))

	words, totalItems, err := h.Service.GetWords(page, perPage)
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

func (h *WordsHandler) GetWord(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	word, err := h.Service.GetWord(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, word)
}
