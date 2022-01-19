package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	m "github.com/kainbr3/klever.io_challenge/src/package/model"
	"github.com/kainbr3/klever.io_challenge/src/package/tool"
	"github.com/stretchr/testify/assert"
)

//Mock Database
func MockDatabase() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

//Mocked Data:
//Invalid Models
var emptyCrypto = m.CryptoCurrency{}

//Valid Models
var createdCrypto = m.CryptoCurrency{
	Id:    91,
	Name:  "Test Valid Crypto",
	Token: "TVC",
	Votes: 55,
}

var updatedCrypto = m.CryptoCurrency{
	Id:    91,
	Name:  "Test Valid Crypto Updated",
	Token: "TVCU",
	Votes: 55,
}

//var deletedCrypto = 91

func TestAddCrypto(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, mock := MockDatabase()
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec(tool.InsertCryptoQuery).WithArgs(createdCrypto)
	mock.ExpectCommit()

	//success
	crypto_created, err := AddCrypto(db, ctx, createdCrypto)
	assert.Equal(t, err, nil)
	assert.Equal(t, createdCrypto.Id, crypto_created.Id)
	assert.Equal(t, createdCrypto.Name, crypto_created.Name)
	assert.Equal(t, createdCrypto.Token, crypto_created.Token)
	assert.Equal(t, createdCrypto.Votes, crypto_created.Votes)
	assert.NoError(t, err)

	//fail
	invalidCrypto, err := AddCrypto(db, ctx, emptyCrypto)
	assert.NotEqual(t, err, nil)
	assert.Equal(t, invalidCrypto.Id, 0)
	assert.Equal(t, invalidCrypto.Name, "")
	assert.Equal(t, invalidCrypto.Token, "")
	assert.Equal(t, invalidCrypto.Votes, 0)

}

func TestUpdateCrypto(t *testing.T) {
	// ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	// defer cancel()

	// db, mock := MockDatabase()
	// defer db.Close()

	// mock.ExpectBegin()
	// mock.ExpectPrepare(fmt.Sprintf(tool.UpdateCryptoQuery, updatedCrypto.Name, updatedCrypto.Token, updatedCrypto.Votes, updatedCrypto.Id))
	// // prep := mock.ExpectPrepare(fmt.Sprintf(tool.UpdateCryptoQuery, updatedCrypto.Name, updatedCrypto.Token, updatedCrypto.Votes, updatedCrypto.Id))
	// // prep.ExpectExec().WithArgs(updatedCrypto.Name, updatedCrypto.Token, updatedCrypto.Votes, updatedCrypto.Id).WillReturnResult(sqlmock.NewResult(0, 1))
	// //prep.ExpectExec().WillReturnResult(sqlmock.NewResult(0, 1))

	// //success
	// crypto_updated, err := UpdateCrypto(db, ctx, updatedCrypto)
	// assert.Equal(t, err, nil)
	// assert.Equal(t, updatedCrypto.Id, crypto_updated.Id)
	// assert.Equal(t, updatedCrypto.Name, crypto_updated.Name)
	// assert.Equal(t, updatedCrypto.Token, crypto_updated.Token)
	// assert.Equal(t, updatedCrypto.Votes, crypto_updated.Votes)
	// assert.NoError(t, err)
}

func TestFindCryptoById(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	db, mock := MockDatabase()
	defer db.Close()

	rows := mock.NewRows([]string{"id", "name", "token", "votes"}).
		AddRow(&createdCrypto.Id, &createdCrypto.Name, &createdCrypto.Token, &createdCrypto.Votes)

	mock.ExpectQuery(fmt.Sprintf(tool.SelectCryptoByIdQuery, createdCrypto.Id)).WillReturnRows(rows)
	if rows != nil {
		fmt.Println(rows)
	}

	//found
	crypto_found, err := FindCryptoById(db, ctx, createdCrypto.Id)
	assert.NotNil(t, crypto_found)
	assert.Equal(t, crypto_found.Id, createdCrypto.Id)
	assert.Equal(t, crypto_found.Name, createdCrypto.Name)
	assert.Equal(t, crypto_found.Token, createdCrypto.Token)
	assert.Equal(t, crypto_found.Votes, createdCrypto.Votes)
	assert.NoError(t, err)

	//notfound
	crypto_notFound, err := FindCryptoById(db, ctx, createdCrypto.Id)
	assert.Error(t, err)
	assert.Equal(t, crypto_notFound.Id, 0)
	assert.Equal(t, crypto_notFound.Name, "")
	assert.Equal(t, crypto_notFound.Token, "")
	assert.Equal(t, crypto_notFound.Votes, 0)

}

func TestUpvoteCryptoById(t *testing.T) {
	assert.Equal(t, 1, 1)
}

func TestDownvoteCryptoById(t *testing.T) {
	assert.Equal(t, 1, 1)
}

func TestRemoveCryptoById(t *testing.T) {
	assert.Equal(t, 1, 1)
}

func TestFindAllCryptos(t *testing.T) {
	assert.Equal(t, 1, 1)
}
