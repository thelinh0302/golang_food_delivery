package ginupload

import (
	"Tranning_food/common"
	"Tranning_food/component"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Upload(appCtx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		c.SaveUploadedFile(fileHeader, fmt.Sprintf("./static/%s", fileHeader.Filename))

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
