package events

type InvoiceCreated struct {
	ID         uint      `json:"id"`
	Products   []Product `json:"products"`
	Quantities []int     `json:"quantities"`
}

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
}
