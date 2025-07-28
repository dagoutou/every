package ipc

import (
	"encoding/json"
	"log"
)

type Request struct {
	Method string `json:"method"`
	Params string `json:"params"`
}

type Response struct {
	Code string `json:"code"`
	Body string `json:"body"`
}

type Server interface {
	Name() string
	Handle(method, params string) *Response
}

type IPCServer struct {
	Server
}

func NewIPCServer(server Server) *IPCServer {
	return &IPCServer{
		Server: server,
	}
}

func (s *IPCServer) Connection() chan string {
	session := make(chan string)
	go func(c chan string) {
		for {
			request := <-c
			if request == "close" {
				close(c)
				break
			}
			var req Request
			if err := json.Unmarshal([]byte(request), &req); err != nil {
				log.Println("Invalid request format:", request)
			}
			response := s.Handle(req.Method, req.Params)
			b, err := json.Marshal(response)
			if err != nil {
				log.Println("json.Marshal response error:", err)
			}
			c <- string(b)
		}

	}(session)
	log.Println("A new session has been created successfully.")
	return session
}
