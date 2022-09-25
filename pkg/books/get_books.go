package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rayyanhunerkar/gin-tryout/pkg/common/models"
)

// GetBooks godoc
// @Summary Retrieves an array of Books
// @Description Retrieves an array of Books
// @Tags books
// @Produce json
// @Success 200 {object} []models.Book
// @Router /books [get]
func (h handler) GetBooks(ctx *gin.Context) {
	var books []models.Book

	if result := h.DB.Find(&books); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	ctx.JSON(http.StatusOK, &books)
}
