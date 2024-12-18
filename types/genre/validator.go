package types

type Genre struct {
	Id          string `json:"-" db:"id"`
	Code        string `json:"code,omitempty" db:"code"`
	Name        string `json:"name,omitempty" db:"name" search:"allow"`
	Slug        string `json:"slug,omitempty" db:"slug" search:"allow"`
	Description string `json:"description,omitempty" db:"description"`
}

type GenreStore interface {
	CreateGenre(Genre) (*Genre, error)
	FindGenreByCode(string) (*Genre, error)
	UpdateGenre(string, Genre) (*Genre, error)
	DeleteGenres([]string) error
	SearchGenres(GenreSearch, int, int) ([]Genre, error)
	CountGenres(GenreSearch) (int, error)
}

type CreateGenrePayload struct {
	Name        string `json:"name" binding:"required"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
}

type FindGenrePayload struct {
	Code string `uri:"code" binding:"required"`
}

type GenreSearch struct {
	Filters   map[string]string `json:"filters"`
	SortField string            `json:"sortField"`
	SortOrder string            `json:"sortOrder"`
	Limit     int               `json:"limit"`
	Page      int               `json:"page"`
}

type DeleteGenrepayload struct {
	Codes []string `json:"codes"`
}

type GenreSearchResponse struct {
	Total int     `json:"total"`
	Limit int     `json:"limit"`
	Page  int     `json:"page"`
	Data  []Genre `json:"data,omitempty"`
}
