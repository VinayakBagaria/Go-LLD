package splitwise

import "fmt"

type BalanceSheet struct {
	balances map[*User](map[*User]float64)
}

func NewBalanceSheet() *BalanceSheet {
	return &BalanceSheet{
		balances: make(map[*User]map[*User]float64),
	}
}

func (b *BalanceSheet) UpdateBalances(expense *Expense) {
	paidBy := expense.PaidBy

	for _, split := range expense.Splits {
		owedBy := split.GetUser()
		amount := split.GetAmount()

		if owedBy != paidBy {
			b.updateBalanceForUser(paidBy, owedBy, amount)
			b.updateBalanceForUser(owedBy, paidBy, -amount)
		}
	}
}

func (b *BalanceSheet) updateBalanceForUser(paidBy, owedBy *User, amount float64) {
	_, ok := b.balances[paidBy]
	if !ok {
		b.balances[paidBy] = map[*User]float64{}
	}

	b.balances[paidBy][owedBy] += amount
}

func (b *BalanceSheet) SettleBalances(toUser, fromUser *User) {
	owed := b.balances[toUser][fromUser]
	if owed > 0 {
		b.createTransaction(toUser, fromUser, owed)
	} else if owed < 0 {
		b.createTransaction(fromUser, toUser, -owed)
	}

	b.balances[fromUser][toUser] = 0
	b.balances[toUser][fromUser] = 0
}

func (b *BalanceSheet) createTransaction(sender, receiver *User, amount float64) {
	fmt.Printf("Transaction: %s receives from %s an amount of %.2f\n", sender.Name, receiver.Name, amount)
}

func (b *BalanceSheet) PrintBalances() {
	for sender, sheet := range b.balances {
		fmt.Printf("\nUser: %s\n", sender.Name)

		for receiver, amount := range sheet {
			if amount > 0 {
				fmt.Printf("\tBalance with %s : %.2f\n", receiver.Name, amount)
			}
		}
	}
}
