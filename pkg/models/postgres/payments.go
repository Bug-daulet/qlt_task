package postgres

import (
	"context"
	"fmt"
	"github.com/Bug-daulet/qlt_task/pkg/models"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
)

type PaymentRepository struct {
	Pool *pgxpool.Pool
}

const (
	queryGetAllPayments = "SELECT * FROM payment"
	queryGetPayment     = "SELECT * FROM payment WHERE id=$1"
	queryInsertPayment  = "INSERT INTO payment (title, date, type, comments, category_id) VALUES ($1,$2,$3,$4,$5) RETURNING id"
	queryUpdatePayment  = "UPDATE payment SET title=$1, date=$2, type=$3, comments=$4 WHERE id=$5"
	queryDeletePayment  = "DELETE FROM payment WHERE id=$1"
)

func (r *PaymentRepository) GetAll() []*models.Payment {
	rows, err := r.Pool.Query(context.Background(), queryGetAllPayments)
	if err != nil {
		fmt.Fprint(os.Stderr, "QueryRow failed: %v\n", err)
	}
	defer rows.Close()

	var payments []*models.Payment
	for rows.Next() {
		payment := &models.Payment{}
		err = rows.Scan(&payment.ID, &payment.Title, &payment.Date, &payment.Type, &payment.Comments, &payment.CategoryId)
		if err != nil {
			fmt.Fprint(os.Stderr, "Error: %v\n", err)
		}
		payments = append(payments, payment)
	}
	if err = rows.Err(); err != nil {
		fmt.Fprint(os.Stderr, "Error: %v\n", err)
	}
	return payments
}

func (r *PaymentRepository) Get(id int) *models.Payment {
	row := r.Pool.QueryRow(context.Background(), queryGetPayment, id)
	payment := &models.Payment{}
	err := row.Scan(&payment.ID, &payment.Title, &payment.Date, &payment.Type, &payment.Comments, &payment.CategoryId)
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: %v\n", err)
	}
	return payment
}

func (r *PaymentRepository) Save(payment *models.Payment) int {
	var paymentId int
	row := r.Pool.QueryRow(context.Background(), queryInsertPayment, payment.Title, payment.Date, payment.Type, payment.Comments, payment.CategoryId)
	err := row.Scan(&paymentId)
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: %v\n", err)
	}
	payment.ID = paymentId
	return paymentId
}

func (r *PaymentRepository) Update(payment *models.Payment) string {
	_, err := r.Pool.Exec(context.Background(), queryUpdatePayment, payment.Title, payment.Date, payment.Type, payment.Comments, payment.ID)
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: %v\n", err)
	}
	return "successfully updated"
}

func (r *PaymentRepository) Delete(id int) string {
	_, err := r.Pool.Exec(context.Background(), queryDeletePayment, id)
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: %v\n", err)
	}
	return "successfully deleted"
}
