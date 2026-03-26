package Routes

import (
	"net/http"

	"github.com/DonShanilka/user-service/internal/Handler"
)

func RegisterUserRoutes(mux *http.ServeMux, handler *Handler.UserHandler) {
	mux.HandleFunc("/api/user/creatUser", handler.CreateUser)
	mux.HandleFunc("/api/user/getAllUsers", handler.GetAllUsers)
	mux.HandleFunc("/api/user/updateUser", handler.UpdateUser)
	mux.HandleFunc("/api/user/deleteUser", handler.DeleteUser)
}
