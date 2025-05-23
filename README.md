# 🏥 Medication Scheduler Bot (Go + Telegram)

Умный планировщик для управления приёмом лекарств с интеграцией в Telegram. Написан на Go с использованием модуля Gotel.

## ✨ Особенности

- 🕒 **Напоминания о приёме** лекарств через Telegram
- 💊 Учёт текущих препаратов с дозировкой и графиком
- 📅 Гибкое расписание (ежедневно, по дням недели, разово)
- 🔔 Настройка пользовательских уведомлений
- 📊 История приёмов с статистикой
- ➕ Простое добавление новых препаратов

## 🚀 Быстрый старт

### Требования
- Go 1.21+
### Установка

1. Клонировать репозиторий:
```bash
git clone https://github.com/yourusername/MedicationScheduler.git
cd MedicationScheduler
```

2. Установить зависимости:
```bash
go mod download
```
3. Настройте .env файл на основе:
```ini
TELEGRAM_BOT_TOKEN=your_bot_token_here
DATABASE_URL=postgres://user:password@localhost:5432/medications?sslmode=disable
```
При наличии готового шаблона используйте:
```bash
cp .env.example .env
```
4.  Запустить приложение:
```bash
go run main.go
```