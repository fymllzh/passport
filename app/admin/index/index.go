package index

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wuzehv/passport/model/client"
	"github.com/wuzehv/passport/model/user"
	"github.com/wuzehv/passport/service/db"
	"net/http"
)

func Index(c *gin.Context) {
	var u []user.User
	db.Db.Find(&u)

	var cl []client.Client
	db.Db.Find(&cl)
	c.HTML(http.StatusOK, "admin/index/index", gin.H{
		"users": u,
		"clients": cl,
	})
}

func Test(c *gin.Context) {
	fmt.Fprint(c.Writer, "this is test page")
}
