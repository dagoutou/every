package cg

import "log"

type Player struct {
	Name string       `json:"name"`
	Leve int          `json:"leve"`
	Exp  int          `json:"exp"`
	Room int          `json:"room"`
	Mq   chan Message `json:"msg"`
}

func NewPlayer() *Player {
	m := make(chan Message, 1024)
	player := &Player{
		Name: "",
		Leve: 0,
		Exp:  0,
		Room: 0,
		Mq:   m,
	}
	go func(p *Player) {
		for {
			msg := <-p.Mq
			log.Println(p.Name, "received message:", msg.Content)
		}
	}(player)
	return player
}
