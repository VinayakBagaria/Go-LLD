package splitwise

import "fmt"

type Expense struct {
	ID          string
	Amount      float64
	Description string
	PaidBy      *User
	Splits      []Split
}

func NewExpense(id string, amount float64, description string, paidBy *User) *Expense {
	return &Expense{
		ID:          id,
		Amount:      amount,
		Description: description,
		PaidBy:      paidBy,
		Splits:      []Split{},
	}
}

func (e *Expense) AddSplit(split Split) {
	e.Splits = append(e.Splits, split)
}

func (e *Expense) CalculateSplit() {
	totalAmount := e.Amount
	equalSplitsCount := 0

	remaining := totalAmount

	for _, split := range e.Splits {
		switch v := split.(type) {
		case *EqualSplit:
			equalSplitsCount++
		case *PercentSplit:
			currentAmount := totalAmount * v.Percent / 100.0
			v.SetAmount(currentAmount)
			remaining -= currentAmount
		}
	}

	for _, split := range e.Splits {
		switch v := split.(type) {
		case *EqualSplit:
			v.SetAmount(remaining / float64(equalSplitsCount))
		}
	}
}

func (e *Expense) Print() {
	fmt.Printf("Expense added of %.2f by %s\n", e.Amount, e.PaidBy.Name)
}
