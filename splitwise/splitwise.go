package splitwise

type SplitWise struct {
	users  map[string]*User
	groups map[string]*Group
}

func NewSplitWise() *SplitWise {
	return &SplitWise{
		users:  make(map[string]*User),
		groups: make(map[string]*Group),
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
}

func (s *SplitWise) SettleBalances(userID1, userID2 string) {
	_, ok1 := s.users[userID1]
	_, ok2 := s.users[userID2]

	if !ok1 || !ok2 {
		return
	}
}
