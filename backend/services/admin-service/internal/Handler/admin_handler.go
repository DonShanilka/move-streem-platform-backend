package Handler

import (
	"encoding/json"
	"io"
	"strconv"

	"net/http"

	"github.com/DonShanilka/admin-service/internal/Models"
	"github.com/DonShanilka/admin-service/internal/Service"
)

type AdminHandler struct {
	Service *Service.AdminService
}

func NewAdminHandler(service *Service.AdminService) *AdminHandler {
	return &AdminHandler{Service: service}
}

func atoiSafe(s string) int {
	i, _ := strconv.Atoi(s)
	return i
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

func (handler *AdminHandler) UpdateAdmin(writer http.ResponseWriter, request *http.Request) {

	err := request.ParseMultipartForm(100 << 20)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	idstr := request.FormValue("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(writer, "Invalid ID", http.StatusBadRequest)
		return
	}

	admin := Models.Admin{
		ID:       uint(id),
		Name:     request.FormValue("name"),
		Email:    request.FormValue("email"),
		Password: request.FormValue("password"),
	}

	if file, _, err := request.FormFile("profile_image"); err == nil {
		admin.ProfileImage, _ = io.ReadAll(file)
		file.Close()
	}

	if err := handler.Service.UpdateAdmin(uint(id), &admin); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(map[string]string{
		"message": "Admin updated successfully",
	})

}

func (handler *AdminHandler) DeleteAdmin(writer http.ResponseWriter, request *http.Request) {
	id, _ := strconv.Atoi(request.URL.Query().Get("id"))

	if err := handler.Service.DeleteAdmin(uint(id)); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(writer).Encode(map[string]string{
		"message": "Admin deleted successfully",
	})
}

func (handler *AdminHandler) GetAllAdmins(writer http.ResponseWriter, request *http.Request) {
	admins, err := handler.Service.GetAllAdmins()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(writer).Encode(admins)
}
