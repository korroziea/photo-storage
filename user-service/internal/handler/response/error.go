package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	internalCode = 10
)

type errorResp struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

func Error(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, errorResp{
		Code: internalCode,
		Message: "something went wrong",
	})
}
