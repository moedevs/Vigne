package weebsh

import (
	"github.com/bwmarrin/discordgo"
	"github.com/moedevs/Vigne/commands"
	"github.com/moedevs/Vigne/errors"
	"github.com/moedevs/Vigne/messages"
)

type ImageCommand struct {
	module *WeebshModule
	predicate string
}

func (c *ImageCommand) Check(cmd string) bool {
	c.predicate = ""
	types, err := c.module.GetTypes()
	if err == nil {
		for _, typeName := range types {
			if cmd == typeName {
				c.predicate = cmd
				return true
			}
		}
	}
	return cmd == "image"
}

func (c *ImageCommand) Action(m *discordgo.MessageCreate, args []string, creator *messages.MessageCreator) error {
	typeName := ""
	if c.predicate != "" {
		typeName = c.predicate
	}else if len(args) >= 1 {
		typeName = args[0]
	}else {
		return errors.New("", "Please provide a type")
	}
	image, err := c.module.GetImageByType(typeName)
	if err != nil {
		return err
	}
	msg := creator.NewMessage()
	embed := msg.GetEmbedBuilder()
	//msg.SetContent("<@" + m.Author.ID + ">")
	embed.SetAuthor(m.Author.Username + " sent an image.",
		"",
		"https://cdn.discordapp.com/avatars/" + m.Author.ID + "/" + m.Author.Avatar + ".webp?size=128")
	embed.SetFooter("Provided by weeb.sh", "https://docs.weeb.sh/images/logo.png")
	embed.SetImage(image)
	embed.SetColor(0x00897B)
	return nil
}

func (c *ImageCommand) ShouldRemoveOriginal() bool {
	return false
}

func (c *ImageCommand) GetHelpPageEntry() commands.HelpPageEntry {
	return commands.HelpPageEntry{
		IsHidden: false,
		Description: "Sends a random image by type",
		Command:"image",
		Usage:"[type]",
	}
}
