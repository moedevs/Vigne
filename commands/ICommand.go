package commands

import (
	"github.com/moedevs/Vigne/messages"
	"github.com/bwmarrin/discordgo"
)

type ICommand interface {
	Check(cmd string) bool
	Action(m *discordgo.MessageCreate,args []string, creator *messages.MessageCreator) error
	ShouldRemoveOriginal() bool
	GetHelpPageEntry() HelpPageEntry
}
