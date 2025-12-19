package Handler

import (
	"encoding/json"
	"io"
	//"strconv"

	//"io"
	"net/http"
	//"strconv"

	"github.com/DonShanilka/admin-service/internal/Models"
	"github.com/DonShanilka/admin-service/internal/Service"
)

type AdminHandler struct {
	Service *Service.AdminService
}

func NewAdminHandler(service *Service.AdminService) *AdminHandler {
	return &AdminHandler{Service: service}
}

func (handler *AdminHandler) CreateAdmin(writer http.ResponseWriter, request *http.Request) {

	if err := request.ParseMultipartForm(100 << 20); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	admin := Models.Admin{
		Name:     request.FormValue("name"),
		Email:    request.FormValue("email"),
		Password: request.FormValue("password"),
	}

	if file, _, err := request.FormFile("profile_image"); err == nil {
		admin.ProfileImage, _ = io.ReadAll(file)
		file.Close()
	}

	if err := handler.Service.CreateAdmin(&admin); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(map[string]string{
		"message": "Admin created successfully",
	})
}
