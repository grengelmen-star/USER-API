package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"user-api/models"
	"user-api/storage"
)

func main() {

	testUser := models.User{
		Username: "EgorTester",
		Email:    "test@gmail.com",
	}

	createdUser, err := storage.CreateUser(testUser)
	if err != nil {
		fmt.Println("Ошибка", err)
	} else {
		fmt.Printf("Создан пользователь: %+v\n", createdUser)
	}
	allUsers := storage.GetAllUsers()
	fmt.Printf("Всего пользователей:%d\n", len(allUsers))
	//http.HandleFunc("/", homeHandler) // регистрация обработчика, в скобках первое значение - адресс, второй - функция обработчик
	//http.HandleFunc("/health", healthHandler)
	//http.HandleFunc("/time", timeHandler)
	/*http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			handlers.GetUsersHandler(w, r)
		case "POST":
			handlers.CreateUsersHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Сервер запущен на http://localhost:8080")
	fmt.Println("Доступные эндпоинты:")
	fmt.Println("GET /    - Главная страница")
	fmt.Println("GET /health - Проверка состояния")
	fmt.Println("GET /time   - Текущее время")
	fmt.Println("GET /users  - Все пользователи")
	fmt.Println("GET /users/create  - Создать пользователя ")

	http.ListenAndServe(":8080", nil)*/
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain") // w.Header().Set("1","2"); 1 - имя заголовка 2 - значение заголовка
	fmt.Fprintf(w, "User API v1.0")              // w - куда писать, второй аргумент - строка для отправки
}
func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/health" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	data := map[string]string{"status": "ok"}
	json.NewEncoder(w).Encode(data) // кодировка данных в формат Json
}
func timeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/time" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	now := time.Now()                                        // получаем текущее время
	formattedTime := now.Format(time.RFC3339)                // кодируем его в формат RFC3339(обязательно сохранив в новую переменную, иначе возвращенная строка нигде не сохранится)
	data := map[string]string{"current_time": formattedTime} // создаем карту с ключем current_time и значением formattedTime(значение указывается обязательно !)
	json.NewEncoder(w).Encode(data)                          // кодируем данные в формат json
}
