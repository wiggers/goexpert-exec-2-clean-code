package event

import "time"

type GetOrders struct {
	Name    string
	Payload interface{}
}

func NewGetOrders() *GetOrders {
	return &GetOrders{
		Name: "GetOrders",
	}
}

func (e *GetOrders) GetName() string {
	return e.Name
}

func (e *GetOrders) GetPayload() interface{} {
	return e.Payload
}

func (e *GetOrders) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *GetOrders) GetDateTime() time.Time {
	return time.Now()
}
