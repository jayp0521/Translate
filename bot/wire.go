package bot

import (
	"os"
	"sync"
	"time"

	"github.com/google/wire"
	"github.com/jayp0521/Translate/utils"
	"go.uber.org/zap"
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

func provideBot(key botKey, log *zap.SugaredLogger) TBot {
	pref := tele.Settings{
		Token:   string(provideBotKey()),
		Poller:  &tele.LongPoller{Timeout: 10 * time.Second},
		Verbose: true,
	}
	b, err := tele.NewBot(pref)
	if err != nil {
		panic(err)
	}
	return TBot{
		log: log,
		key: key,
		Bot: b,
	}
}

func ProvideBot() TBot {
	botInit.Do(func() {
		bot = provideBot(provideBotKey(), utils.ProvideLogger())
	})
	return bot
}
