package repository

import (
	"context"
	"fmt"
	"github.com/fichca/junior-astrologer-service/internal/model"
	"github.com/jmoiron/sqlx"
	"time"
)

type apodRepo struct {
	db *sqlx.DB
}

func NewAPODRepo(db *sqlx.DB) *apodRepo {
	return &apodRepo{
		db: db,
	}
}

func (a *apodRepo) GetByDate(ctx context.Context, date *time.Time) (*model.APOD, error) {
	query := `
        SELECT id, title, explanation, copyright, date 
        FROM apod 
        WHERE date = $1
    `
	var apod model.APOD
	err := a.db.GetContext(ctx, &apod, query, date)
	if err != nil {
		return nil, fmt.Errorf("failed to get APOD by date: %w", err)
	}

	return &apod, nil
}

func (a *apodRepo) GetAll(ctx context.Context) ([]*model.APOD, error) {
	query := `
        SELECT id, title, explanation, copyright, date 
        FROM apod
    `
	var apods []*model.APOD
	err := a.db.SelectContext(ctx, &apods, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get APODs: %w", err)
	}

	return apods, nil
}

func (a *apodRepo) Save(ctx context.Context, apod *model.APOD) error {
	query := `
        INSERT INTO apod (id, title, explanation, copyright, date) VALUES (:id, :title, :explanation, :copyright, :date)`
	_, err := a.db.NamedExecContext(ctx, query, apod)
	if err != nil {
		return fmt.Errorf("failed to save APOD: %w", err)
	}
	return nil
}
