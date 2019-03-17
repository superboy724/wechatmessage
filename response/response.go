package response

import (
	"encoding/xml"
)

type Response struct {
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int      `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	XMLName      xml.Name `xml:"xml"`
}

func NewResponse(msgType string) Response {
	r := Response{
		MsgType: msgType,
	}

	return r
}

func ResponseToString(resp interface{}) (string, error) {
	res, err := xml.Marshal(resp)
	if err != nil {
		return "", err
	}
	return string(res), nil
}
