package moviebooking

type Theater struct {
	ID       string
	Name     string
	Location string
	Shows    []*Show
}

func NewTheater(id, name, location string) *Theater {
	return &Theater{
		ID:       id,
		Name:     name,
		Location: location,
		Shows:    []*Show{},
	}
}
