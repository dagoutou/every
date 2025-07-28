package cg

import (
	"encoding/json"
	"errors"
	"every/src/ipc"
	"sync"
)

var _ ipc.Server = &CenterServer{}

type Message struct {
	Form    string `json:"form"`
	To      string `json:"to"`
	Content string `json:"content"`
}

type CenterServer struct {
	servers map[string]ipc.Server
	players []*Player
	rooms   []*Room
	mutex   sync.RWMutex
}

type Room struct {
}

func NewCenterServer() *CenterServer {
	servers := make(map[string]ipc.Server)
	players := make([]*Player, 0)
	return &CenterServer{
		servers: servers,
		players: players,
		mutex:   sync.RWMutex{},
	}
}

func (c *CenterServer) addPlayer(params string) error {
	player := NewPlayer()
	if err := json.Unmarshal([]byte(params), &player); err != nil {
		return err
	}
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.players = append(c.players, player)
	return nil
}

func (c *CenterServer) broadcast(params string) error {
	var msg Message
	var err error
	if err = json.Unmarshal([]byte(params), msg); err != nil {
		return err
	}
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if len(c.players) > 0 {
		for _, player := range c.players {
			player.Mq <- msg
		}
	} else {
		err = errors.New("no player online")
	}
	return err
}

func (c *CenterServer) Handle(method string, params string) *ipc.Response {
	switch method {
	case "addPlayer":
		if err := c.addPlayer(params); err != nil {
			return &ipc.Response{
				Code: "500",
				Body: err.Error(),
			}
		}
	case "broadcast":
		if err := c.broadcast(params); err != nil {
			return &ipc.Response{
				Code: "500",
				Body: err.Error(),
			}
		}
	default:
		return &ipc.Response{
			Code: "404",
			Body: "No Router",
		}
	}
	return &ipc.Response{
		Code: "200",
		Body: "",
	}
}

func (c *CenterServer) Name() string {
	return "CenterServer"
}
