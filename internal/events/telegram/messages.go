package telegram

const msgHelp = `Я сохраняю ваши показатели сна из Samsung Health. 

Чтобы сохранить показатели сна за день, отправьте в сообщении показатель (число от 1 до 100)

Позже я смогу читать данные со скриншота, если вы купите полную версию программы 😁
`

const msgHello = "Здорова, пёс! 🐶\n\n" + msgHelp

const (
	msgUnknownCommand = "Неизвестная команда 🤔"
	msgSavedSleepRate = "Сохранено! 👌"
)

const (
	msgSaveRateErr = "Ошибка при сохранении показателя сна"
)
