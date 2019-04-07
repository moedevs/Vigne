package weebsh

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/moedevs/Vigne/commands"
	"github.com/moedevs/Vigne/messages"
	"time"
)

type TypesCommand struct {
	module *WeebshModule
}

func (TypesCommand) Check(cmd string) bool {
	return cmd == "type" || cmd == "types"
}

func (c TypesCommand) Action(m *discordgo.MessageCreate, args []string, creator *messages.MessageCreator) error {
	types, err := c.module.GetTypes()
	if err != nil {
		return err
	}
	msg := creator.NewMessage()
	embed := msg.GetEmbedBuilder()
	embed.SetTitle("Image types")
	embed.SetFooter("Use with --image", "")
	desc := ""
	for _, typeName := range types {
		desc += fmt.Sprintf("`%s` ", typeName)
	}
	embed.SetDescription(desc)
	msg.SetExpiry(time.Second*30)
	return nil
}

func (TypesCommand) ShouldRemoveOriginal() bool {
	return true
}

func (TypesCommand) GetHelpPageEntry() commands.HelpPageEntry {
	return commands.HelpPageEntry{
		Command: "types",
		Description: "List all types that can be used with --image",
	}
}

