package webchat

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var routes []Route

type Route struct {
	Method   string
	Path     string
	Handlers []gin.HandlerFunc
}

func Add(r Route) {
	routes = append(routes, r)
}

/**
分组路由
user := router.Group("user", gin.Logger(),gin.Recovery())
{
    user.GET("info", func(context *gin.Context) {

    })
    user.GET("article", func(context *gin.Context) {

    })
}


过滤器
    router.PUT("/users", handler.MyHandler(), controller.UpdateUser)
    router.DELETE("/users/:name", handler.MyHandler(), controller.DeleteUser)

*/
func init() {
	Add(Route{
		Method: "GET",
		Path:   "/webchat",
		Handlers: []gin.HandlerFunc{CheckToken(), func(context *gin.Context) {
			//context.String(http.StatusOK, "pong")
			echostr := context.Query("echostr")
			_, _ = context.Writer.WriteString(echostr)
		}},
	})

	Add(Route{
		Method: "POST",
		Path:   "/webchat",
		Handlers: []gin.HandlerFunc{CheckToken(), func(context *gin.Context) {
			var textMsg WXTextMsg
			err := context.ShouldBindXML(&textMsg)
			if err != nil {
				log.Printf("[消息接收] - XML数据包解析失败: %v\n", err)
				return
			}

			log.Printf("[消息接收] - 收到消息, 消息类型为: %s, 消息内容为: %s\n", textMsg.MsgType, textMsg.Content)
			WXMsgReply(context, textMsg.ToUserName, textMsg.FromUserName)
		}},
	})

	Add(Route{
		Method: "GET",
		Path:   "/index",
		Handlers: []gin.HandlerFunc{func(context *gin.Context) {
			context.HTML(http.StatusOK, "index.html", gin.H{
				"title":  "DEMO",
				"movies": routes,
			})
		}},
	})
}

func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		if WXCheckSignature(c) == false {
			c.JSON(http.StatusForbidden, gin.H{"error": "token not found"})
			c.Abort()
		} else {
			c.Next()
		}
	}
}

func Create(engine *gin.Engine) {
	engine.LoadHTMLGlob("html/template/*")
	for _, route := range routes {
		if route.Method == "GET" {
			engine.GET(route.Path, route.Handlers...)
		}
		if route.Method == "POST" {
			engine.POST(route.Path, route.Handlers...)
		}
		if route.Method == "PUT" {
			engine.PUT(route.Path, route.Handlers...)
		}
		if route.Method == "DELETE" {
			engine.DELETE(route.Path, route.Handlers...)
		}
	}
}
