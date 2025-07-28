package ipc

import "encoding/json"

type IPCClient struct {
	conn chan string
}

func NewIPCClient(server *IPCServer) *IPCClient {
	c := server.Connection()
	return &IPCClient{c}
}

func (s *IPCClient) Call(method, params string) (*Response, error) {
	req := &Request{method, params}
	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	s.conn <- string(b)

	resp := <-s.conn

	var respData *Response
	if err = json.Unmarshal([]byte(resp), &respData); err != nil {
		return nil, err
	}
	return respData, nil
}

func (s *IPCClient) Close() {
	s.conn <- "close"
}
