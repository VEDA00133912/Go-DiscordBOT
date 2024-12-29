package autoreply

import (
	"github.com/bwmarrin/discordgo"
)

func HelloReply(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	if m.Content == "こんにちは" {
		reply := "こんにちは！"
		s.ChannelMessageSendReply(m.ChannelID, reply, m.Reference())
		return true
	}
	return false
}