package database

import (
	"strings"

	"github.com/joho/sqltocsv"
	"github.com/ssergomol/Balance-Manager/pkg/models"
)

type OrderRepo struct {
	store *Storage
}

func (r *OrderRepo) CreateOrder(order models.Order) error {
	_, err := r.store.db.Query("INSERT INTO orders(user_id, service_id, is_positive, price, description)"+
		" VALUES ($1, $2, $3, $4, $5)",
		order.UserID, order.ServiceID, order.IsPositive, order.Price, order.Description,
	)
	return err
}

func (r *OrderRepo) GetReport(date string) (string, error) {
	dateStr := strings.Split(date, "-")
	year := dateStr[0]
	month := dateStr[1]
	rows, err := r.store.db.Query("SELECT SUM(price), description, execution_date FROM orders WHERE"+
		" EXTRACT (MONTH FROM execution_date) = $1 AND EXTRACT (YEAR FROM execution_date) = $2"+
		" GROUP BY description, execution_date", month, year)

	path := "reports/report_" + year + "_" + month
	if err != nil {
		return path, err
	}
	defer rows.Close()

	// f, err := os.Create("reports/report_" + year + "_" + month)
	// if err != nil {
	// 	return path, err
	// }
	// f.Close()

	err = sqltocsv.WriteFile(path, rows)
	if err != nil {
		return path, err
	}

	return path, nil
}
