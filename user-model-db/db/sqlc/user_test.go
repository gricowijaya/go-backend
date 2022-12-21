package db

import (
	"context"
	"database/sql"
	"log"
	"testing"
	"user-model-db/app/util"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User { 
	arg := CreateUserParams{
		// reference : https://stackoverflow.com/questions/60792313/unable-to-use-type-string-as-sql-nullstring
		Email:    sql.NullString{String: util.RandomOwner(), Valid: true},
		Username: sql.NullString{String: util.RandomOwner(), Valid: true},
		Password: sql.NullString{String: "super-password", Valid: true},
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)   // check the error must be nil
	require.NotEmpty(t, user) // user should be not an empty object

	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.Username, user.Username)
  
  return user
}

func TestCreateUser(t *testing.T) {
  createRandomUser(t);
}

func TestGetOneUser(t *testing.T) {
  createUser := createRandomUser(t);  
  user, err := testQueries.GetOneUser(context.Background(), createUser.ID)
  if (err != nil) { 
    log.Fatal(err)
  }

  require.Equal(t, createUser.ID, user.ID)
  require.Equal(t, createUser.Email, user.Email)
  require.Equal(t, createUser.Password, user.Password)
  require.Equal(t, createUser.Username, user.Username)
}

func TestGetAllUsers(t *testing.T) {
  allUser, err := testQueries.GetAllUsers(context.Background())
  require.NoError(t, err)
  require.Len(t, allUser, len(allUser))
}

func TestUpdateEmail(t *testing.T) { 
  createUser := createRandomUser(t)

  arg := UpdateEmailParams{
    ID: createUser.ID,
    Email: createUser.Email,
  }

  err := testQueries.UpdateEmail(context.Background(), arg)

  if (err !=nil ) { 
    log.Fatal(err)
  }
}

func TestDeleteUserTestDeleteUser(t *testing.T) {
  createUser := createRandomUser(t)

  err := testQueries.DeleteUser(context.Background(), createUser.ID)
  if err != nil { log.Fatal(err) }

  // for checking the user that we create
  checkUser, err := testQueries.GetOneUser(context.Background(), createUser.ID)
  require.Error(t, err)
  require.Empty(t, checkUser)
}

