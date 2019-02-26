package debug

import (
	"github.com/moedevs/Vigne/commands"
	"github.com/moedevs/Vigne/server"
)

type DebugModule struct {

}

func (DebugModule) GetName() string {
	return "debug"
}

func (DebugModule) Init(server *server.Server) error {
	cmdi, err := server.GetModuleByName("Commands")
	if err != nil {
		return err
	}
	cmd := cmdi.(*commands.CommandsModule)
	cmd.RegisterCommand(RolesCommand{server:server})
	cmd.RegisterCommand(TestReplacingCommand{})
	return nil
}


