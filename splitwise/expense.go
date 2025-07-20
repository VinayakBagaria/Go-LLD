package splitwise

type Expense struct {
	ID          string
	Amount      float64
	Description string
	PaidBy      *User
}

func NewExpense(id string, amount float64, description string, paidBy *User) *Expense {
	return &Expense{
		ID:          id,
		Amount:      amount,
		Description: description,
		PaidBy:      paidBy,
	}
}
