package db

import (
	models "github.com/Daniel-Kasem48/multitenant-sqlc-bolierplate/internal/db/autogenerated"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DatasourceService struct {
	Queries *models.Queries
}

func NewDatasourceServiceService(pool *pgxpool.Pool) *DatasourceService {
	queries := models.New(pool) // Initialize the queries using sqlc-generated code
	return &DatasourceService{Queries: queries}
}