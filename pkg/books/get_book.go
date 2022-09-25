package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rayyanhunerkar/gin-tryout/pkg/common/models"
)

// GetBook godoc
// @Summary Retrieves a single Book
// @Description Takes an ID and retireves the book by the ID
// @Tags books
// @Produce json
// @Param id path int true "search by ID"
// @Success 200 {object} models.Book
// @Router /books/{id} [get]
func (h handler) GetBook(ctx *gin.Context) {
	id := ctx.Param("id")

	var book models.Book
	if result := h.DB.First(&book, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusFound, &book)
}
