package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// PageData - структура, которую мы будем передавать в HTML
type PageData struct {
	Title  string
	Result string
}

func main() {
	// 1. Настраиваем раздачу статических файлов (наш CSS)
	// Это решает главную задачу - "Linking CSS and HTML"
	fs := http.FileServer(http.Dir("static"))

	// Говорим серверу: все запросы, начинающиеся с /static/,
	// перенаправлять в папку static на компьютере
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// 2. Настраиваем обработчик для главной страницы
	http.HandleFunc("/", homeHandler)

	fmt.Println("Сервер запущен. Открой браузер по ссылке: http://localhost:8080")

	// 3. Запускаем сервер на порту 8080.
	// Используем только стандартный пакет net/http
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}

// homeHandler обрабатывает запросы к главной странице "/"
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Если пользователь ввел несуществующий адрес, выдаем 404
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Читаем наш HTML-шаблон
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		// Обработка ошибки, чтобы программа не падала
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Подготавливаем данные для отображения
	data := PageData{
		Title:  "ASCII Art Generator",
		Result: "Тут будет красивый ASCII-арт...",
	}

	// Отправляем HTML с вшитыми данными обратно в браузер пользователя
	tmpl.Execute(w, data)
}
