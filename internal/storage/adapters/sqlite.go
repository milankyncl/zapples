package adapters

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/sqlite3"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/mattn/go-sqlite3"
	"github.com/milankyncl/feature-toggles/internal/storage"
)

const (
	featuresTable = "features"
)

var (
	getFeaturesQuery = fmt.Sprintf(
		"SELECT id, key, description, enabled, created_at "+
			"FROM %s",
		featuresTable,
	)
	getFeatureByIdQuery = fmt.Sprintf(
		"SELECT id, key, description, enabled, created_at "+
			"FROM %s WHERE id = ?",
		featuresTable,
	)
	insertFeatureQuery = fmt.Sprintf(
		"INSERT INTO %s (`key`, `description`, `enabled`, `created_at`) "+
			"VALUES (?, ?, true, CURRENT_TIMESTAMP)",
		featuresTable,
	)
	updateFeatureQuery = fmt.Sprintf(
		"UPDATE %s SET key = ?, description = ? WHERE id = ?",
		featuresTable,
	)
	toggleFeatureQuery = fmt.Sprintf(
		"UPDATE %s SET enabled = ? WHERE id = ?",
		featuresTable,
	)
	deleteFeatureQuery = fmt.Sprintf(
		"DELETE FROM %s WHERE id = ?",
		featuresTable,
	)
)

type SQLite struct {
	db *sql.DB
}

func NewSQLite(dbpath string, migrationsDir string) (error, *SQLite) {
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		return err, nil
	}

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return err, nil
	}

	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", migrationsDir), "main", driver)
	if err != nil {
		return err, nil
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err, nil
	}

	return nil, &SQLite{
		db,
	}
}

func (s *SQLite) GetAll() ([]storage.Feature, error) {
	recs := make([]storage.Feature, 0)
	rows, err := s.db.Query(getFeaturesQuery)
	defer rows.Close()
	if err != nil {
		return recs, err
	}
	for rows.Next() {
		f := storage.Feature{}
		err = rows.Scan(&f.Id, &f.Key, &f.Description, &f.Enabled, &f.CreatedAt)
		if err != nil {
			return recs, err
		}
		recs = append(recs, f)
	}
	return recs, nil
}

func (s *SQLite) GetOne(id int) (storage.Feature, error) {
	f := storage.Feature{}
	row := s.db.QueryRow(getFeatureByIdQuery, id)
	err := row.Scan(&f.Id, &f.Key, &f.Description, &f.Enabled, &f.CreatedAt)
	if err == sql.ErrNoRows {
		return f, storage.ErrFeatureNotFound
	}
	if err != nil {
		return f, err
	}
	return f, nil
}

func (s *SQLite) Create(data storage.CreateFeatureData) error {
	stmt, err := s.db.Prepare(insertFeatureQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(data.Key, data.Description)

	return err
}

func (s *SQLite) Update(id int, data storage.UpdateFeatureData) error {
	stmt, err := s.db.Prepare(updateFeatureQuery)
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(data.Key, data.Description, id)

	return err
}

func (s *SQLite) Toggle(id int, enabled bool) error {
	stmt, err := s.db.Prepare(toggleFeatureQuery)
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(enabled, id)
	return err
}

func (s *SQLite) Delete(id int) error {
	stmt, err := s.db.Prepare(deleteFeatureQuery)
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	return err
}

func (s *SQLite) Close() error {
	return s.db.Close()
}
