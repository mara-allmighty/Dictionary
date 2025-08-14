package reports

import (
	"database/sql"
)

type ReportsRepo struct {
	db *sql.DB
}

func NewReportsRepo(db *sql.DB) *ReportsRepo {
	return &ReportsRepo{db: db}
}

// получаем репорт по id
func (rr *ReportsRepo) GetReport() (*Report, error) {
	return nil, nil
}

// создаем репорт
func (rr *ReportsRepo) CreateReport() error {
	return nil
}

// редактируем репорт
func (rr *ReportsRepo) UpdateReport() error {
	return nil
}

// удаляем репорт
func (rr *ReportsRepo) DeleteReport() error {
	return nil
}
