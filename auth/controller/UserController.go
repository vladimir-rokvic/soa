package controller

import (
	"auth_service/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type UserController struct {
	UserService *service.UserService
}

func (user_controller *UserController) GetAll(writer http.ResponseWriter, 
req *http.Request){
	users := user_controller.UserService.GetAll()

	writer.Header().Set("Content-Type", "application/json")
	if users == nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(users)
}

func (user_controller *UserController) GetById(writer http.ResponseWriter, 
req *http.Request) {
	id := mux.Vars(req)["id"]
	user := user_controller.UserService.GetById(id)

	writer.Header().Set("Content-Type", "application/json")
	if user == nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(user)
}

func (user_controller *UserController) Create() {
}
