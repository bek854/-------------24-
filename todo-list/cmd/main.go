package main

import (
    "fmt"
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/bek854/todo-list/database"
    "github.com/bek854/todo-list/handlers"
)

func main() {
    database.ConnectDB()

    r := chi.NewRouter()

    // Добавляем корневой маршрут
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("TODO List API is working! Use /tasks endpoints"))
    })

    // Добавляем health check
    r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("OK"))
    })

    // Регистрируем handlers для задач
    r.Get("/tasks", handlers.GetTasks)
    r.Post("/tasks", handlers.PostTask)
    r.Get("/tasks/{id}", handlers.GetTaskByID)
    r.Delete("/tasks/{id}", handlers.DeleteTask)

    fmt.Println("Server starting on :3000")
    if err := http.ListenAndServe(":3000", r); err != nil {
        fmt.Printf("Start server error: %s", err.Error())
    }
}
