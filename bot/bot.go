package bot

import (
	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/layout"
)

type Bot struct {
	*tele.Bot

	*layout.Layout
}

func New(path string) (*Bot, error) {
	lt, err := layout.New(path)
	if err != nil {
		return nil, err
	}

	b, err := tele.NewBot(lt.Settings())
	if err != nil {
		return nil, err
	}

	return &Bot{b, lt}, nil
}

func (b *Bot) Run() {
	b.Use(b.Middleware("en"))

	b.registerHandlers()

	b.Bot.Start()
}

func (b *Bot) registerHandlers() {
	b.Handle("/start", b.onStart)

	b.Handle(b.Callback("menu"), b.onMenu)
	b.Handle(b.Callback("id"), b.onID)
	b.Handle(b.Callback("utils"), b.onUtils)
}
