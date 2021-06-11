package index

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	fmt.Fprint(c.Writer, "this is index page")
}

func Test(c *gin.Context) {
	fmt.Fprint(c.Writer, "this is test page")
}
