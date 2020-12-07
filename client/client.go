package client

import (
	"html/template"
	"log"
	"net/http"
)

type dataStruct struct {
	Port int
}

// Data порт для подключения. Я бы вынес в конфик
var Data = dataStruct{
	Port: 8080,
}

// Handler отображение страницы клиента
func Handler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("client/index.html")
	if err != nil {
		log.Printf("client.Handler #1:\nError: %s\n\n", err)
		return
	}

	err = t.Execute(w, Data)
	if err != nil {
		log.Printf("client.Handler #2:\nError: %s\n\n", err)
		return
	}
}
