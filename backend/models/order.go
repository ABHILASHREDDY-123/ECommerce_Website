package models

type Obj struct {
	ProductID   uint   `json:"product_id"`
	Quantity    uint64 `json:"quantity"`
	CostPerUnit uint64 `json:"cost_per_unit"`
}
type Order struct {
	ID     uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Items  []byte `json:"items"`
	UserID uint   `json:"user_id"`
	Address  string `json:"address"`
}