package vendingmachine

type Inventory struct {
	products map[*Product]int
}

func NewInventory() *Inventory {
	return &Inventory{products: make(map[*Product]int)}
}

func (i *Inventory) AddProduct(product *Product, quantity int) {
	i.products[product] = quantity
}

func (i *Inventory) IsAvailable(product *Product) bool {
	qty, exists := i.products[product]
	return exists && qty > 0
}

func (i *Inventory) ReduceStock(product *Product) {
	i.products[product]--
}
