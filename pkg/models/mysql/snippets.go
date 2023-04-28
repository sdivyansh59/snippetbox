package mysql

import (
	"database/sql"

	"github.com/sdivyansh59/snippetbox/pkg/models"
	
)

// Define a SnippetModel type which wraps a sql.DB connection pool.
type SnippetModel struct {
	DB *sql.DB
}

 
// This will insert a new snippet into the database.
func (m *SnippetModel) Insert (title, content , expires string) ( int, error) {
	stmt := `INSERT INTO snippets (title, created, expires) 
	VALUES (?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	//ge newely inserted obj id
	id ,err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	
	return int(id), nil
}

// This will return a specific snippet based on it's id.
func (m *SnippetModel) Get (id int) ( []*models.Snippet, error) {
	return nil, nil
}

// This will return the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}