package transport

type Order struct {
	OrderID   string `json:"orderId"`
	UserEmail string `json:"userEmail"`
}