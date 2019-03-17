package processer

type Processer interface {
	GetRequest(map[string]string) string          //接受http get参数，返回字符串
	PostRequest(map[string]string, []byte) string //接受http post body,同时可接受url value，返回字符串
}
