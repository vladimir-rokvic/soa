package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"purchase-service/service"

	"github.com/gorilla/mux"
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
