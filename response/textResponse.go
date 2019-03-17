package response

type TextResponse struct {
	Response
	Content string `xml:"Content"`
}

func NewTextResponse(text string, toUser string) TextResponse {
	res := TextResponse{}

	res.Response = NewResponse("text")
	res.CreateTime = 11155566
	res.FromUserName = "gh_fba62a0ffce7"
	res.ToUserName = toUser
	res.Content = text

	return res
}

func (t TextResponse) TextResponseToString() (string, error) {
	res, err := ResponseToString(t)
	return res, err
}
