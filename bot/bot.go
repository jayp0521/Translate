package bot

import (
	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
)

type botKey string

type TBot struct {
	log *zap.SugaredLogger
	key botKey
	Bot *tele.Bot
}
