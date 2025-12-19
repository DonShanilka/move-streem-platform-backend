package Handler

import (
	"encoding/json"
	"strconv"

	//"io"
	"net/http"
	//"strconv"

	"github.com/DonShanilka/genres-service/internal/Models"
	"github.com/DonShanilka/genres-service/internal/Service"
)

type GenreHandler struct {
	Service *Service.GenreService
}

func NewGenreHandler(service *Service.GenreService) *GenreHandler {
	return &GenreHandler{Service: service}
}

func (handler *GenreHandler) CreateGenre(writer http.ResponseWriter, request *http.Request) {

	var genre Models.Genre

	if err := json.NewDecoder(request.Body).Decode(&genre); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if genre.Name == "" {
		http.Error(writer, "Genre name is required", http.StatusBadRequest)
		return
	}

	err := handler.Service.CreateGenre(&genre)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(genre)

}

func (handler *GenreHandler) UpdateGenre(w http.ResponseWriter, r *http.Request) {

	// 1️⃣ Get ID from query param
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Genre id is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid genre id", http.StatusBadRequest)
		return
	}

	// 2️⃣ Decode JSON body
	var genre Models.Genre
	if err := json.NewDecoder(r.Body).Decode(&genre); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 3️⃣ Validate input
	if genre.Name == "" {
		http.Error(w, "Genre name is required", http.StatusBadRequest)
		return
	}

	// 4️⃣ Call service
	err = handler.Service.UpdateGenre(uint(id), &genre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 5️⃣ Success response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Genre updated successfully",
	})
}

func (handler *GenreHandler) DeleteGenre(writer http.ResponseWriter, request *http.Request) {
	id, _ := strconv.Atoi(request.URL.Query().Get("id"))

	if err := handler.Service.DeleteGenre(uint(id)); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(map[string]string{})
}

func (handler *GenreHandler) GetAllGenres(writer http.ResponseWriter, request *http.Request) {
	genres, err := handler.Service.GetAllGenres()

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(genres)
}
