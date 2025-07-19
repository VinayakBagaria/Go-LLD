// https://github.com/ashishps1/awesome-low-level-design/blob/main/problems/vending-machine.md
package vendingmachine

import "fmt"

func DoWork() {
	vm := NewVendingMachine()

	// Create some products
	coke := NewProduct("Coke", 0.95)
	bread := NewProduct("Bread", 0.5)
	water := NewProduct("Water", 1.0)

	// Add products to inventory
	vm.inventory.AddProduct(coke, 5)
	vm.inventory.AddProduct(bread, 3)
	vm.inventory.AddProduct(water, 2)

	fmt.Println("Selecting coke")
	vm.SelectProduct(coke)

	fmt.Println("Inserting coins")
	vm.InsertCoin(QUARTER)
	vm.InsertCoin(QUARTER)
	vm.InsertCoin(QUARTER)
	vm.InsertCoin(QUARTER)

	fmt.Println("Dispensing product")
	vm.DispenseProduct()

	fmt.Println("Returning change")
	vm.ReturnChange()

	fmt.Println("\n\nWith insufficient funds")
	vm.SelectProduct(bread)
	vm.InsertCoin(QUARTER)
	fmt.Println("Trying to dispense bread")
	vm.DispenseProduct()

	fmt.Println("\nAdding more coins")
	vm.InsertCoin(QUARTER)
	fmt.Println("Adding extra money - shouldn't be accepted")
	vm.InsertCoin(QUARTER)

	fmt.Println("Dispensing product")
	vm.DispenseProduct()

	fmt.Println("Returning change")
	vm.ReturnChange()
}
