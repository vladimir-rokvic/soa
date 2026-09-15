package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"purchase-service/dto"
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

	var ret []dto.TourDTO
	for _, t := range tokens {
		resp, err := http.Get("http://tours-service:8080/tours/" + t.TourId)
		if err != nil {
			fmt.Println("Error getting tour")
			fmt.Println(err)
		}

		var n dto.TourDTO
		n.TokenId = t.ID.String()
		json.NewDecoder(resp.Body).Decode(&n)
		resp.Body.Close()

		ret = append(ret, n)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ret)
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
		json.NewEncoder(w).Encode(nil)
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

	var currentPosition dto.CurrentPosDTO
	err := json.NewDecoder(r.Body).Decode(&currentPosition)

	if err != nil {
		fmt.Println("Error decoding current position body")
		fmt.Println(err)
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

	te, err := tc.Service.CreateTourExecution(currentPosition, token)

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

func (tc *TokenController) UpdateTour(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		fmt.Println("Error getting vars")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var currentPosition dto.CurrentPosDTO
	err := json.NewDecoder(r.Body).Decode(&currentPosition)

	if err != nil {
		fmt.Println("Error decoding current position body")
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	te, err := tc.Service.GetTEById(id)

	if err != nil {
		fmt.Println("Error getting tour execution")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	te, err, completed := tc.Service.UpdateTour(&te, currentPosition)
	
	if err != nil {
		fmt.Println("Error updating tour")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(completed)
}

func (tc *TokenController) GetTourFromToken(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		fmt.Println("Error getting vars")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := tc.Service.GetById(id)
	if err != nil {
		fmt.Println("Error getting token by id")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := http.Get("http://tours-service:8080/tours/" + token.TourId)
	if err != nil {
		fmt.Println("Error getting tour from tours service")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var n dto.TourDTO
	n.TokenId = token.ID.String()
	json.NewDecoder(res.Body).Decode(&n)
	res.Body.Close()

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(n)
}

func (tc *TokenController) AbandonTour(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		fmt.Println("Error getting vars")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	te, err := tc.Service.GetTEById(id)
	if err != nil {
		fmt.Println("Error getting tour execution")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ret, err := tc.Service.AbandonTE(te)
	if err != nil {
		fmt.Println("Error updating te")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ret)
}

func (tc *TokenController) UpdateActivity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		fmt.Println("Error getting vars")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	te, err := tc.Service.GetTEById(id)
	if err != nil {
		fmt.Println("Error getting tour execution")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ret, err := tc.Service.UpdateActivity(te)
	if err != nil {
		fmt.Println("Error updating te")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ret)
}
