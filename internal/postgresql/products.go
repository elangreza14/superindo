package postgresql

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/elangreza14/superindo/internal/domain"
	"github.com/elangreza14/superindo/internal/params"
)

type ProductRepo struct {
	db *sql.DB
}

func NewProductsRepo(db *sql.DB) *ProductRepo {
	return &ProductRepo{db}
}

func (pr *ProductRepo) ListProduct(ctx context.Context, req params.ProductQueryParams) ([]domain.Product, error) {
	q := squirrel.
		Select("id", "name", "quantity", "price", "product_type_name", "created_at", "updated_at").
		From("products p")

	if len(req.Search) != 0 {
		q = q.Where(squirrel.Or{
			squirrel.Eq{"p.id": req.Search},
			squirrel.Eq{"p.name": "%" + req.Search + "%"},
		})
	}

	if len(req.Types) != 0 {
		q = q.Where(squirrel.Eq{"p.product_type_name": req.Types})
	}

	for _, sort := range req.Sort {
		q = q.OrderBy(sort)
	}

	qr, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := pr.db.QueryContext(ctx, qr, args)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []domain.Product{}
	for rows.Next() {
		product := domain.Product{}
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Quantity,
			&product.Price,
			&product.ProductType.Name,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
