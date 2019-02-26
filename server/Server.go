package server

import (
	"github.com/bwmarrin/discordgo"
	"github.com/moedevs/Vigne/database"
)

type Server struct {
	Modules map[string]Module
	Session *discordgo.Session
	Database *database.Database
	config *database.Config
}

func NewServer(identifier, address, password string) (*Server, error) {
	//Create database
	var err error
	s := Server{}
	s.Modules = make(map[string]Module)
	s.Database = database.NewDatabase(identifier, address, password)
	//Get config from database
	config := s.Database.Config()
	if err != nil {
		return nil, err
	}
	s.Session, err = discordgo.New(config.Token())
	if err != nil {
		return nil, err
	}
	s.Session.ShouldReconnectOnError = true
	s.Session.StateEnabled = true
	return &s, nil
}

func (s Server) Start() error {
	err := s.Session.Open()
	if err != nil {
		return err
	}
	//TODO: Find a better method for waiting
	<- make(chan bool)
	return nil
}