package router

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/wuzehv/passport/util/config"
	"io"
	"log"
	"os"
	"path/filepath"
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

	logDir := config.Log.Dir
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err = os.MkdirAll(logDir, 0777); err != nil {
			log.Fatalf("log dir create error: %v\n", err)
		}
	}

	finalFile := logDir + "/gin.log"

	f, err := os.Create(finalFile)
	if err != nil {
		log.Fatalf("log file create error: %v\n", err)
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

	router.LoadHTMLFiles(loadTemplates("template")...)

	construct(router)

	return router
}

func loadTemplates(templatesDir string) []string {
	other, err := filepath.Glob(templatesDir + "/**/*.html")
	if err != nil {
		log.Fatalf("load template error: %v\n", err)
	}

	admin, err := filepath.Glob(templatesDir + "/**/**/*.html")
	if err != nil {
		log.Fatalf("load template error: %v\n", err)
	}

	return append(admin, other...)
}
