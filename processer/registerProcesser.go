package processer

import (
	"fmt"
	"github.com/superboy724/wechatmessage/utils"
	"sort"
	"strings"
)

type RegisterProcesser struct {
}

func (t *RegisterProcesser) GetRequest(values map[string]string) string {
	var signature, timestamp, nonce, echostr string
	if _, ok := values["signature"]; !ok {
		fmt.Println("error")
		return ""
	}
	if _, ok := values["timestamp"]; !ok {
		fmt.Println("error")
		return ""
	}
	if _, ok := values["nonce"]; !ok {
		fmt.Println("error")
		return ""
	}
	if _, ok := values["echostr"]; !ok {
		fmt.Println("error")
		return ""
	}

	signature = values["signature"]
	timestamp = values["timestamp"]
	nonce = values["nonce"]
	echostr = values["echostr"]

	tmp := []string{timestamp, nonce, "bbbb"}
	sort.Strings(tmp)
	tmpStr := strings.Join(tmp, "")
	sha1Res := utils.Sha1(tmpStr)

	if sha1Res == signature {
		return echostr
	} else {
		return ":("
	}
}

func (t *RegisterProcesser) PostRequest(values map[string]string, body []byte) string {
	return ""
}

func NewRegisterProcesser() *RegisterProcesser {
	res := &RegisterProcesser{}

	return res
}
