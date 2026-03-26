package Handler

import (
	"backend/payment-service/internal/Models"
	"backend/payment-service/internal/Service"
	"encoding/json"
	"net/http"
	"strconv"
	//"github.com/klauspost/compress/gzhttp/writer"
)

type SubsHandler struct {
	Service *Service.SubsService
}

func NewSubsHandler(service *Service.SubsService) *SubsHandler {
	return &SubsHandler{Service: service}
}

func (handler *SubsHandler) CreateSubs(writer http.ResponseWriter, request *http.Request) {

	var subs Models.Subscription

	if err := json.NewDecoder(request.Body).Decode(&subs); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if subs.UserID == 0 {
		http.Error(writer, "User ID is required", http.StatusBadRequest)
		return
	}

	err := handler.Service.CreateSubs(&subs)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(subs)

}

func (handler *SubsHandler) UpdateSubs(w http.ResponseWriter, r *http.Request) {

	// 1️⃣ Get ID from query param
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Subs id is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid Subs id", http.StatusBadRequest)
		return
	}

	// 2️⃣ Decode JSON body
	var subs Models.Subscription
	if err := json.NewDecoder(r.Body).Decode(&subs); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 3️⃣ Validate input
	if subs.UserID == 0 {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	// 4️⃣ Call service
	err = handler.Service.UpdateSubs(uint(id), &subs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 5️⃣ Success response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Subs updated successfully",
	})
}

func (handler *SubsHandler) DeleteSubs(writer http.ResponseWriter, request *http.Request) {
	id, _ := strconv.Atoi(request.URL.Query().Get("id"))

	if err := handler.Service.DeleteSubs(uint(id)); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(map[string]string{})
}

func (handler *SubsHandler) GetAllSubs(writer http.ResponseWriter, request *http.Request) {
	plan, err := handler.Service.GetAllSubs()

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(plan)
}
