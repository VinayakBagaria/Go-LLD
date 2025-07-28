package moviebooking

type Movie struct {
	ID              string
	Title           string
	DurationMinutes int
}

func NewMovie(id, title string, durationMinutes int) *Movie {
	return &Movie{
		ID:              id,
		Title:           title,
		DurationMinutes: durationMinutes,
	}
}
