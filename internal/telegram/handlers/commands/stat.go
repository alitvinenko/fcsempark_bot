package commands

import tele "gopkg.in/telebot.v3"

func StatHandler(c tele.Context) error {

	var menu = &tele.ReplyMarkup{
		//ResizeKeyboard: true,
		InlineKeyboard: [][]tele.InlineButton{{tele.InlineButton{
			Text: "Сезон 21/22",
			URL:  "https://clck.ru/fNhzi",
		}}, {tele.InlineButton{
			Text: "Сезон 22/23",
			URL:  "https://clck.ru/32BuoH",
		}}, {tele.InlineButton{
			Text: "Сезон 23/24",
			URL:  "https://clck.ru/36ZVCG",
		}}},
	}

	return c.Send("📊 Статистика ФК Семёновский парк:", menu)
}
