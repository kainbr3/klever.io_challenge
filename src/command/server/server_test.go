package server

import (
	"testing"

	m "github.com/kainbr3/klever.io_challenge/src/package/model"
	pb "github.com/kainbr3/klever.io_challenge/src/protobuf"
)

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

//Mocked Data:
//Invalid Models
//var mockedEmptyCrypto = m.CryptoCurrency{}
//var mockedInvalidPbCrypto = pb.Crypto{}

//Valid Models
var meckedValidCrypto = m.CryptoCurrency{
	Id:    91,
	Name:  "Test Valid Crypto",
	Token: "TVC1",
	Votes: 55,
}
var mockedValidPbCrypto = pb.Crypto{
	Id:    92,
	Name:  "Teste Valid Protobuf Crypto",
	Token: "TVC2",
	Votes: 12,
}

//Invalid Requests
//var mockedInvalidCreateCryptoRequest = pb.CreateNewCryptoRequest{}

//Valid Requests
var mockedValidCreateCryptoRequest = pb.CreateNewCryptoRequest{
	Name:  meckedValidCrypto.Name,
	Token: meckedValidCrypto.Token,
}

//Valid Responses
//var mockedInvalidCreateCryptoResponse = pb.CreateNewCryptoResponse{}
var mockedValidCryptoResponse = pb.CreateNewCryptoResponse{
	Crypto: &mockedValidPbCrypto,
}

//Tests
func TestCreateNewCrypto(t *testing.T) {
	// ctx := context.Background()
	// conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

	// log.Println(conn)
	// log.Println(err)

	// if err != nil {
	// 	t.Errorf("Failed to dial bufnet: %v", err)
	// }

	// client := pb.NewCryptoServiceClient(conn)

	// resp, err := client.CreateNewCrypto(ctx, &pb.CreateNewCryptoRequest{
	// 	Name:  meckedValidCrypto.Name,
	// 	Token: meckedValidCrypto.Token})
	// if err != nil {
	// 	t.Errorf("CreateNewCrypto failed: %v", err)
	// }
	// log.Printf("Response: %+v", resp)
}
