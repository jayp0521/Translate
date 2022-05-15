package main

import (
	"github.com/jayp0521/Translate/bot"
	"github.com/jayp0521/Translate/env"
	"github.com/jayp0521/Translate/utils"
	tele "gopkg.in/telebot.v3"
)

func main() {
	// path, _ := os.Getwd()
	// key := encrypt.GenerateKey()
	// fmt.Println(key)
	// _, err := encrypt.EncryptFile(filepath.Join(path, ".env.production"), key)
	// if err != nil {
	// 	panic(err)
	// }
	log := utils.ProvideLogger()
	loaded := env.InjectEnvLoad()
	err := loaded.Load()
	if err != nil {
		panic(err)
	}

	b := bot.ProvideBot().Bot
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	b.Start()
}
