package usecase

import (
	"github.com/devfullcycle/20-CleanArch/internal/entity"
	"github.com/devfullcycle/20-CleanArch/pkg/events"
)

type getOrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type GetOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	GetOrders       events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewGetOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	GetOrders events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,

) *GetOrdersUseCase {
	return &GetOrdersUseCase{
		OrderRepository: OrderRepository,
		GetOrders:       GetOrders,
		EventDispatcher: EventDispatcher,
	}
}

func (c *GetOrdersUseCase) Execute() ([]getOrderOutputDTO, error) {

	orders, err := c.OrderRepository.GetAll()
	if err != nil {
		return nil, err
	}

	var results []getOrderOutputDTO
	for _, order := range orders {
		resul := getOrderOutputDTO{ID: order.ID, Tax: order.Tax, Price: order.Price, FinalPrice: order.FinalPrice}
		results = append(results, resul)
	}

	c.GetOrders.SetPayload(results)
	c.EventDispatcher.Dispatch(c.GetOrders)

	return results, nil
}
