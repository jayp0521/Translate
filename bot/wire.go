//go:build wireinject
// +build wireinject

package bot

import (
	"os"
	"sync"
	"time"

	"github.com/google/wire"
	"github.com/jayp0521/Translate/utils"
	tele "gopkg.in/telebot.v3"
)

var (
	bot     TBot
	botInit sync.Once
)

var SuperSet = wire.NewSet(
	ProvideBot,
)

func provideBotKey() botKey {
	secret := os.Getenv("BOT_KEY")
	if len(secret) == 0 {
		panic("BOT_KEY is not defined!")
	}
	return botKey(secret)
}

func provideTeleBotSettings(key botKey) tele.Settings {
	return tele.Settings{
		Token:   string(key),
		Poller:  &tele.LongPoller{Timeout: 10 * time.Second},
		Verbose: true,
	}
}

func provideTeleBot(settings tele.Settings) *tele.Bot {
	b, err := tele.NewBot(settings)
	if err != nil {
		panic(err)
	}
	return b
}

func ProvideBot() TBot {
	botInit.Do(func() {
		bot = injectBot()
	})
	return bot
}

func injectBot() TBot {
	panic(wire.Build(
		utils.SuperSet,
		provideBotKey,
		provideTeleBotSettings,
		provideTeleBot,
		wire.Struct(new(TBot), "*"),
	))
}
