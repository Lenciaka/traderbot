package model

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type OrderEntity struct {
	ID           string   `db:"id"`
	MtID         int      `db:"mt_id"`
	PortID       string   `db:"portId"`
	StrategyID   int      `db:"strategyId"`
	Action       string   `db:"action"`
	Price        float64  `db:"price"`
	Volume       float64  `db:"volume"`
	TakeProfit   *float64 `db:"takeprofit"`
	StopLoss     *float64 `db:"stoploss"`
	IsOrdered    bool     `db:"isOrdered"`
	IsFinished   bool     `db:"isFinished"`
	MtStatus     *int     `db:"mt_status"`
	Profit       *float64 `db:"profit"`
	ClosePrice   *float64 `db:"close_price"`
	IsProfitable *bool    `db:"is_profitable"`
	UpdatedAt    string   `db:"updated_at"`
	CreatedAt    string   `db:"created_at"`
}

type OrderModel struct {
	db *sqlx.DB
}

func NewOrderModel(db *sqlx.DB) *OrderModel {
	return &OrderModel{
		db: db,
	}
}

func (model *OrderModel) GetAll(ctx context.Context) ([]OrderEntity, error) {
	orders := []OrderEntity{}
	err := model.db.SelectContext(ctx, &orders, "SELECT * FROM orders")
	if err == sql.ErrNoRows {
		return orders, nil
	}
	return orders, err
}

func (model *OrderModel) GetLatestOrderByPortIDAndStrategyID(ctx context.Context, portID, strategyID string) (*OrderEntity, error) {
	orders := OrderEntity{}
	err := model.db.GetContext(ctx, &orders, `
		SELECT * FROM orders
		WHERE portId = ? AND strategyId = ? AND isOrdered = 1
		ORDER BY created_at desc
		LIMIT 1
	`, portID, strategyID)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &orders, err
}

func (model *OrderModel) FindByPortIDAndStrategyID(ctx context.Context, portID, strategyID string) ([]OrderEntity, error) {
	orders := []OrderEntity{}
	err := model.db.SelectContext(ctx, &orders, `
		SELECT * FROM orders
		WHERE portId = ? AND strategyId = ? AND isOrdered = 1
		ORDER BY created_at desc
	`, portID, strategyID)

	if err == sql.ErrNoRows {
		return orders, nil
	}

	return orders, err
}

func (model *OrderModel) FindByPortID(ctx context.Context, portID string) ([]OrderEntity, error) {
	orders := []OrderEntity{}
	err := model.db.SelectContext(ctx, &orders, `
		SELECT * FROM orders
		WHERE portId = ? AND isOrdered = 1
		ORDER BY created_at desc
	`, portID)

	if err == sql.ErrNoRows {
		return orders, nil
	}

	return orders, err
}

func (model *OrderModel) FindByStrategyID(ctx context.Context, strategyID string) ([]OrderEntity, error) {
	orders := []OrderEntity{}
	err := model.db.SelectContext(ctx, &orders, `
		SELECT * FROM orders
		WHERE strategyId = ? AND isOrdered = 1
		ORDER BY created_at desc
	`, strategyID)

	if err == sql.ErrNoRows {
		return orders, nil
	}

	return orders, err
}
