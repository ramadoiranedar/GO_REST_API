package repository

import (
	"context"
	"database/sql"

	"github.com/ramadoiranedar/go_restapi/helper"
	"github.com/ramadoiranedar/go_restapi/model/domain"
)

type CategoryRepositoryImpl struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &CategoryRepositoryImpl{db: db}
}

func (repository *CategoryRepositoryImpl) Create(ctx context.Context, category domain.Category) domain.Category {
	q := "insert into category (name) values (?)"
	result, err := repository.db.ExecContext(ctx, q, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	category.Id = int(id)

	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, category domain.Category) domain.Category {
	_, err := repository.db.ExecContext(ctx, "UPDATE category SET name = ? WHERE id = ?", category.Name, category.Id)
	helper.PanicIfError(err)
	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, category domain.Category) {
	q := "delete from category where id = ?"
	_, err := repository.db.ExecContext(ctx, q, category.Id)
	helper.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, categoryId int) (domain.Category, error) {
	q := "select id, name from category where id = ?"
	row := repository.db.QueryRowContext(ctx, q, categoryId)

	category := domain.Category{}
	err := row.Scan(&category.Id, &category.Name)
	return category, err
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context) []domain.Category {
	q := "select id, name from category"
	rows, err := repository.db.QueryContext(ctx, q)
	helper.PanicIfError(err)

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(
			&category.Id,
			&category.Name,
		)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}
