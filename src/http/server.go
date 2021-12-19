package http

import (
	"github.com/gin-gonic/gin"
	"webchartweb/src/http/webchat"
)

func Start() {
	//gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()

	webchat.Create(engine)
	//静态html页面的工程内路径
	engine.Static("/html", "html")

	//服务监听的端口
	engine.Run(":9311")
}
