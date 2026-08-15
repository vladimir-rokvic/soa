package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"purchase-service/models"
	"purchase-service/service"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type TokenController struct {
	Service *service.TourTokenService
}

func (tc *TokenController) GetByUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		fmt.Println("Error getting vars")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokens, err := tc.Service.GetByUserId(id)

	if err != nil {
		fmt.Println("Error getting by user id")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tokens)
}

func (tc *TokenController) GetActiveByUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		fmt.Println("Error getting vars")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	te, err := tc.Service.GetActiveByUserId(id)

	if err == gorm.ErrRecordNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err != nil {
		fmt.Println("Error getting an active tour")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(te)
}

func (tc *TokenController) StartTour(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		fmt.Println("Error getting vars")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := tc.Service.GetById(id)

	if err == gorm.ErrRecordNotFound {
		fmt.Println("Error token not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	
	if err != nil {
		fmt.Println("Error getting token by id")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	token.Status = models.ACTIVE
	tc.Service.Update(&token)

	te, err := tc.Service.CreateTourExecution(token)

	if err != nil {
		fmt.Println("Error creating tour execution")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(te)
}
