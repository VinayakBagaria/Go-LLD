package vendingmachine

type VendingMachine struct {
	inventory       Inventory
	selectedProduct *Product
	payment         float64

	idleState         VendingMachineState
	readyState        VendingMachineState
	dispenseState     VendingMachineState
	returnChangeState VendingMachineState

	currentState VendingMachineState
}

func NewVendingMachine() *VendingMachine {
	vm := &VendingMachine{
		inventory: *NewInventory(),
		payment:   0,
	}
	vm.selectedProduct = nil

	vm.idleState = &IdleState{vm}
	vm.readyState = &ReadyState{vm}
	vm.dispenseState = &DispenseState{vm}
	vm.returnChangeState = &ReturnChangeState{vm}

	vm.currentState = vm.idleState
	return vm
}

func (vm *VendingMachine) SetState(state VendingMachineState) {
	vm.currentState = state
}

func (vm *VendingMachine) ResetPayment() {
	vm.payment = 0
}

func (vm *VendingMachine) ResetSelectedProduct() {
	vm.selectedProduct = nil
}

func (vm *VendingMachine) SelectProduct(product *Product) {
	vm.currentState.SelectProduct(product)
}

func (vm *VendingMachine) InsertCoin(coin Coin) {
	vm.currentState.InsertCoin(coin)
}

func (vm *VendingMachine) DispenseProduct() {
	vm.currentState.DispenseProduct()
}

func (vm *VendingMachine) ReturnChange() {
	vm.currentState.ReturnChange()
}
