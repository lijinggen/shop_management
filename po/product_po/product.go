package product_po

type Product struct {
	ID               string  `json:"id,omitempty"`
	ImageURL         string  `json:"image_url,omitempty"`
	StorageCode      string  `json:"storage_code,omitempty"`
	StoragePos       string  `json:"storage_pos,omitempty"`
	Name             string  `json:"name,omitempty"`
	Color            string  `json:"color,omitempty"`
	BasePrice        float64 `json:"base_price,omitempty"`
	CostPrice        float64 `json:"cost_price,omitempty"`
	PurchasePrice    float64 `json:"purchase_price,omitempty"`
	Factory          string  `json:"factory,omitempty"`
	Stock            int     `json:"stock,omitempty"`
	InProductionNums int     `json:"in_production_nums,omitempty"`
	InOrderNums      int     `json:"in_order_nums,omitempty"`
}
