package ipc

import (
	"log"
	"testing"
)

type EchoServer struct {
}

func (e *EchoServer) Name() string {
	return "ECHO SERVER"
}

func (e *EchoServer) Handle(method, params string) *Response {
	log.Println("method:", method, "params:", params)
	return &Response{
		Body: "",
		Code: "200",
	}
}

func TestNewIPCServer(t *testing.T) {
	ipcS := NewIPCServer(&EchoServer{})
	
	cl1 := NewIPCClient(ipcS)
	cl2 := NewIPCClient(ipcS)
	cl3 := NewIPCClient(ipcS)
	
	resp1, err := cl1.Call("GET", "client1")
	if err != nil {
		log.Println("cl1.Call error:", err)
		return
	}
	if resp1.Body != "client1" {
		t.Error("IpcClient.Call failed. resp1:", resp1)
	}
	
	resp2, err := cl2.Call("GET", "client2")
	if err != nil {
		log.Println("cl2.Call error:", err)
		return
	}
	if resp2.Body != "client2" {
		t.Error("IpcClient.Call failed. resp2:", resp2)
	}
	
	resp3, err := cl3.Call("GET", "client3")
	if err != nil {
		log.Println("cl3.Call error:", err)
		return
	}
	if resp3.Body != "client3" {
		t.Error("IpcClient.Call failed. resp3:", resp3)
	}
	cl1.Close()
	
	// resp1, err = cl1.Call("GET", "client1")
	// if err != nil {
	// 	log.Println("cl1.Call error:", err)
	// 	return
	// }
	// if resp1.Body != "client1" {
	// 	t.Error("IpcClient.Call failed. resp1:", resp1)
	// }
}
