package bot

import (
	tele "gopkg.in/telebot.v3"
)

func (b *Bot) onStart(c tele.Context) error {
	return c.Send(b.Text(c, "start"), b.Markup(c, "start"))
}

func (b *Bot) onUtils(c tele.Context) error {
	return c.Edit(b.Text(c, "utils"), b.Markup(c, "utils"))
}

func (b *Bot) onID(c tele.Context) error {
	var args tele.M = tele.M{
		"ID":     c.Sender().ID,
		"ChatID": c.Chat().ID,
	}
	return c.Edit(b.Text(c, "id", args), b.Markup(c, "menu"))
}

func (b *Bot) onMenu(c tele.Context) error {
	return c.Edit(b.Text(c, "menu"), b.Markup(c, "start"))
}
