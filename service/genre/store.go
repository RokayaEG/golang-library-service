package genre

import (
	"database/sql"
	"errors"
	"strings"

	types "github.com/RokayaEG/golang-library-service/types/genre"
	"github.com/jmoiron/sqlx"
)

type Store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) CreateGenre(genre types.Genre) (*types.Genre, error) {
	response, err := s.db.NamedExec("INSERT INTO genres (code, name, slug, description) VALUES (:code, :name, :slug, :description)", &genre)
	if err != nil {
		return nil, err
	}

	// Get the last inserted id
	lastId, err := response.LastInsertId()
	if err != nil {
		return nil, err
	}

	var _genre types.Genre

	err = s.db.Get(&_genre, "SELECT id, code, name, slug, description FROM genres where id = ?", lastId)
	if err != nil {
		return nil, err
	}
	return &_genre, nil
}

func (s *Store) FindGenreByCode(code string) (*types.Genre, error) {
	var _genre types.Genre

	err := s.db.Get(&_genre, "SELECT * FROM genres where code = ?", code)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, errors.New("genre not found")
		}
		return nil, err
	}

	return &_genre, nil

}

func (s *Store) UpdateGenre(code string, updatedGenre types.Genre) (*types.Genre, error) {

	_genre, err := s.FindGenreByCode(code)
	if err != nil {
		return nil, err
	}

	_, err = s.db.Query("UPDATE genres SET name = ?, slug = ?, description = ? WHERE id = ?", updatedGenre.Name, updatedGenre.Slug, updatedGenre.Description, _genre.Id)

	if err != nil {
		return nil, err
	}

	return s.FindGenreByCode(code)
}

func (s *Store) SearchGenres(srch types.GenreSearch) ([]types.Genre, error) {

	var searchClause string
	if srch.Filters != nil {
		for key, value := range srch.Filters {
			if len(value) > 0 {
				switch key {
				case "search":
					searchClause += "g.code LIKE '%" + value + "%' OR g.name LIKE '%" + value + "%' OR g.slug LIKE '%" + value + "%'"
				default:
				}
			}
		}
	}
	var selectClause string
	if len(searchClause) > 0 {
		selectClause += "\nWHERE " + searchClause
	}

	var orderClause string
	if len(srch.SortField) > 0 {
		orderClause += "\nORDER BY " + srch.SortField + " "
		if len(srch.SortOrder) > 0 && (strings.ToUpper(srch.SortOrder) == "ASC" || strings.ToUpper(srch.SortOrder) == "DESC") {
			orderClause += strings.ToUpper(srch.SortOrder)
		} else {
			orderClause += "ASC"
		}
	}

	var genres []types.Genre
	err := s.db.Select(&genres, "SELECT * FROM genres AS g"+selectClause+orderClause+";")

	if err != nil {
		return nil, err
	}

	return genres, nil
}

func (s *Store) DeleteGenres(codes []string) error {
	query, args, err := sqlx.In("DELETE FROM genres WHERE code IN (?);", codes)

	if err != nil {
		return err
	}

	query = s.db.Rebind(query)
	_, err = s.db.Query(query, args...)

	if err != nil {
		return err
	}
	return nil
}
