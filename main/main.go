package main

import (
	"github.com/cpalone/gobot/config"
	"github.com/cpalone/gobot/handlers"
	"github.com/cpalone/maimai.v2"
)


func main() {
	b, err := config.BotFromCfgFile("test.yml")
	if err != nil {
		panic(err)
	}
	for _, room := range b.Rooms {
		room.Handlers = append(room.Handlers, 
		&maimai.LinkTitleHandler{},
		 &handlers.UptimeHandler{}, 
		&handlers.PongHandler{},
		&handlers.HelpHandler{ShortDesc:"MaiMai provides titles for links as well as a base set of commands. Try !help @MaiMai.", LongDesc: "MaiMai provides link titles as well as some basic commands. !ping will prompt a response of 'pong!', as will !ping @MaiMai. !uptime or !uptime @MaiMai will tell you how many hours the bot has been up."})
	}
	b.RunAllRooms()
}