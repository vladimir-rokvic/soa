package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"purchase-service/dto"
	"purchase-service/models"
	"purchase-service/service"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type ShoppingCartController struct {
	Service *service.ShoppingCartService
	ItemService *service.OrderItemService
	TokenService *service.TourTokenService
}


func (sc *ShoppingCartController) Update(w http.ResponseWriter, r *http.Request) {
	var cart models.ShoppingCart
	json.NewDecoder(r.Body).Decode(&cart)

	ret_cart, err := sc.Service.Update(&cart)

	if err != nil {
		fmt.Println("Error updating cart")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ret_cart)
}

func (sc *ShoppingCartController) Save(w http.ResponseWriter, r *http.Request) {
	var cart models.ShoppingCart
	json.NewDecoder(r.Body).Decode(&cart)

	ret_cart, err := sc.Service.Save(&cart)

	if err != nil {
		fmt.Println("Error saving cart")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ret_cart)
}

func (sc *ShoppingCartController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		fmt.Println("Error getting vars for delete")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cart, err := sc.Service.Delete(id)

	if err != nil {
		fmt.Println("Error deleting cart")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
}

func (sc *ShoppingCartController) GetAll(w http.ResponseWriter, r *http.Request) {
	carts, err := sc.Service.GetAll()

	if err != nil {
		fmt.Println("Error getting all carts")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(carts)
}

func (sc *ShoppingCartController) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		fmt.Println("Error getting vars for get by id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cart, err := sc.Service.GetById(id)

	if err != nil {
		fmt.Println("Error getting all carts")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
}

func (sc *ShoppingCartController) GetByUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		fmt.Println("Error getting vars for get by id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cart, err := sc.Service.GetByUserId(id)

	if err != nil {
		fmt.Println("Error getting cart by user id")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
}

func (sc *ShoppingCartController) AddItem(w http.ResponseWriter, r *http.Request) {
	var item dto.OrderItemDTO
	err := json.NewDecoder(r.Body).Decode(&item)

	if err != nil {
		fmt.Println("Error getting item body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userId, err := uuid.Parse(item.UserId)
	if err != nil {
		fmt.Println("Error parsing user uuid")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cart, err := sc.Service.GetByUserId(item.UserId)

	if err != nil {
		cart = models.ShoppingCart{
			UserId: userId,
		}

		_, err = sc.Service.Save(&cart)

		if err != nil {
			fmt.Println("Error saving cart")
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	//TODO: Proveri da li je ovo potrebno uopste
	cart, err = sc.Service.GetByUserId(item.UserId)
	if err != nil {
		fmt.Println("Error saving cart")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var it models.OrderItem
	it.ShoppingCartID = cart.ID
	it.TourName = item.TourTitle
	it.Price = item.TourPrice
	it.TourId = item.TourId

	_, err = sc.ItemService.Save(&it)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(it)
}

func (sc *ShoppingCartController) BuyItems(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		fmt.Println("Error getting id from vars")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cart, err := sc.Service.GetByUserId(id)

	if err != nil {
		fmt.Println("Error getting items for user")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userUUID, err := uuid.Parse(id)
	if err != nil {
		fmt.Println("Error parsing user id")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, i := range cart.Items {
		token := models.GenerateToken(userUUID, i)
		err := sc.TokenService.Save(&token)
		if err != nil {
			fmt.Println("Error saving token")
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = sc.ItemService.Delete(i.ID.String())
		if err != nil {
			fmt.Println("Error deleting item")
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
