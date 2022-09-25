package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rayyanhunerkar/gin-tryout/pkg/common/models"
)

type AddBooksRequestBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

// AddBook godoc
// @Summary Creates a new book
// @Description Takes a book JSON and store in DB. Return saved JSON.
// @Tags books
// @Accept json
// @Produce json
// @Param book body AddBooksRequestBody true "AddBooksRequestBody JSON"
// @Success 201 {object} models.Book
// @Router /books [post]
func (h handler) AddBook(ctx *gin.Context) {
	body := AddBooksRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	var book models.Book

	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description

	if result := h.DB.Create(&book); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	ctx.JSON(http.StatusCreated, &book)
}
