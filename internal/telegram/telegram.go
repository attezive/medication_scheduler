// Package telegram реализует взаимодействие с TelegramAPI с помощью Gotel
package telegram

import (
	"MedicationScheduler/internal/medication"
	"MedicationScheduler/internal/user"
	"github.com/attezive/gotel/api/bot"
	"github.com/attezive/gotel/data"
	"strconv"
)

var Bot bot.GotelBot    // Bot содержит структуру Gotel бота для работы с TelegramAPI
var NextMedication bool // NextMedication содержит флаг будет ли следующее сообщение лекарством

// StartCommandHandle обрабатывает сообщение с командой \start
func StartCommandHandle(message *data.Message) {
	user.NewUser(message.Chat) // Создание нового пользователя
}

// AddCommandHandle обрабатывает сообщение с командой \start
// Следующий обработчик будет считывать сообщение как лекарство
func AddCommandHandle(message *data.Message) {
	NextMedication = true
}

// AddMedicationHandle обрабатывает сообщение лекарства и добавляет его в базу планировщика
func AddMedicationHandle(message *data.Message) {
	if NextMedication { //Если нужно добавить
		med, err := medication.MedicationFromString(message.Text)
		if err != nil {
			ErrorNotification(&message.Chat, "Неправильная форма рецепта")
		}
		medication.AddMedication(med) // Добавление в базу данных
		NextMedication = false
	}
}

// ErrorNotification уведомляет пользователя об ошибки в виде msg отправляя в чат с ботом
func ErrorNotification(chat *data.Chat, msg string) {
	Bot.SendMessage(
		&data.SendingEntity{
			ChatId: strconv.FormatInt(chat.Id, 10),
			Value:  msg},
	)
}
