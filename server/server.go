package server

import (
	"fmt"
	"github.com/superboy724/wechatmessage/processer"
	"net/http"
	"strconv"
)

type Server struct {
	port      int
	processer processer.Processer
}

func (t *Server) Run() {
	portStr := strconv.Itoa(t.port)
	http.HandleFunc("/", t.read)
	http.ListenAndServe(":"+portStr, nil)
}

func (t *Server) SetProcesser(p processer.Processer) {
	t.processer = p
}

func NewServer(port int) *Server {
	server := &Server{
		port: port,
	}

	return server
}

func (t *Server) read(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == "GET" {
		r.ParseForm()
		values := make(map[string]string, len(r.Form))
		for key, value := range r.Form {
			values[key] = value[0]
		}
		res := t.processer.GetRequest(values)
		fmt.Fprintln(w, res)
	}
}
