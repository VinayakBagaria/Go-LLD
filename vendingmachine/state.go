package vendingmachine

import "fmt"

type VendingMachineState interface {
	SelectProduct(product *Product)
	InsertCoin(c Coin)
	InsertNote(note Note)
	DispenseProduct()
	ReturnChange()
}

type IdleState struct {
	vm *VendingMachine
}

func (s *IdleState) SelectProduct(product *Product) {
	if s.vm.inventory.IsAvailable(product) {
		s.vm.selectedProduct = product
		s.vm.SetState(s.vm.readyState)
		fmt.Println("Product selected:", product.name)
	} else {
		fmt.Println("Product not available:", product.name)
	}
}

func (s *IdleState) InsertCoin(c Coin)    { fmt.Println("Select a product first") }
func (s *IdleState) InsertNote(note Note) { fmt.Println("Select a product first") }
func (s *IdleState) DispenseProduct()     { fmt.Println("Select a product first and make payment") }
func (s *IdleState) ReturnChange()        { fmt.Println("No change to return") }

type ReadyState struct {
	vm *VendingMachine
}

func (s *ReadyState) SelectProduct(product *Product) {
	fmt.Println("Product already selected. Please make payment.")
}
func (s *ReadyState) InsertCoin(c Coin) {
	s.vm.payment += float64(c)
	s.checkPaymentStatus()
}
func (s *ReadyState) InsertNote(note Note) {
	s.vm.payment += float64(note)
	s.checkPaymentStatus()
}
func (s *ReadyState) DispenseProduct() { fmt.Println("Please make payment first.") }
func (s *ReadyState) ReturnChange()    { fmt.Println("Change returned after payment") }

func (s *ReadyState) checkPaymentStatus() {
	if s.vm.payment >= s.vm.selectedProduct.price {
		s.vm.SetState(s.vm.dispenseState)
	}
}

type DispenseState struct {
	vm *VendingMachine
}

func (s *DispenseState) SelectProduct(product *Product) { fmt.Println("Product already selected.") }
func (s *DispenseState) InsertCoin(c Coin)              { fmt.Println("Please collect the product.") }
func (s *DispenseState) InsertNote(n Note)              { fmt.Println("Please collect the product.") }
func (s *DispenseState) DispenseProduct() {
	s.vm.SetState(s.vm.returnChangeState)
	s.vm.inventory.ReduceStock(s.vm.selectedProduct)
}
func (s *DispenseState) ReturnChange() { fmt.Println("Please collect the product first.") }

type ReturnChangeState struct {
	vm *VendingMachine
}

func (s *ReturnChangeState) SelectProduct(product *Product) {
	fmt.Println("Please collect the change first.")
}
func (s *ReturnChangeState) InsertCoin(c Coin) {
	fmt.Println("Please collect the change first.")
}
func (s *ReturnChangeState) InsertNote(n Note) {
	fmt.Println("Please collect the change first.")
}
func (s *ReturnChangeState) DispenseProduct() {
	fmt.Println("Product already dispensed. Please collect the change.")
}

func (s *ReturnChangeState) ReturnChange() {
	change := s.vm.payment - s.vm.selectedProduct.price
	if change > 0 {
		fmt.Printf("Change returned: %.2f\n", change)
	} else {
		fmt.Println("No change")
	}
	s.vm.ResetPayment()
	s.vm.ResetSelectedProduct()
	s.vm.SetState(s.vm.idleState)
}
