package main

import (
	"fmt"
	"github.com/moedevs/Vigne/commands"
	"github.com/moedevs/Vigne/messages"
	"github.com/moedevs/Vigne/modules/debug"
	"github.com/moedevs/Vigne/modules/help"
	"github.com/moedevs/Vigne/modules/joined"
	"github.com/moedevs/Vigne/modules/levels"
	"github.com/moedevs/Vigne/modules/music"
	"github.com/moedevs/Vigne/modules/ping"
	"github.com/moedevs/Vigne/modules/reactionMenu"
	"github.com/moedevs/Vigne/modules/roles"
	"github.com/moedevs/Vigne/modules/weebsh"
	"github.com/moedevs/Vigne/modules/welcome"
	"github.com/moedevs/Vigne/server"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Print("Creating bot... ")
	s, err := server.NewServer("vigne", "localhost:6379", "")
	if err != nil {
		panic(err)
	}
	fmt.Println("Done!")
	//Service modules
	s.RegisterModule(&commands.CommandsModule{})
	s.RegisterModule(&messages.MessagesModule{})
	//User modules
	s.RegisterModule(&ping.PingModule{})
	s.RegisterModule(&debug.DebugModule{})
	s.RegisterModule(&roles.RolesModule{})
	s.RegisterModule(&welcome.WelcomeModule{})
	s.RegisterModule(&help.HelpModule{})
	s.RegisterModule(&reactionMenu.ReactionModule{})
	s.RegisterModule(&music.MusicModule{})
	s.RegisterModule(&weebsh.WeebshModule{})
	s.RegisterModule(&joined.JoinedModule{})
	s.RegisterModule(&levels.LevelsModule{})
	fmt.Print("Running bot... ")
	err = s.Start()
	if err != nil {
		panic(err)
	}
	fmt.Println("Done!")

}
