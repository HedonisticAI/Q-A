package httpserver

import (
	"golangqatestdesu/config"
	"net/http"
)

type HttpServer struct {
	Server *http.ServeMux
	Port   string
}

func NewServer(Config config.Config) *HttpServer {
	Server := http.NewServeMux()
	return &HttpServer{Port: Config.HttpPort, Server: Server}
}

func (S *HttpServer) Map(path string, Handler func(w http.ResponseWriter, r *http.Request)) {
	S.Server.HandleFunc(path, Handler)
}
func (S *HttpServer) Run() error {
	return http.ListenAndServe(S.Port, S.Server)
}
