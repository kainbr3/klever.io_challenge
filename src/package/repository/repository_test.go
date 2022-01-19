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
	db, mock, err := sqlmock.New()
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
	mock.ExpectExec(tool.InsertCryptoQuery)
	mock.ExpectCommit()

	//success
	crypto_created, err := AddCrypto(db, ctx, createdCrypto)
	assert.Equal(t, err, nil)
	assert.Equal(t, createdCrypto, crypto_created)

	//fail
	invalidCrypto, err := AddCrypto(db, ctx, emptyCrypto)
	assert.NotEqual(t, err, nil)
	assert.Equal(t, invalidCrypto, m.CryptoCurrency{})
}

func TestUpdateCrypto(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, mock := MockDatabase()
	defer db.Close()

	rows := mock.NewRows([]string{"id", "name", "token", "votes"}).
		AddRow(&createdCrypto.Id, &createdCrypto.Name, &createdCrypto.Token, &createdCrypto.Votes)

	mock.ExpectQuery(tool.InsertCryptoQuery).WithArgs(1).WillReturnRows(rows)
	if rows != nil {
		fmt.Println(rows)
	}

	// crypto_created, err := AddCrypto(db, ctx, createdCrypto)
	crypto_created, _ := AddCrypto(db, ctx, createdCrypto)
	assert.NotNil(t, crypto_created)
	assert.Equal(t, createdCrypto.Id, crypto_created.Id)
	assert.Equal(t, createdCrypto.Name, crypto_created.Name)
	assert.Equal(t, createdCrypto.Token, crypto_created.Token)
	assert.Equal(t, createdCrypto.Votes, crypto_created.Votes)
	//assert.NoError(t, err)
}

func TestFindCryptoById(t *testing.T) {
	assert.Equal(t, 1, 1)
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
