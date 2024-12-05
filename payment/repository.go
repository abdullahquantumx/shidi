package account

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)
// RemittanceDetail represents the result of processing a single order's remittance.
type RemittanceDetail struct {
	OrderID   string
	Amount    float64
	Processed bool
}

// Transaction represents a single transaction record.
type Transaction struct {
	TransactionID string
	TransactionType string
	Amount         float64
	OrderID        sql.NullString
	Timestamp      time.Time
}


type Repository interface {
	Close()
	RechargeWallet(ctx context.Context, userID string, amount float64) (float64, error)
	DeductBalance(ctx context.Context, userID string, amount float64, orderID string) (float64, error)
	ProcessRemittance(ctx context.Context, userID string, orderIDs []string) ([]RemittanceDetail, error)
	GetWalletDetails(ctx context.Context, userID string) (float64, []Transaction, error)
}

type postgresRepository struct {
	db *sql.DB
}


func NewPostgresRepository(url string) (*postgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &postgresRepository{db}, nil
}

func (r *postgresRepository) Close() {
	r.db.Close()
}

func (r *postgresRepository) RechargeWallet(ctx context.Context, userID string, amount float64) (float64, error) {
	
}

func (r *postgresRepository) DeductBalance(ctx context.Context, userID string, amount float64, orderID string) (float64, error) {

}

func (r *postgresRepository) ProcessRemittance(ctx context.Context, userID string, orderIDs []string) ([]RemittanceDetail, error) {

}

func (r *postgresRepository) GetWalletDetails(ctx context.Context, userID string) (float64, []Transaction, error) {

}

