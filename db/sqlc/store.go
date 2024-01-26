package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	// similiar to inheritance
	*Queries
	// sql db object
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execute a function within a database transaction, fn func(*Queries) error in this case
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil) // Begin takes in context and isolation options

	if err != nil {
		return nil
	}
	q := New(tx)

	err = fn(q)
	if err != nil {
		//check if Rollback happened
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("transaction error %w, rollback error: %w ", err, rbErr)
		}
		return err
	}
	// when all Tx successs, commit it
	return tx.Commit()
}

// TransferTx: Transaction to transfer money from one account to another
// create a transfer record, add entries, update accounts balance in one transaction
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Accounts `json:"from_account"`
	ToAccount   Accounts `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:" to_entry"`
}

func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, err) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccounID: arg.FromAccountID,
			Amount:   -arg.Amount,
		})
		if err != nill {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccounID: arg.FromAccountID,
			Amount:   -arg.Amount,
		})
		if err != nill {
			return err
		}

		return nil
	})

	return result, err
}
