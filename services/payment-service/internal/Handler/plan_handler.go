package Handler

import (
	"encoding/json"
	"strconv"

	//"io"
	"net/http"
	//"strconv"

	"backend/payment-service/internal/Models"
	"backend/payment-service/internal/Service"
)

type PlanHandler struct {
	Service *Service.PlanService
}

func NewGenreHandler(service *Service.PlanService) *PlanHandler {
	return &PlanHandler{Service: service}
}

func (handler *PlanHandler) CreatePlane(writer http.ResponseWriter, request *http.Request) {

	var plan Models.Plan

	if err := json.NewDecoder(request.Body).Decode(&plan); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if plan.Name == "" {
		http.Error(writer, "Plan name is required", http.StatusBadRequest)
		return
	}

	err := handler.Service.CreatePlan(&plan)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(plan)

}

func (handler *PlanHandler) UpdatePlan(w http.ResponseWriter, r *http.Request) {

	// 1️⃣ Get ID from query param
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Plan id is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid Plan id", http.StatusBadRequest)
		return
	}

	// 2️⃣ Decode JSON body
	var plan Models.Plan
	if err := json.NewDecoder(r.Body).Decode(&plan); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 3️⃣ Validate input
	if plan.Name == "" {
		http.Error(w, "Plan name is required", http.StatusBadRequest)
		return
	}

	// 4️⃣ Call service
	err = handler.Service.UpdatePlan(uint(id), &plan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 5️⃣ Success response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Plan updated successfully",
	})
}

func (handler *PlanHandler) DeletePlan(writer http.ResponseWriter, request *http.Request) {
	id, _ := strconv.Atoi(request.URL.Query().Get("id"))

	if err := handler.Service.DeletePlan(uint(id)); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(map[string]string{})
}

func (handler *PlanHandler) GetAllPlan(writer http.ResponseWriter, request *http.Request) {
	plan, err := handler.Service.GetAllPlan()

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(plan)
}
