package Routes

import (
	"backend/payment-service/internal/Handler"
	"net/http"
)

func RegisterPlanRoutes(mux *http.ServeMux, handler *Handler.PlanHandler) {
	mux.HandleFunc("/api/plan/creatPlan", handler.CreatePlane)
	mux.HandleFunc("/api/plan/getAllPlan", handler.GetAllPlan)
	mux.HandleFunc("/api/plan/updatePlan", handler.UpdatePlan)
	mux.HandleFunc("/api/plan/deletePlan", handler.DeletePlan)
}

func RegisterSubsRoutes(mux *http.ServeMux, handler *Handler.SubsHandler) {
	mux.HandleFunc("/api/subs/creatSubs", handler.CreateSubs)
	mux.HandleFunc("/api/subs/getAllSubs", handler.GetAllSubs)
	mux.HandleFunc("/api/subs/updateSubs", handler.UpdateSubs)
	mux.HandleFunc("/api/subs/deleteSubs", handler.DeleteSubs)
}
