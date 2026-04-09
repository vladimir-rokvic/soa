package controller

import (
	"auth_service/dto"
	"auth_service/model"
	"auth_service/service"
	"auth_service/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
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

func (user_controller *UserController) Save(writer http.ResponseWriter,
req *http.Request) {
	var user model.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		fmt.Printf("Error encoding user when adding new user ")
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = user_controller.UserService.Save(&user)
	if err != nil {
		fmt.Println("Error saving user")
		fmt.Println(err)
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	writer.WriteHeader(http.StatusCreated)
}

func (user_controller *UserController) LogIn(writer http.ResponseWriter,
req *http.Request){
	var loginDto dto.LoginDTO
	err := json.NewDecoder(req.Body).Decode(&loginDto)
	if err != nil {
		fmt.Println("Error loging in")
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	user := user_controller.UserService.GetByUsername(loginDto.Username)
	if user == nil {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDto.Password))
	if err != nil {
		fmt.Println("Error logging in")
		fmt.Println(err)
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateToken(user.Username)
	if err != nil {
		fmt.Println("Error creating token")
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(map[string]string{"token": token})
}

func (user_controller *UserController) MyProfile(writer http.ResponseWriter,
req *http.Request) {
	claims, ok := utils.GetClaims(req)
	if !ok {
		fmt.Println("Error getting claims from request header")
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	username := claims["username"].(string)
	user := user_controller.UserService.GetByUsername(username)
	if user == nil {
		fmt.Printf("Error getting user by username: %s ", username)
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	userDto := dto.UserToDTO(user)
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(userDto)
}

