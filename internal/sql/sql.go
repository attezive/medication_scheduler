// Package sql содержит функции для работы с базой данных PostgreSQL
package sql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "fugs_12"
	dbname   = "medication_db"
)

// CreateTable создает таблицу medications, если она не была ранее инициализирована
// При инициализации заполняет ее по следующему шаблону:
// id SERIAL PRIMARY KEY,
// user_id BIGINT NOT NULL,
// medicament VARCHAR(255) NOT NULL,
// time_interval BIGINT,
// updated_at TIMESTAMP DEFAULT NOW()
func CreateTable() error {
	conn, err := GetConnection() //Подключение
	if err != nil {
		return err
	}
	// SQL-запрос для создания таблицы
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS medications (
	   id SERIAL PRIMARY KEY,
	   user_id BIGINT NOT NULL,
	   medicament VARCHAR(255) NOT NULL,
	   time_interval BIGINT,
	   updated_at TIMESTAMP DEFAULT NOW()
	);
	`

	// Выполняем запрос
	_, err = conn.Exec(context.Background(), createTableSQL)
	if err != nil {
		return fmt.Errorf("ошибка создания таблицы: %v", err)
	}
	return nil
}

// DeleteTable удалеяет таблицу medications, если она существует
func DeleteTable() error {
	conn, err := GetConnection()
	if err != nil {
		return err
	}

	_, err = conn.Exec(context.Background(), "DROP TABLE IF EXISTS medications")
	if err != nil {
		return fmt.Errorf("ошибка удаления таблицы: %v", err)
	}
	return nil
}

// GetConnection возвращает объект подключения к базе данных
func GetConnection() (*pgx.Conn, error) {
	// Формируем строку подключения
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		user, password, host, port, dbname)

	// Подключаемся к базе данных
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("не удалось подключиться к базе: %v", err)
	}
	defer conn.Close(context.Background())
	return conn, nil
}
