package joined

import (
	"github.com/moedevs/Vigne/commands"
	"github.com/moedevs/Vigne/server"
)

type JoinedModule struct {
	s *server.Server
}

func (m JoinedModule) GetName() string {
	return "joined"
}

func (m *JoinedModule) Init(s *server.Server) error {
	m.s = s
	//Get command handler module
	cmdInterface, err := s.GetModuleByName("Commands")
	if err != nil {
		return err
	}
	cmd := (cmdInterface).(*commands.CommandsModule)
	err = cmd.RegisterCommand(&JoinedCommand{m})
	if err != nil {
		return err
	}
	return nil
}


