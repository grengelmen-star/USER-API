package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", homeHandler) // регистрация обработчика, в скобках первое значение - адресс, второй - функция обработчик
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/time", timeHandler)

	http.ListenAndServe(":8080", nil)
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
