package processer

import (
	"encoding/xml"
	"github.com/superboy724/wechatmessage/message"
	"github.com/superboy724/wechatmessage/response"
)

type MessageProcesser struct {
}

func (t *MessageProcesser) GetRequest(values map[string]string) string {
	return ""
}

func (t *MessageProcesser) PostRequest(values map[string]string, body []byte) string {
	v := &message.MixMessage{}
	err := xml.Unmarshal(body, v)
	if err != nil {
		return ""
	}
	//test code
	resp := response.NewTextResponse("hello world,your content:"+v.Content, v.FromUserName)
	if resStr, err := resp.TextResponseToString(); err == nil {
		return resStr
	} else {
		return err.Error()
	}
}

func NewMessageProcesser() *MessageProcesser {
	res := &MessageProcesser{}

	return res
}
