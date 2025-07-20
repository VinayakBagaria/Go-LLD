// https://github.com/ashishps1/awesome-low-level-design/blob/main/problems/splitwise.md
package splitwise

func DoWork() {
	service := NewSplitWise()

	// Create users
	user1 := NewUser("1", "Alice")
	user2 := NewUser("2", "Bob")
	user3 := NewUser("3", "Charlie")
	service.AddUser(user1)
	service.AddUser(user2)
	service.AddUser(user3)

	// Create group and add users
	group := NewGroup("1", "Apartment")
	group.AddMember(user1)
	group.AddMember(user2)
	group.AddMember(user3)
	service.AddGroup(group)

	// Add expense
	expense := NewExpense("1", 300, "Rent", user1)
	service.AddExpense(group.ID, expense)

	// Settle balances
	service.SettleBalances(user1.ID, user2.ID)
	service.SettleBalances(user1.ID, user3.ID)
}
