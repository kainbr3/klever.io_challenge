// package server

// import (
// 	"context"
// 	"database/sql"
// 	"log"
// 	"net"
// 	"testing"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	m "github.com/kainbr3/klever.io_challenge/src/package/model"
// 	repo "github.com/kainbr3/klever.io_challenge/src/package/repository"
// 	pb "github.com/kainbr3/klever.io_challenge/src/protobuf"
// 	grpc "google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"
// 	"google.golang.org/grpc/test/bufconn"
// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/assert"
// )

// //Configuration for Enable Testing
// const bufSize = 1024 * 1024

// var lis *bufconn.Listener

// //var ctx context.Context

// func init() {
// 	lis = bufconn.Listen(bufSize)
// 	s := grpc.NewServer()
// 	pb.RegisterCryptoServiceServer(s, &CryptoServiceServer{})

// 	go func() {
// 		if err := s.Serve(lis); err != nil {
// 			log.Fatalf("Server exited with error: %v", err)
// 		}
// 	}()
// }

// func bufDialer(context.Context, string) (net.Conn, error) {
// 	return lis.Dial()
// }

// //Invalid Requests
// var mockedInvalidCreateCryptoRequest = pb.CreateNewCryptoRequest{}

// //Valid Requests
// var mockedValidCreateCryptoRequest = pb.CreateNewCryptoRequest{
// 	Name:  meckedValidCrypto.Name,
// 	Token: meckedValidCrypto.Token,
// }

// //Valid Responses
// var mockedInvalidCreateCryptoResponse = pb.CreateNewCryptoResponse{}
// var mockedValidCryptoResponse = pb.CreateNewCryptoResponse{
// 	Crypto: &mockedValidPbCrypto,
// }

// func NewMock() (*sql.DB, sqlmock.Sqlmock) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}

// 	return db, mock
// }

// //Tests
// func TestCreateNewCrypto(t *testing.T) {
// 	var cryptoServer *CryptoServiceServer = NewCryptoServiceServer()
// 	cryptoServer.conn = repo.DatabaseInit()
// 	ctx := context.Background()
// 	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

// 	log.Println(conn)
// 	log.Println(err)

// 	if err != nil {
// 		t.Errorf("Failed to dial bufnet: %v", err)
// 	}

// 	client := pb.NewCryptoServiceClient(conn)

// 	resp, err := client.CreateNewCrypto(ctx, &pb.CreateNewCryptoRequest{
// 		Name:  meckedValidCrypto.Name,
// 		Token: meckedValidCrypto.Token})
// 	if err != nil {
// 		t.Errorf("CreateNewCrypto failed: %v", err)
// 	}
// 	log.Printf("Response: %+v", resp)
// }

package server

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	m "github.com/kainbr3/klever.io_challenge/src/package/model"
	repo "github.com/kainbr3/klever.io_challenge/src/package/repository"
	"github.com/kainbr3/klever.io_challenge/src/package/tool"
	pb "github.com/kainbr3/klever.io_challenge/src/protobuf"
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
var mockedEmptyCrypto = m.CryptoCurrency{}
var mockedInvalidPbCrypto = pb.Crypto{}

//Valid Models
var mockedValidCrypto = m.CryptoCurrency{
	Id:    91,
	Name:  "Test Valid Crypto",
	Token: "TVC1",
	Votes: 55,
}
var mockedValidPbCrypto = pb.Crypto{
	Id:    91,
	Name:  "Test Valid Crypto",
	Token: "TVC1",
	Votes: 55,
}

//Valid Requests
var mockedValidCreateCryptoRequest = pb.CreateNewCryptoRequest{
	Name:  mockedValidCrypto.Name,
	Token: mockedValidCrypto.Token,
}

//Valid Responses
var mockedInvalidCreateCryptoResponse = pb.CreateNewCryptoResponse{}
var mockedValidCryptoResponse = pb.CreateNewCryptoResponse{
	Crypto: &mockedValidPbCrypto,
}

//Success Case
func TestCreateNewCrypto(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	db, mock := MockDatabase()
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec(tool.InsertCryptoQuery).WithArgs(mockedValidCrypto)
	mock.ExpectCommit()

	crypto_created, err := repo.AddCrypto(db, ctx, mockedValidCrypto)
	if err != nil {
		log.Println(err)
	}

	assert.Equal(t, mockedValidCrypto.Id, crypto_created.Id)
	assert.Equal(t, mockedValidCrypto.Name, crypto_created.Name)
	assert.Equal(t, mockedValidCrypto.Token, crypto_created.Token)
	assert.Equal(t, mockedValidCrypto.Votes, crypto_created.Votes)
}

func TestCreateNewCryptoSuccess(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	db, mock := MockDatabase()
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec(tool.InsertCryptoQuery).WithArgs(mockedValidCrypto)
	//mock.ExpectExec("use kleverchallenge INSERT INTO cryptoCurrencies (name, token, votes) OUTPUT Inserted.id, Inserted.name, Inserted.token, Inserted.votes VALUES ($1, $2, $3);").WithArgs(mockedValidCrypto.Name, mockedValidCrypto.Token, mockedValidCrypto.Votes)
	mock.ExpectCommit()

	Instance = &CryptoServiceServer{}
	Instance.conn = &repo.Klever{DB: db}

	response, err := Instance.CreateNewCrypto(ctx, &mockedValidCreateCryptoRequest)
	fmt.Println(err)
	fmt.Println(response)
	fmt.Println(response.Crypto.GetName())

	assert.Equal(t, nil, err)
	assert.Equal(t, mockedValidCryptoResponse.Crypto.Name, response.Crypto.Name)
	assert.Equal(t, mockedValidCryptoResponse.Crypto.Token, response.Crypto.Token)
	assert.Equal(t, 0, int(response.Crypto.Votes))

}
