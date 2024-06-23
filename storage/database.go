package storage

import (
	// Импортирование необходимых пакетов
	"gorm.io/driver/sqlite" // Драйвер для работы с SQLite
	"gorm.io/gorm"      	// Основная библиотека GORM для работы с ORM
	"gin-notes-api/models" // Импортирование пакета с определением моделей данных
)

// Объявление глобальной переменной для хранения экземпляра базы данных
var db *gorm.DB

// Функция инициализации базы данных
func InitDatabase() error {
	var err error
	// Открытие подключения к базе данных SQLite с использованием GORM
	db, err = gorm.Open(sqlite.Open("notes.db"), &gorm.Config{})
	if err != nil {
    	// Возвращение ошибки, если подключение не удалось
    	return err
	}
	// Автоматическое создание таблицы для модели Note, если она еще не существует
	return db.AutoMigrate(&models.Note{})
}

// Функция для получения экземпляра базы данных
func GetDB() *gorm.DB {
	// Возвращение глобальной переменной db, содержащей подключение к базе данных
	return db
}



