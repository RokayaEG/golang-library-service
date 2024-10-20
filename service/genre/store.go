package genre

import (
	"database/sql"
	"errors"

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
			return nil, errors.New("Genre not found")
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
