package discordbot

import (
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	channelId := ""
	token := ""
	msg := "haha"
	err := ChannelMsg(channelId, msg, token, time.Second*5)
	if err != nil {
		t.Fatal(err)
	}
}
