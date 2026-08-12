package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"purchase-service/models"
	"purchase-service/service"

	"github.com/gorilla/mux"
)

type OrderItemController struct {
	Service *service.OrderItemService
}


func (oc *OrderItemController) Update(w http.ResponseWriter, r *http.Request) {
	var item models.OrderItem
	json.NewDecoder(r.Body).Decode(&item)

	ret_item, err := oc.Service.Update(&item)

	if err != nil {
		fmt.Println("Error updating item")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ret_item)
}

func (oc *OrderItemController) Save(w http.ResponseWriter, r *http.Request) {
	var item models.OrderItem
	json.NewDecoder(r.Body).Decode(&item)
	
	_, err := oc.Service.Save(&item)

	if err != nil {
		fmt.Println("Error saving item")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func (oc *OrderItemController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		fmt.Println("Error getting id from vars")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	item, err := oc.Service.Delete(id)

	if err != nil {
		fmt.Println("Error deleting items by id")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func (oc *OrderItemController) GetAll(w http.ResponseWriter, r *http.Request) {
	items, err := oc.Service.GetAll()

	if err != nil {
		fmt.Println("Error getting items by id")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func (oc *OrderItemController) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		fmt.Println("Error getting id from vars")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	item, err := oc.Service.GetById(id)


	if err != nil {
		fmt.Println("Error getting items by id")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func (oc *OrderItemController) GetByShoppingCartId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		fmt.Println("Error getting id from vars")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	items, err := oc.Service.GetByShoppingCartId(id)

	if err != nil {
		fmt.Println("Error getting items for cart")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
