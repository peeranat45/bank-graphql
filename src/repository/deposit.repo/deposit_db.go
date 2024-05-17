package depositrepo

import (
	"bank/src/pkg/data"
	"bank/src/pkg/data/models"
)

type DepositRepo struct {
	db *data.DB
}

func NewDepositRepo(db *data.DB) DepositRepository {
	return &DepositRepo{db: db}
}

func (d *DepositRepo) GetDeposits() ([]*models.DepositeTransaction, error){
	rows, err := d.db.Connection.Query("SELECT * FROM DepositeTransaction")
	if err != nil {
		return nil, err
	}
	var deposits []*models.DepositeTransaction
	for rows.Next() {
		var deposit models.DepositeTransaction
		rows.Scan(&deposit.TransactionID, &deposit.AccountID, &deposit.TransactionDate, &deposit.Amount, &deposit.Description)
		deposits = append(deposits, &deposit)
	}
	return deposits, nil
	
}

func (d* DepositRepo) CreateDeposit(dt models.DepositeTransaction) (*models.DepositeTransaction, error) {
	query := "INSERT INTO DepositeTransaction (AccountID, TransactionDate, Amount, Description) VALUES (?, ?, ?, ?)"
	res, err := d.db.Connection.Exec(query, dt.AccountID, dt.TransactionDate, dt.Amount, dt.Description)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	dt.AccountID = int(id)

	return &dt, nil
}