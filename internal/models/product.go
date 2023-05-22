package models

type (
	Products struct {
		ID       uint64 `json:"product_id"`
		Code     string `form:"code"`
		Title    string `json:"title"`
		Price    string `json:"price"`
		Weight   string `json:"weight"`
		Stock    string `json:"stock"`
		Address  string `json:"address"`
		Category string `json:"category"`
		Timestamp
	}
)
