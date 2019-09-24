package levels

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/moedevs/Vigne/commands"
	"github.com/moedevs/Vigne/messages"
	"math"
	"strconv"
	"time"
)

type LevelCommand struct {
	m *LevelsModule
}

func (l LevelCommand) Check(cmd string) bool {
	return cmd == "level" || cmd == "levels"
}

func (l LevelCommand) Action(m *discordgo.MessageCreate, args []string, creator *messages.MessageCreator) error {
	message := creator.NewMessage()
	message.SetExpiry(time.Second*30)
	embed := message.GetEmbedBuilder()
	level, err := l.m.GetLevels(m.Author.ID)
	if err != nil {
		return err
	}
	levelW, levelF := math.Modf(level)
	embed.SetTitle(fmt.Sprintf("%d%% until next level", int(100-levelF*100)))
	embed.SetColor(0xbc2ff0)
	embed.SetThumbnail(m.Author.AvatarURL(""))
	field1 :=embed.NewField()
	field1.SetName("Level")
	field1.SetValue(strconv.Itoa(int(levelW)))
	return nil
}

func (l LevelCommand) ShouldRemoveOriginal() bool {
	return true
}

func (l LevelCommand) GetHelpPageEntry() commands.HelpPageEntry {
	return commands.HelpPageEntry{
		Description: "Shows how many levels you have",
		Usage:       "",
		Command:     "level",
		IsHidden:    false,
	}
}
