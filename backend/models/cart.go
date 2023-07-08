package models

type CartObj struct {
	ProductId uint  `json:"product_id"`
	Quantity  int64 `json:"quantity"`
}
type Cart struct {
	Objs []CartObj  `json:"cart"`
}