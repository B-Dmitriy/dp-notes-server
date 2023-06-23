package transport

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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
			writer.WriteHeader(500)
		}
	}

	return nil
}

func createTag(writer http.ResponseWriter, request *http.Request) error {
	// Читаем файл
	jsonFile, _ := os.ReadFile("../../data/data.json")

	// Создаём переменную для Tags
	var tagsList Tags

	err := json.Unmarshal(jsonFile, &tagsList)
	if err != nil {
		return err
	}

	var body Tag

	_ = json.NewDecoder(request.Body).Decode(&body)

	newTag := Tag{
		Id:          int(time.Now().Unix()),
		Title:       body.Title,
		Description: body.Description,
	}

	tagsList.Tags = append(tagsList.Tags, newTag)

	jsonTags, _ := json.Marshal(tagsList)

	err = os.WriteFile("../../data/data.json", jsonTags, 666)
	if err != nil {
		writer.WriteHeader(500)
		return err
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(201)
	err = json.NewEncoder(writer).Encode(newTag)

	if err != nil {
		return err
	}

	return nil
}

func deleteTag(writer http.ResponseWriter, request *http.Request, id int) error {
	// Читаем файл
	jsonFile, _ := os.ReadFile("../../data/data.json")

	// Создаём переменную для Tags
	var tagsList Tags
	var deletedTag Tag
	var newTags Tags

	err := json.Unmarshal(jsonFile, &tagsList)
	if err != nil {
		return err
	}

	for _, v := range tagsList.Tags {
		if v.Id == id {
			deletedTag = v
		} else {
			newTags.Tags = append(newTags.Tags, v)
		}
	}

	jsonNewTags, _ := json.Marshal(newTags)

	err = os.WriteFile("../../data/data.json", jsonNewTags, 666)
	if err != nil {
		writer.WriteHeader(500)
		return err
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(200)
	err = json.NewEncoder(writer).Encode(deletedTag)

	if err != nil {
		return err
	}

	return nil
}

func putTag(writer http.ResponseWriter, request *http.Request, id int) error {
	// Читаем файл
	jsonFile, _ := os.ReadFile("../../data/data.json")

	// Создаём переменную для Tags
	var tagsList Tags
	var editedTag Tag

	err := json.Unmarshal(jsonFile, &tagsList)
	if err != nil {
		return err
	}

	var body Tag

	_ = json.NewDecoder(request.Body).Decode(&body)

	for i, v := range tagsList.Tags {
		if v.Id == id {
			tagsList.Tags[i].Title = body.Title
			tagsList.Tags[i].Description = body.Description
			editedTag.Id = id
			editedTag.Title = body.Title
			editedTag.Description = body.Description
		}
	}

	jsonTags, _ := json.Marshal(tagsList)

	err = os.WriteFile("../../data/data.json", jsonTags, 666)
	if err != nil {
		writer.WriteHeader(500)
		return err
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(201)
	err = json.NewEncoder(writer).Encode(editedTag)

	if err != nil {
		return err
	}

	return nil
}

func TagsHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method)
	switch request.Method {
	case "GET":
		id := strings.TrimPrefix(request.URL.Path, "/api/tags/")
		intId, _ := strconv.Atoi(id)

		if intId != 0 {
			_ = getTagById(writer, request, intId)
			break
		}

		_ = getTagsList(writer, request)
		break
	case "POST":
		_ = createTag(writer, request)
	case "DELETE":
		id := strings.TrimPrefix(request.URL.Path, "/api/tags/")
		intId, _ := strconv.Atoi(id)

		_ = deleteTag(writer, request, intId)
	case "PUT":
		id := strings.TrimPrefix(request.URL.Path, "/api/tags/")
		intId, _ := strconv.Atoi(id)

		_ = putTag(writer, request, intId)
	default:
		fmt.Println("Default")
	}
}
