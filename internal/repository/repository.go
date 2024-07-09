package repository

import (
	"database/sql"
	"go_learning/internal/models"
)

type DatabaseRepository struct {
	db *sql.DB
}


func NewDatabaseRepository(db *sql.DB) *DatabaseRepository {
	return &DatabaseRepository{db: db}
}

func (r *DatabaseRepository) GetAll() ([]*models.Testing, error) {
	query := `
	SELECT id, name, value
	FROM users
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	testingList := []*models.Testing{}

	for rows.Next() {
		var testing models.Testing
		err := rows.Scan(&testing.ID, &testing.Email, &testing.Username)
		if err != nil {
			return nil, err
		}
		testingList = append(testingList, &testing)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return testingList, nil
}