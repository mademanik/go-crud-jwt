package books

import (
	"github.com/gin-gonic/gin"
	"github.com/mademanik/go-crud-jwt/pkg/common/models"
	"net/http"
)

func (h handler) GetBooks(ctx *gin.Context) {
	var books []models.Book

	if result := h.DB.Find(&books); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &books)
}
