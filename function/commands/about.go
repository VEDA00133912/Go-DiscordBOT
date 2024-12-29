package commands

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func AboutCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	botUser, err := s.User("@me")
	if err != nil {
		return
	}

	embed := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("%s [%s#%s]", botUser.Username, botUser.Username, botUser.Discriminator),
		Description: "discord.goで作ったBOTです",
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: botUser.AvatarURL(""),
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("%s || about", botUser.Username+"#"+botUser.Discriminator),
		},
		Timestamp: time.Now().Format(time.RFC3339),
		Color:     0x00BFFF,
	}

	message := &discordgo.MessageSend{
		Content:   "",
		Reference: m.Reference(),
		Embed:     embed,
	}

	s.ChannelMessageSendComplex(m.ChannelID, message)
}