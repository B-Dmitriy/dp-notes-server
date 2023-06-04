package transport

import (
	"encoding/json"
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

func TagsHandler(writer http.ResponseWriter, request *http.Request) {
	// Добавляем заголовок ответа
	writer.Header().Add("Content-Type", "application/json")
	// Читаем файл
	jsonFile, _ := os.ReadFile("../../data/data.json")
	// Создаём переменную для Tags
	var result Tags
	//
	json.Unmarshal(jsonFile, &result)
	json.NewEncoder(writer).Encode(result.Tags)
}
