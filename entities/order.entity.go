package entities

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Order struct {
	ID       *string  `json:"id"`
	Action   string   `json:"action"`
	Price    float32  `json:"price"`
	Limit    float32  `json:"limit"`
	Stoploss float32  `json:"stoploss"`
	Time     string   `json:"time"`
	Lot      *float32 `json:"lot"`
	Symbol   string   `json:"symbol"`
}

func NewOrder(symbol string) *Order {
	id := uuid.Must(uuid.NewV7()).String()
	return &Order{
		ID:     &id,
		Symbol: symbol,
	}
}

func (sp *Order) SetLot(lot float32) {
	sp.Lot = &lot
}

func (sp *Order) SetPrice(price float32) {
	sp.Price = price
}

func (sp *Order) SetTakeProfit(limit float32) {
	sp.Limit = limit
}

func (sp *Order) SetStopLoss(stoploss float32) {
	sp.Stoploss = stoploss
}

func (sp *Order) SetBuy() {
	sp.Action = "BUY"
}

func (sp *Order) SetSell() {
	sp.Action = "SELL"
}

func (sp *Order) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, sp)
}

func (sp Order) MarshalBinary() ([]byte, error) {
	return json.Marshal(sp)
}
