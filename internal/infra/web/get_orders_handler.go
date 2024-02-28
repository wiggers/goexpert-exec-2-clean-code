package web

import (
	"encoding/json"
	"net/http"

	"github.com/devfullcycle/20-CleanArch/internal/entity"
	"github.com/devfullcycle/20-CleanArch/internal/usecase"
	"github.com/devfullcycle/20-CleanArch/pkg/events"
)

type GetWebOrdersHandler struct {
	OrderRepository entity.OrderRepositoryInterface
	GetOrders       events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewGetWebOrdersHandler(
	OrderRepository entity.OrderRepositoryInterface,
	GetOrders events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *GetWebOrdersHandler {
	return &GetWebOrdersHandler{
		OrderRepository: OrderRepository,
		GetOrders:       GetOrders,
		EventDispatcher: EventDispatcher,
	}
}

func (h *GetWebOrdersHandler) Get(w http.ResponseWriter, r *http.Request) {

	getOrders := usecase.NewGetOrdersUseCase(h.OrderRepository, h.GetOrders, h.EventDispatcher)
	output, err := getOrders.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
