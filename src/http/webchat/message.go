package webchat

import (
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type WXTextMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
	MsgId        int64
}


type WXRepTextMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
	// 若不标记XMLName, 则解析后的xml名为该结构体的名称
	XMLName      xml.Name `xml:"xml"`
}

func WXMsgReply(c *gin.Context, fromUser string, toUser string) {
	fmt.Println("!!!!!!!!!!!!!!!!!!!!! 开始回复消息！！！！！！！！！！！！！！！！")
	repTextMsg := WXRepTextMsg{
		ToUserName:   toUser,
		FromUserName: fromUser,
		CreateTime:   time.Now().Unix(),
		MsgType:      "text",
		Content:      fmt.Sprintf("[消息回复] 收到%s的来信，由%s回复 - %s", fromUser, toUser ,time.Now().Format("2021-11-02 15:04:05")),
	}

	msg, err := xml.Marshal(&repTextMsg)
	if err != nil {
		log.Printf("[消息回复] - 将对象进行XML编码出错: %v\n", err)
		return
	}
	_, _ = c.Writer.Write(msg)
}

