package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // Каждый запрос вызывает обработчик
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// Обработчик возвращает компонент пути из URL запроса,
func handler(w http.ResponseWriter, r *http.Request) {
	var resp string
	fmt.Fprintf(resp, "URL.Path = %q\n", r.URL.Path)
	w = formatter(resp)
}

func formatter(input string) string {
	return input
}
