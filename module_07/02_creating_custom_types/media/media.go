package media

import utils "shaphil.me/go-standard-library"

type Catalogable interface {
	// constructor
	NewMovie(title string, rating Rating, boxOffice float32)
	// getters
	GetTitle() string
	GetRating() string
	GetBoxOffice() float32

	// setters
	SetTitle(title string)
	SetRating(rating string)
	SetBoxOffice(boxOffice float32)
}

type Movie struct {
	title     string
	rating    Rating
	boxOffice float32
}

type Rating string

const (
	R    = "R (Restricted)"
	G    = "G (General Audiences)"
	PG   = "PG (Parental Guidance)"
	PG13 = "PG-13 (Parental Caution)"
	NC17 = "NC-17 (No children under 17)"
)

func (m *Movie) NewMovie(title string, rating Rating, boxOffice float32) {
	m.title = title
	m.rating = rating
	m.boxOffice = boxOffice
}

func (m *Movie) GetTitle() string {
	return utils.ToTitle(m.title)
}

func (m *Movie) GetRating() Rating {
	return m.rating
}

func (m *Movie) GetBoxOffice() float32 {
	return m.boxOffice
}

func (m *Movie) SetTitle(title string) {
	m.title = title
}

func (m *Movie) SetRating(rating Rating) {
	m.rating = rating
}

func (m *Movie) SetBoxOffice(boxOffice float32) {
	m.boxOffice = boxOffice
}
