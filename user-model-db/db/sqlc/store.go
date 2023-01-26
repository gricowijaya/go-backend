package db

import (
	"context"
	"database/sql"
	"fmt"
)

/*
provides functions to execute db entries and transactions
inherit function from Queries into Store
provide the sql db object to execute into the database
*/
type Store struct {
	*Queries
	db *sql.DB
}

/*
Create a New Store
DB is return a queries object
*/
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

/*
write a function to a Store called execTx so will won't export the function
call the callback function based on the error return by that function
*/
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		// rollback error
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	// commit the transaction if there's no error
	return tx.Commit()
}

// TransferTXParams contains the input params of the transfer transaction
type TransferTXParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// TransferTXResult is the results of the transfer transaction
type TransferTXResult struct {
	Transfer    Transfers `json:"transfer"`
	FromAccount Accounts  `json:"from_account"`
	ToAccount   Accounts  `json:"to_account"`
	FromEntry   Entries   `json:"from_entry"`
	ToEntry     Entries   `json:"to_entry"`
}

/*
write a function to a Store for creating transfer record
and update the account balance within a single database transaction
1. Create Transfer Record
2. Add Account Entries
3. Update Account Balance
*/
func (store *Store) TransferTX(ctx context.Context, arg TransferTXParams) (TransferTXResult, error) {
	var result TransferTXResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// create Transfer information with the transaction amount
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		// create the from entry which is out transaction
		result.FromEntry, err = q.CreateEntries(ctx, CreateEntriesParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}

		// create the into entry which is in transcaction
		result.ToEntry, err = q.CreateEntries(ctx, CreateEntriesParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		// get account balance -> Update Account balance using the locking mechanism
		result.FromAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			ID:     arg.FromAccountID,
			Amount: -arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			ID:     arg.ToAccountID,
			Amount: arg.Amount,
		})

		if err != nil {
			return err
		}
		return nil
	})
	return result, err
}
