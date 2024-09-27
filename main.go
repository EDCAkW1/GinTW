package main


import (
	"log"                     	// Пакет для логирования
	"github.com/gin-gonic/gin"	// Веб-фреймворк Gin
	"gin-notes-api/handlers" // Импортирование модуля с обработчиками запросов
	"gin-notes-api/storage"  // Импортирование модуля для работы с базой данных
)

func main() {
	// Инициализация базы данных
	if err := storage.InitDatabase(); err != nil {
    	log.Fatalf("Failed to initialize database: %v", err) // Логирование ошибки и завершение программы, если инициализация базы данных не удалась
	}

	// Создание нового роутера Gin с настройками по умолчанию
	router := gin.Default()

	// Определение маршрутов и привязка их к соответствующим обработчикам
	router.GET("/notes", handlers.GetNotes)         	// Маршрут для получения всех заметок
	router.GET("/notes/:id", handlers.GetNoteByID)  	// Маршрут для получения заметки по ID
	router.POST("/notes", handlers.CreateNote)      	// Маршрут для создания новой заметки
	router.PUT("/notes/:id", handlers.UpdateNoteByID)   // Маршрут для обновления заметки по ID
	router.DELETE("/notes/:id", handlers.DeleteNoteByID) // Маршрут для удаления заметки по ID

	// Запуск веб-сервера на порту 8080
	router.Run(":8080")
	
}

