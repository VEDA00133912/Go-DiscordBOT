package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func PingCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	latency := s.HeartbeatLatency().Milliseconds()
	reply := fmt.Sprintf("🏓 Pong! Latency: `%d`ms", latency)
	s.ChannelMessageSendReply(m.ChannelID, reply, m.Reference())
}
