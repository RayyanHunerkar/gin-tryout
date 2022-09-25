package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rayyanhunerkar/gin-tryout/pkg/common/models"
)

// DeleteBook godoc
// @Summary Delete a single Book
// @Description Takes an ID and deletes the book by the ID
// @Tags books
// @Produce json
// @Param id path int true "search by ID"
// @Success 200 {object} models.Book
// @Router /books/{id} [delete]
func (h handler) DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&book)

	ctx.Status(http.StatusNoContent)
}
