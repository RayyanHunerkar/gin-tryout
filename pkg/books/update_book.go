package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rayyanhunerkar/gin-tryout/pkg/common/models"
)

type UpdateBookRequestBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

// UpdateBook godoc
// @Summary Updates a single Book
// @Description Takes an ID and updates the book by the ID
// @Tags books
// @Produce json
// @Param id path int true "search by ID"
// @Param book body UpdateBookRequestBody true "UpdateBookRequestBody JSON"
// @Success 200 {object} models.Book
// @Router /books/{id} [put]
func (h handler) UpdateBook(ctx *gin.Context) {
	id := ctx.Param("id")
	body := UpdateBookRequestBody{}

	if err := ctx.Bind(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description
	h.DB.Save(&book)
	ctx.JSON(http.StatusOK, &book)
}
