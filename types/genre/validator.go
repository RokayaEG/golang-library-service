package types

type Genre struct {
	Id          string `json:"-" db:"id"`
	Code        string `json:"code,omitempty" db:"code"`
	Name        string `json:"name,omitempty" db:"name"`
	Slug        string `json:"slug,omitempty" db:"slug"`
	Description string `json:"description,omitempty" db:"description"`
}

type GenreStore interface {
	CreateGenre(Genre) (*Genre, error)
	FindGenreByCode(string) (*Genre, error)
	UpdateGenre(string, Genre) (*Genre, error)
}

type CreateGenrePayload struct {
	Name        string `json:"name" binding:"required"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
}

type FindGenrePayload struct {
	Code string `uri:"code" binding:"required"`
}
