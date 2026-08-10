package controller

import (
	"net/http"
	"purchase-service/service"
)

type OrderItemController struct {
	Service *service.OrderItemService
}


func (oc *OrderItemController) Update(w http.ResponseWriter, r *http.Request) {
}
func (oc *OrderItemController) Save(w http.ResponseWriter, r *http.Request) {
}
func (oc *OrderItemController) Delete(w http.ResponseWriter, r *http.Request) {
}
func (oc *OrderItemController) GetAll(w http.ResponseWriter, r *http.Request) {
}
func (oc *OrderItemController) GetById(w http.ResponseWriter, r *http.Request) {
}
func (oc *OrderItemController) GetByShoppingCartId(w http.ResponseWriter, r *http.Request) {
}
