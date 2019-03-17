package message

type MixMessage struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int    `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	MsgId        int    `xml:"MsgId"`

	Content string `xml:"Content"`
}