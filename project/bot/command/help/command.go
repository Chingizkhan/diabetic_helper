package help

import tele "gopkg.in/telebot.v3"

func Handle(c tele.Context) error {
	const text = `
/poel - Уведомит вас через два часа для проверки сахара
/stat - Сделает анализ ваших сахаров
/help - Показать все доступные комманды
/start - Начать работать с ботом
`
	return c.Send(text)
}
