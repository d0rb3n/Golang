package app

import (
	"context"
	"log"
	"net/http"

	"time"

	"practice4/internal/handler"
	"practice4/internal/middleware"
	"practice4/internal/repository"
	_mysql "practice4/internal/repository/mysql"
	"practice4/internal/usecase"
	"practice4/pkg/modules"
)

func Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 1. Подключаемся к БД
	dbConfig := initMySQLConfig()
	db := _mysql.NewMySQLDialect(ctx, dbConfig)

	// 2. Создаём слои: Repository → Usecase → Handler
	repos := repository.NewRepositories(db)
	userUsecase := usecase.NewUserUsecase(repos.UserRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	// 3. Роутер
	mux := http.NewServeMux()

	// Healthcheck (без авторизации)
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "ok"}`))
	})

	// User endpoints
	mux.HandleFunc("GET /users", userHandler.GetUsers)
	mux.HandleFunc("GET /users/{id}", userHandler.GetUserByID)
	mux.HandleFunc("POST /users", userHandler.CreateUser)
	mux.HandleFunc("PUT /users/{id}", userHandler.UpdateUser)
	mux.HandleFunc("DELETE /users/{id}", userHandler.DeleteUser)

	// 4. Оборачиваем в middleware (порядок важен!)
	// Auth → Logging → Handler
	chain := middleware.LoggingMiddleware(middleware.AuthMiddleware(mux))

	// 5. Запускаем сервер
	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", chain); err != nil {
		log.Fatal(err)
	}
}

func initMySQLConfig() *modules.MySQLConfig {
	return &modules.MySQLConfig{
		Host:        "localhost",
		Port:        "3306",
		Username:    "root",
		Password:    "darkhan1709",
		DBName:      "mydb",
		ExecTimeout: 5 * time.Second,
	}
}
