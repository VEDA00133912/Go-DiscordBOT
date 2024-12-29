package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/bwmarrin/discordgo"

	"GO-DISCORDBOT/function/autoreply"
	"GO-DISCORDBOT/function/commands"
)

type Config struct {
	Prefix string `json:"prefix"`
}

func loadConfig() (*Config, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

func main() {
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("config.jsonの読み込みに失敗しました: %v", err)
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatalf(".envファイルの読み込みに失敗しました: %v", err)
	}

	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		log.Fatalf("DISCORD_TOKENが設定されていません")
	}

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Discordセッションの作成に失敗しました: %v", err)
	}

	dg.AddHandler(messageCreate(config))
	if err := dg.Open(); err != nil {
		log.Fatalf("Discord接続のオープンに失敗しました: %v", err)
	}
	defer dg.Close()

	user, err := dg.User("@me")
	if err != nil {
		log.Fatalf("ユーザー情報の取得に失敗しました: %v", err)
	}

	fmt.Printf("%s起動！CTRL+Cで終了します\n", user.Username)

	select {}
}

func messageCreate(config *Config) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.Bot {
			return
		}

		if autoreply.HelloReply(s, m) {
			return
		}
		if autoreply.GoodNightReply(s, m) {
			return
		}
		
		if strings.HasPrefix(m.Content, config.Prefix) {
			command := strings.TrimPrefix(m.Content, config.Prefix)
			switch command {
			case "about":
				commands.AboutCommand(s, m)
			case "ping":
				commands.PingCommand(s, m)
			}
		} else {
			return
		}
	}
}