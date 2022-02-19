package postgres

import (
	"context"
	"fmt"
	"github.com/Bug-daulet/qlt_task/pkg/models"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
)

type CategoryRepository struct {
	Pool *pgxpool.Pool
}

const (
	queryGetAllCategories	= "SELECT * FROM category"
	queryGetCategory     	= "SELECT * FROM category WHERE id=$1"
	queryInsertCategory  	= "INSERT INTO category (name) VALUES ($1) RETURNING id"
	queryUpdateCategory  	= "UPDATE category SET name=$1 WHERE id=$2"
	queryDeleteCategory  	= "DELETE FROM category WHERE id=$1"
)

func (r *CategoryRepository) GetAll() []*models.Category {
	rows, err := r.Pool.Query(context.Background(), queryGetAllCategories)
	if err != nil {
		fmt.Fprint(os.Stderr, "QueryRow failed: %v\n", err)
	}
	defer rows.Close()

	var categories []*models.Category
	for rows.Next() {
		category := &models.Category{}
		err = rows.Scan(&category.ID, &category.Name)
		if err != nil {
			fmt.Fprint(os.Stderr, "Error: %v\n", err)
		}
		categories = append(categories, category)
	}
	if err = rows.Err(); err != nil {
		fmt.Fprint(os.Stderr, "Error: %v\n", err)
	}
	return categories
}

func (r *CategoryRepository) Get(id int) *models.Category {
	row := r.Pool.QueryRow(context.Background(), queryGetCategory, id)
	category := &models.Category{}
	err := row.Scan(&category.ID, &category.Name)
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: %v\n", err)
	}
	return category
}

func (r *CategoryRepository) Save(category *models.Category) int {
	var categoryId int
	row := r.Pool.QueryRow(context.Background(), queryInsertCategory, category.Name)
	err := row.Scan(&categoryId)
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: %v\n", err)
	}
	category.ID = categoryId
	return categoryId
}

func (r *CategoryRepository) Update(category *models.Category) string {
	_, err := r.Pool.Exec(context.Background(), queryUpdateCategory, category.Name, category.ID)
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: %v\n", err)
	}
	return "successfully updated"
}

func (r *CategoryRepository) Delete(id int) string {
	_, err := r.Pool.Exec(context.Background(), queryDeleteCategory, id)
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: %v\n", err)
	}
	return "successfully deleted"
}
