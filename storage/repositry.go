package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepo struct {
	DbPool *pgxpool.Pool
}

func NewPostgresRepo(pool *pgxpool.Pool) *PostgresRepo {
	return &PostgresRepo{
		DbPool: pool,
	}
}

func PrintUrl(url URL) {
	fmt.Printf("ID: %d \nOrginalURL: %s \nShortCode: %s \nCreatedAt: %s\n",
		url.ID, url.OrigUrl, url.ShortCode, url.CreatedTime)
}

func (repo *PostgresRepo) AddOrigUrl(ctx context.Context, url string) (uint64, error) {

	sql := `INSERT INTO urls (original_url, short_code) 
			VALUES ($1, '') 
			ON CONFLICT (original_url)
			DO UPDATE SET original_url = EXCLUDED.original_url
			RETURNING id`

	var id uint64

	err := repo.DbPool.QueryRow(ctx, sql, url).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (repo *PostgresRepo) UpdateShortCode(ctx context.Context, id uint64, shortCode string) error {

	sql := `UPDATE urls SET short_code=$1 WHERE id=$2`
	_, err := repo.DbPool.Exec(ctx, sql, shortCode, id)

	if err != nil {
		return err
	}

	return nil
}

func (repo *PostgresRepo) GetById(ctx context.Context, id int) (URL, error) {

	sql := `SELECT id, original_url, short_code, created_at 
			FROM urls 
			WHERE id=$1`

	var url URL

	err := repo.DbPool.QueryRow(ctx, sql, id).Scan(
		&url.ID,
		&url.OrigUrl,
		&url.ShortCode,
		&url.CreatedTime,
	)

	if err != nil {
		return url, err
	}

	return url, nil
}

func (repo *PostgresRepo) GetByShortCode(ctx context.Context, shortCode string) (URL, error) {

	sql := `SELECT id, original_url, short_code, created_at
			FROM urls
			WHERE short_code=$1`

	var url URL

	err := repo.DbPool.QueryRow(ctx, sql, shortCode).Scan(
		&url.ID,
		&url.OrigUrl,
		&url.ShortCode,
		&url.CreatedTime,
	)

	if err != nil {
		return url, err
	}

	return url, nil
}

func (repo *PostgresRepo) RemoveById(ctx context.Context, id int) error {

	sql := `DELETE FROM urls WHERE id=$1`

	_, err := repo.DbPool.Exec(ctx, sql, id)

	if err != nil {
		return err
	}

	return nil
}
