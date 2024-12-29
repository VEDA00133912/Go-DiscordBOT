package autoreply

import (
	"github.com/bwmarrin/discordgo"
)

func GoodNightReply(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	if m.Content == "おやすみ" {
		reply := "おやすみなさい！"
		s.ChannelMessageSendReply(m.ChannelID, reply, m.Reference())
		return true
	}
	return false
}