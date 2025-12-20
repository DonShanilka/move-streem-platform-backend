package Handler

import (
	"encoding/json"
	"io"
	"strconv"

	"net/http"

	"github.com/DonShanilka/user-service/internal/Models"
	"github.com/DonShanilka/user-service/internal/Service"
)

type UserHandler struct {
	Service *Service.UserService
}

func NewUserHandler(service *Service.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func atoiSafe(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func (handler *UserHandler) CreateUser(writer http.ResponseWriter, request *http.Request) {

	if err := request.ParseMultipartForm(100 << 20); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	user := Models.User{
		Name:     request.FormValue("name"),
		Email:    request.FormValue("email"),
		Password: request.FormValue("password"),
		IsActive: request.FormValue("isActive") == "true",
	}

	if file, _, err := request.FormFile("profile_image"); err == nil {
		user.ProfileImage, _ = io.ReadAll(file)
		file.Close()
	}

	if err := handler.Service.CreateUser(&user); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(map[string]string{
		"message": "User created successfully",
	})
}

func (handler *UserHandler) UpdateUser(writer http.ResponseWriter, request *http.Request) {

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

	user := Models.User{
		ID:       uint(id),
		Name:     request.FormValue("name"),
		Email:    request.FormValue("email"),
		Password: request.FormValue("password"),
		IsActive: request.FormValue("isActive") == "true",
	}

	if file, _, err := request.FormFile("profile_image"); err == nil {
		user.ProfileImage, _ = io.ReadAll(file)
		file.Close()
	}

	if err := handler.Service.UpdateUser(uint(id), &user); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(map[string]string{
		"message": "User updated successfully",
	})

}

func (handler *UserHandler) DeleteUser(writer http.ResponseWriter, request *http.Request) {
	id, _ := strconv.Atoi(request.URL.Query().Get("id"))

	if err := handler.Service.DeleteUser(uint(id)); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(writer).Encode(map[string]string{
		"message": "User deleted successfully",
	})
}

func (handler *UserHandler) GetAllUsers(writer http.ResponseWriter, request *http.Request) {
	admins, err := handler.Service.GetAllUsers()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(writer).Encode(admins)
}
