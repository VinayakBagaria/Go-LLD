package splitwise

type SplitWise struct {
	users        map[string]*User
	groups       map[string]*Group
	balanceSheet *BalanceSheet
}

func NewSplitWise() *SplitWise {
	return &SplitWise{
		users:        make(map[string]*User),
		groups:       make(map[string]*Group),
		balanceSheet: NewBalanceSheet(),
	}
}

func (s *SplitWise) AddUser(user *User) {
	s.users[user.ID] = user
}

func (s *SplitWise) AddGroup(group *Group) {
	s.groups[group.ID] = group
}

func (s *SplitWise) AddExpense(groupId string, expense *Expense) {
	group, ok := s.groups[groupId]
	if !ok {
		return
	}

	group.AddExpense(expense)
	expense.CalculateSplit()
	expense.Print()
	s.balanceSheet.UpdateBalances(expense)
}

func (s *SplitWise) SettleBalances(sender, receiver *User) {
	s.balanceSheet.SettleBalances(sender, receiver)
}

func (s *SplitWise) PrintBalances() {
	s.balanceSheet.PrintBalances()
}
