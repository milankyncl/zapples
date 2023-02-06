package adapters

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/sqlite3"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/mattn/go-sqlite3"
	"github.com/milankyncl/feature-toggles/internal/storage"
	"path/filepath"
)

const (
	featuresTable = "features"
)

var (
	getFeaturesQuery = fmt.Sprintf(
		"SELECT id, key, description, enabled, enabled_since, enabled_until, created_at "+
			"FROM %s",
		featuresTable,
	)
	getFeatureByIdQuery = fmt.Sprintf(
		"SELECT id, key, description, enabled, enabled_since, enabled_until, created_at "+
			"FROM %s WHERE id = ?",
		featuresTable,
	)
	insertFeatureQuery = fmt.Sprintf(
		"INSERT INTO %s (`key`, `description`, `enabled`, `enabled_since`, `enabled_until`, `created_at`) "+
			"VALUES (?, ?, true, ?, ?, CURRENT_TIMESTAMP)",
		featuresTable,
	)
	updateFeatureQuery = fmt.Sprintf(
		"UPDATE %s SET key = ?, description = ?, enabled_since = ?, enabled_until = ? WHERE id = ?",
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

func NewSQLite(basePath string) (error, *SQLite) {
	db, err := sql.Open("sqlite3", filepath.Join(basePath, "db.sqlite"))
	if err != nil {
		return err, nil
	}

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return err, nil
	}

	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", filepath.Join(basePath, "migrations/")), "main", driver)
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
	if err != nil {
		return recs, err
	}
	defer rows.Close()
	for rows.Next() {
		f := storage.Feature{}
		err = rows.Scan(&f.Id, &f.Key, &f.Description, &f.Enabled, &f.EnabledSince, &f.EnabledUntil, &f.CreatedAt)
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
	err := row.Scan(&f.Id, &f.Key, &f.Description, &f.Enabled, &f.EnabledUntil, &f.EnabledSince, &f.CreatedAt)
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

	_, err = stmt.Exec(data.Key, data.Description, data.EnabledSince, data.EnabledUntil)
	return err
}

func (s *SQLite) Update(id int, data storage.UpdateFeatureData) error {
	stmt, err := s.db.Prepare(updateFeatureQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(data.Key, data.Description, data.EnabledUntil, data.EnabledSince, id)
	return err
}

func (s *SQLite) Toggle(id int, enabled bool) error {
	stmt, err := s.db.Prepare(toggleFeatureQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(enabled, id)
	return err
}

func (s *SQLite) Delete(id int) error {
	stmt, err := s.db.Prepare(deleteFeatureQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	return err
}

func (s *SQLite) Close() error {
	return s.db.Close()
}
