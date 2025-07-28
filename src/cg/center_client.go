package cg

import (
	"encoding/json"
	"every/src/ipc"
)

type CenterClient struct {
	*ipc.IPCClient
}

func (c *CenterClient) AddPlayer(player *Player) error {
	b, err := json.Marshal(player)
	if err != nil {
		return err
	}
	resp, err := c.Call("addPlayer", string(b))
	if err == nil && resp.Code == "200" {
		return nil
	}
	return err
}

func (c *CenterClient) Broadcast(message string) error {
	m := &Message{
		Content: message,
	}
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	resp, err := c.Call("broadcast", string(b))
	if err == nil && resp.Code == "200" {
		return nil
	}
	return err
}
