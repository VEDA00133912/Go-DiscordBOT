# Go-DiscordBOT
discord.goの簡単なBOT
# 起動手順
必要なパッケージのインストール
```
go get github.com/bwmarrin/discordgo github.com/joho/godotenv
```
config.jsonでプレフィックス設定
```
{
  "prefix": "!"
}
```
envファイルにTOKENを設定
```
DISCORD_TOKEN= ここに自分のBOTのTOKENをいれる
```
起動
```
go run main.go
```
