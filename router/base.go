package router

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/wuzehv/passport/util"
	"io"
	"os"
	"time"
)

type logFormat struct {
	ClientIp   string
	Timestamp  string
	Method     string
	Path       string
	Code       int
	UserAgent  string
	Message    string
	BodySize   int
	Host       string
	RemoteAddr string
}

func InitRouter() *gin.Engine {
	gin.DisableConsoleColor()

	rootDir := util.ENV("", "root_dir")
	logDir := util.ENV("log", "dir")

	finalDir := rootDir + logDir
	if _, err := os.Stat(finalDir); os.IsNotExist(err) {
		if err = os.MkdirAll(finalDir, 0777); err != nil {
			panic(err)
		}
	}

	finalFile := finalDir + "/gin.log"

	f, err := os.Create(finalFile)
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.New()

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		l := &logFormat{
			ClientIp:   param.ClientIP,
			Timestamp:  param.TimeStamp.Format(time.RFC1123),
			Method:     param.Method,
			Path:       param.Path,
			Code:       param.StatusCode,
			UserAgent:  param.Request.UserAgent(),
			Host:       param.Request.Host,
			RemoteAddr: param.Request.RemoteAddr,
			BodySize:   param.BodySize,
			Message:    param.ErrorMessage,
		}

		res, _ := json.Marshal(l)
		return string(res) + "\n"
	}))

	router.Use(gin.Recovery())

	router.LoadHTMLGlob("templates/*")

	construct(router)

	return router
}
