package main

import (
	"bufio"
	"every/src/cg"
	"every/src/ipc"
	"log"
	"os"
	"strconv"
	"strings"
)

var centerClient *cg.CenterClient

func startCentServer() error {
	sever := ipc.NewIPCServer(&cg.CenterServer{})
	client := ipc.NewIPCClient(sever)
	centerClient = &cg.CenterClient{IPCClient: client}
	return nil
}

func Login(args []string) int {
	leve, _ := strconv.Atoi(args[2])
	exp, _ := strconv.Atoi(args[3])
	player := cg.NewPlayer()
	player.Name = args[1]
	player.Leve = leve
	player.Exp = exp
	if err := centerClient.AddPlayer(player); err != nil {
		log.Println("Invalid add player", err)
	}
	return 0
}

func Send(args []string) int {
	message := strings.Join(args[1:], " ")
	if err := centerClient.Broadcast(message); err != nil {
		log.Println("Failed send", err)
	}
	return 0
}

func GetCommandHandlers() map[string]func(args []string) int {
	return map[string]func(args []string) int{
		"login": Login,
		"send":  Send,
	}
}
func main() {
	log.Println("Casual Game Server Solution")
	
	startCentServer()
	
	r := bufio.NewReader(os.Stdin)
	
	handlers := GetCommandHandlers()
	
	for {
		log.Println("Command>")
		b, _, _ := r.ReadLine()
		line := string(b)
		tokens := strings.Split(line, " ")
		if handler, ok := handlers[tokens[0]]; ok {
			ret := handler(tokens)
			if ret != 0 {
				break
			}
		} else {
			log.Println("Unknown command:", tokens[0])
		}
	}
}
