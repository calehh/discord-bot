package discordbot

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type ChannelMstContent struct {
	Content string `json:"content"`
	nonce   string `json:"nonce"`
	TTS     bool   `json:"tts"`
}

func ChannelMsg(channelId string, msg string, token string, timeout time.Duration) error {
	url := fmt.Sprintf("https://discord.com/api/v9/channels/%s/messages", channelId)
	data, err := json.Marshal(ChannelMstContent{
		Content: msg,
		nonce:   fmt.Sprintf("82329451214%v33232234", rand.Intn(1000)),
		TTS:     false,
	})
	if err != nil {
		return err
	}
	headers := map[string]string{
		"Authorization": token,
		"Content-Type":  "application/json",
		"User-Agent":    "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.61 Safari/537.36",
	}
	payload, err := HttpPost(url, data, timeout, headers)
	if err != nil {
		return err
	}
	fmt.Println(string(payload))
	return nil
}
