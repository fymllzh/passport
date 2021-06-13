package index

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/index/index", gin.H{})
}

func Test(c *gin.Context) {
	fmt.Fprint(c.Writer, "this is test page")
}
