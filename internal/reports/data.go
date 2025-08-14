package reports

import (
	"database/sql"
	"fmt"
)

type ReportsRepo struct {
	db *sql.DB
}

func NewReportsRepo(db *sql.DB) *ReportsRepo {
	return &ReportsRepo{db: db}
}

// GET report
func (rr *ReportsRepo) GetReport(id int) (*Report, error) {
	var report Report

	err := rr.db.QueryRow(`SELECT id, title, overview, created_at, updated_at FROM reports WHERE id = $1`, id).
		Scan(&report.Id, &report.Title, &report.Overview, &report.Created_at, &report.Updated_at)
	if err != nil {
		return nil, err
	}

	return &report, nil
}

// CREATE report
func (rr *ReportsRepo) CreateReport(title, overview string) error {
	_, err := rr.db.Exec(`INSERT INTO reports (title, overview) VALUES ($1, $2)`, title, overview)
	if err != nil {
		return err
	}

	return nil
}

// UPDATE report
func (rr *ReportsRepo) UpdateReport(id int, newTitle, newDescription string) error {
	var title string

	err := rr.db.QueryRow(
		`UPDATE reports
			SET title = $1, overview = $2, updated_at = NOW()
			WHERE id = $3
			RETURNING title`,
		newTitle, newDescription, id,
	).Scan(&title)

	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("report with ID %d not found", id)
		}
		return fmt.Errorf("database error: %w", err)
	}

	return nil
}

// DELETE report
func (rr *ReportsRepo) DeleteReport(id int) error {
	var title string

	err := rr.db.QueryRow(
		`DELETE FROM reports
		 WHERE id = $1
		 RETURNING title`,
		id,
	).Scan(&title)

	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("report with ID %d not found", id)
		}
		return fmt.Errorf("database error: %w", err)
	}

	return nil
}
