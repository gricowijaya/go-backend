package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	// run a concurrent goroutine
	n := 5
	amount := int64(10)
	// send the channel into errors
	errs := make(chan error)
	// send the results into results
	results := make(chan TransferTXResult)
	for i := 0; i < n; i++ {
		go func() {
			// this function return a result or an error
			result, err := store.TransferTX(context.Background(), TransferTXParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})
			// send data into error
			errs <- err
			// send data into results
			results <- result
		}()
	}

	// check the results
	for i := 0; i < n; i++ {
		// error from the channel
		err := <-errs
		require.NoError(t, err)

		// result from the channel
		result := <-results
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, account1.ID, transfer.FromAccountID)
		require.Equal(t, account2.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		// check the transafer in the database
		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		//check the account entries from
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, account1.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount) // money is going out
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		_, err = store.GetEntries(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		//check the account entries
		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, account2.ID, toEntry.AccountID) // to accountID
		require.Equal(t, amount, toEntry.Amount)         // money is going in
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = store.GetEntries(context.Background(), toEntry.ID)
		require.NoError(t, err)

		// TODO check the account balance
	}
}
