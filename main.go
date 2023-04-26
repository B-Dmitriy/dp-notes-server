package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Tag struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Tags struct {
	Tags []Tag `json:"tags"`
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: ping")
	fmt.Fprintf(w, "Ping!")
}

func usersHandler(writer http.ResponseWriter, request *http.Request) {
	// Добавляем заголовок ответа
	writer.Header().Add("Content-Type", "application/json")
	// Читаем файл
	jsonFile, _ := os.ReadFile("data.json")
	// Создаём переменную для Tags
	var result Tags
	//
	json.Unmarshal(jsonFile, &result)

	json.NewEncoder(writer).Encode(result.Tags)
}

func main() {
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/tags", usersHandler)

	err := http.ListenAndServe("localhost:8011", nil)
	log.Fatal(err)
}
