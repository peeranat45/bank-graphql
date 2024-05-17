package depositrepo

import "bank/src/pkg/data/models"

type DepositRepository interface {
	GetDeposits() ([]*models.DepositeTransaction, error)
	CreateDeposit(models.DepositeTransaction) (*models.DepositeTransaction, error)
}