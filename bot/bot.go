package bot

import (
	"strings"

	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/layout"
	"gopkg.in/telebot.v3/middleware"
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
	b.Use(middleware.Logger())
	b.Use(b.Middleware("en"))

	b.registerHandlers()

	b.Bot.Start()
}

// onInline is the handler for inline queries. it is a kind of router for other inline handlers.
func (b *Bot) onInline(c tele.Context) error {
	command, _ := b.parseQuery(c.Data())

	switch command {
	case "s", "short", "shrt":
		return b.onShortenerInline(c)
	}

	return b.onHelp(c)
}

// parseQuery parses the query text and returns the command and the data.
func (b Bot) parseQuery(text string) (command, data string) {
	parts := strings.Split(text, " ")
	if len(parts) >= 1 {
		command = strings.ToLower(parts[0])
		data = strings.Join(parts[1:], " ")
	}
	return
}

// registerHandlers registers all the handlers.
func (b *Bot) registerHandlers() {
	b.Handle("/start", b.onStart)

	b.Handle(b.Callback("menu"), b.onMenu)
	b.Handle(b.Callback("id"), b.onID)
	b.Handle(b.Callback("utils"), b.onUtils)

	b.Handle(tele.OnQuery, b.onInline)
	b.Handle(tele.OnInlineResult, b.onPostShortenerInline)
}
