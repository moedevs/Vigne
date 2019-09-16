package joined

import (
	"github.com/bwmarrin/discordgo"
	"github.com/moedevs/Vigne/commands"
	"github.com/moedevs/Vigne/messages"
	"time"
)

type JoinedCommand struct {
	m *JoinedModule
}

func (c JoinedCommand) GetHelpPageEntry() commands.HelpPageEntry {
	return commands.HelpPageEntry{
		Description: "See when you joined the server",
		Command:"joined",
	}
}

func (c JoinedCommand) ShouldRemoveOriginal() bool {
	return true
}

func (JoinedCommand) Check(command string) bool {
	return command == "joined"
}

func (c *JoinedCommand) Action(m *discordgo.MessageCreate, args []string, creator *messages.MessageCreator) error {
	msg := creator.NewMessage()
	msg.SetExpiry(time.Second*90)
	member, err := c.m.s.Session.GuildMember(m.GuildID, m.Author.ID)
	if err != nil {
		return err
	}
	embed := msg.GetEmbedBuilder()

	embed.SetThumbnail(m.Author.AvatarURL(""))
	embed.SetTimestamp(string(member.JoinedAt))
	embed.SetColor(0xbdc3c7)
	joinTime, err := member.JoinedAt.Parse()
	if err != nil {
		return err
	}
	joinTime = joinTime.UTC()
	joinString := joinTime.Format("2006-01-02 15:04:05 MST")
	field := embed.NewField()
	field.SetName("Joined at")
	field.SetValue(joinString)

	return nil
}