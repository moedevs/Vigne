package commands

import (
	"fmt"
	"github.com/moedevs/Vigne/errors"
	"github.com/moedevs/Vigne/messages"
	"github.com/moedevs/Vigne/server"
	"github.com/bwmarrin/discordgo"
	"regexp"
	"strings"
	"time"
)

type CommandsModule struct {
	Server *server.Server
	Regex *regexp.Regexp
	Commands []ICommand
}

func (module CommandsModule) GetName() string {
	return "Commands"
}

func (module *CommandsModule) Init(s *server.Server) error {
	//Get command regex from database
	config := s.Database.Config()
	var err error
	module.Regex, err = regexp.Compile(config.CommandRegex())
	if err != nil {
		return err
	}
	s.Session.AddHandler(module.handleCommands)
	module.Server = s
	return nil
}



func (module *CommandsModule) handleCommands(s *discordgo.Session, m *discordgo.MessageCreate)  {
	//Does command match?
	if module.Regex.MatchString(m.Content) {
		//Get command
		submatches := module.Regex.FindStringSubmatch(m.Content)
		command := submatches[1]
		//Get arguments for commands
		var args []string
		if len(submatches) > 2 {
			args = strings.Split(submatches[2], " ")
		}
		//Cleanup args array
		for i := 0 ; i < len(args); i++ {
			val := args[i]
			//Remove whitespace
			val = strings.TrimSpace(val)
			args[i] = val
			//If argument is empty
			if val == ""{
				args = append(args[:i], args[i+1:]...)
				i--
			}
		}
		//Loop through every possible command
		for _, commandHandler := range module.Commands {
			//Check
			if commandHandler.Check(command){
				//Delete trigger message if necessary
				if commandHandler.ShouldRemoveOriginal() {
					s.ChannelMessageDelete(m.ChannelID, m.ID)
				}
				//Get messages module
				imessages, err := module.Server.GetModuleByName("messages")
				if err != nil {
					return
				}
				messenger := imessages.(*messages.MessagesModule)
				//execute action
				c := messenger.NewMessageCreator(m.ChannelID)
				err = commandHandler.Action(m, args, c)
				//Handle command error
				if err != nil {
					//Check if err is Public error
					publicErr, ok := err.(*errors.PublicError)
					if ok {
						c := messenger.NewMessageCreator(m.ChannelID)
						msg := c.NewMessage()
						msg.SetContent(fmt.Sprintf("<@%s>, %s", m.Author.ID, publicErr.PublicPart))
						msg.SetExpiry(time.Second*10)
						c.Send()
						private := publicErr.Error()
						if private != "" {
							fmt.Printf("%s (%s) has failed to execute %s. Reason: %s\n", m.Author.Username, m.Author.ID, m.Content, private)
						}
					}else {
						fmt.Printf("%s (%s) has failed to execute %s. Reason: %s\n", m.Author.Username, m.Author.ID, m.Content, err)
					}
					return
				}
				err = c.Send()
				if err != nil {
					fmt.Printf("Couldn't send message: %s\n", err)
					return
				}
				break
			}
		}
	}
}

func (module *CommandsModule) RegisterCommand(command ICommand) error {
	module.Commands = append(module.Commands, command)
	return nil
}