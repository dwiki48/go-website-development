package entity

type Product struct {
	ID    int
	Name  string
	Price int
	Stock int
}

func (p Product) StockStatus() string {
	var status string

	if p.Stock < 5 {
		status = "Stock less than 5"
	} else if p.Stock < 10 {
		status = "Limited stock"
	}

	return status
}
