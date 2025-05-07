package entities

import (
	"encoding/json"
	"time"
)

type StrategyPrice struct {
	Key       string    `json:"key"`
	Signal    string    `json:"signal"`
	Open      float64   `json:"open"`
	Close     float64   `json:"close"`
	NotifyAt  time.Time `json:"notify_at"`
	StoredAt  time.Time `json:"stored_at"`
	TimeFrame string    `json:"timeframe"`
}

func (sp *StrategyPrice) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, sp)
}
func (sp StrategyPrice) MarshalBinary() ([]byte, error) {
	return json.Marshal(sp)
}
