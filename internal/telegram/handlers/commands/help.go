package commands

import (
	tele "gopkg.in/telebot.v3"
)

func HelpHandler(c tele.Context) error {
	const message = `Привет! Я - бот-помощник ФК Семёновский парк ⚽️.

Доступные команды:

🪧 /rules - правила чата, респисание игр и т.п.

📊 /stat - ссылки на статистику игроков за разные сезоны.

ℹ️ /help - ты сейчас здесь.

Еще я автоматически создаю голосования на будущие игры и закрываю эти голосования когда набирается необходимое число игроков.`

	return c.Send(message, &tele.SendOptions{ParseMode: tele.ModeMarkdown})
}
