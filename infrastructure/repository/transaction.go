package repository

import "database/sql"

type TransactionRepositoryDb struct {
	db *sql.db
}

// Creating DB
func NewTransactionRepositoryDb(db *sql.DB) *TransactionRepositoryDb{
	return &TransactionRepositoryDb{db:db}
}

func (t *TransactionRepositoryDb) SaveTransaction( transaction domain.Transaction, creditCard domain.CreditCard ) error {

}