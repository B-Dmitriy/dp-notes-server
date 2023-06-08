package transport

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Tag struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Tags struct {
	Tags []Tag `json:"tags"`
}

func getTagsList(writer http.ResponseWriter, request *http.Request) error {
	// Добавляем заголовок ответа
	writer.Header().Add("Content-Type", "application/json")
	// Читаем файл
	jsonFile, _ := os.ReadFile("../../data/data.json")
	// Создаём переменную для Tags
	var result Tags

	err := json.Unmarshal(jsonFile, &result)
	if err != nil {
		return err
	}

	err = json.NewEncoder(writer).Encode(result.Tags)
	if err != nil {
		return err
	}

	return nil
}

func getTagById(writer http.ResponseWriter, request *http.Request, id int) error {
	// Читаем файл
	jsonFile, _ := os.ReadFile("../../data/data.json")

	// Создаём переменную для Tags
	var tagsList Tags

	err := json.Unmarshal(jsonFile, &tagsList)
	if err != nil {
		return err
	}

	var result Tag

	for i := range tagsList.Tags {
		if tagsList.Tags[i].Id == id {
			result = tagsList.Tags[i]
			break
		}
	}

	if result.Id != 0 { // Добавляем заголовок ответа
		writer.Header().Add("Content-Type", "application/json")
		err = json.NewEncoder(writer).Encode(result)

		if err != nil {
			return err
		}
	} else {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(404)

		_, err = writer.Write([]byte(`{ "message": "Tag not found"}`))
		if err != nil {
			return err
		}
	}

	return nil
}

func TagsHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		id := strings.TrimPrefix(request.URL.Path, "/tags/")
		intId, _ := strconv.Atoi(id)

		if intId != 0 {
			_ = getTagById(writer, request, intId)
			break
		}

		_ = getTagsList(writer, request)
		break
	default:
		fmt.Println("Default")
	}
}
