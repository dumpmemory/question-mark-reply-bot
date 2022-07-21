package questionmarkreply

import (
	"github.com/AkinoKaede/question-mark-reply-bot/common"
	"github.com/AkinoKaede/question-mark-reply-bot/features"
	tele "gopkg.in/telebot.v3"
)

var (
	QuestionMarks = []string{"?", "¿", "？"}
)

func OnText(c tele.Context) error {
	text := c.Message().Text

	for _, b := range text {
		if !common.Contains(string(b), QuestionMarks) {
			return nil
		}
	}

	c.Reply(text)
	return nil
}

func init() {
	features.RegisterFeature(tele.OnText, OnText)
}
