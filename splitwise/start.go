// https://github.com/ashishps1/awesome-low-level-design/blob/main/problems/splitwise.md
package splitwise

import "fmt"

func DoWork() {
	service := NewSplitWise()

	// Create users
	user1 := NewUser("1", "Vinayak")
	user2 := NewUser("2", "Shirley")
	user3 := NewUser("3", "Mummy")
	service.AddUser(user1)
	service.AddUser(user2)
	service.AddUser(user3)

	// Create group and add users
	group := NewGroup("1", "Apartment")
	group.AddMember(user1)
	group.AddMember(user2)
	group.AddMember(user3)
	service.AddGroup(group)

	// Add expense1
	expense1 := NewExpense("1", 100, "Rent", user1)
	expense1.AddSplit(NewEqualSplit(user1))
	expense1.AddSplit(NewEqualSplit(user2))
	expense1.AddSplit(NewPercentSplit(user3, 20))
	service.AddExpense(group.ID, expense1)

	// Add expense1
	// expense2 := NewExpense("1", 40, "Utility", user2)
	// expense2.AddSplit(NewEqualSplit(user1))
	// expense2.AddSplit(NewEqualSplit(user2))
	// service.AddExpense(group.ID, expense2)

	service.PrintBalances()

	// Print balances
	// for _, user := range []*User{user1, user2, user3} {
	// 	user.PrintBalances()
	// }

	// Settle balances
	fmt.Println()
	service.SettleBalances(user1, user2)
	service.SettleBalances(user1, user3)
	service.SettleBalances(user2, user3)

	service.PrintBalances()
}
