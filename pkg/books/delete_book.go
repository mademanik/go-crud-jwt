package books

import (
	"github.com/gin-gonic/gin"
	"github.com/mademanik/go-crud-jwt/pkg/common/models"
	"net/http"
)

func (h handler) DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")
	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&book)
	ctx.Status(http.StatusOK)
}
