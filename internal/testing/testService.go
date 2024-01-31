package testing

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleTest(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Test{Message: "hello!"})
}
