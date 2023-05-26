package models

type (
	Products struct {
		ID       uint64 `json:"product_id"`
		Code     string `form:"code"`
		Title    string `json:"title"`
		Price    int    `json:"price"`
		Weight   int    `json:"weight"`
		Stock    int    `json:"stock"`
		Address  string `json:"address"`
		Category string `json:"category"`
		Timestamp
	}
)
