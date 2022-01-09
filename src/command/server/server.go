package main

import (
	"context"
	"fmt"
	"log"
	"net"

	m "github.com/kainbr3/klever.io_challenge/package/model"
	repo "github.com/kainbr3/klever.io_challenge/package/repository"
	t "github.com/kainbr3/klever.io_challenge/package/tool"
	pb "github.com/kainbr3/klever.io_challenge/protobuf"
	grpc "google.golang.org/grpc"
)

func NewCryptoServiceServer() *CryptoServiceServer {
	return &CryptoServiceServer{}
}

type CryptoServiceServer struct {
	conn *repo.Klever
	pb.UnimplementedCryptoServiceServer
}

func (server *CryptoServiceServer) Run() error {
	//Starts to Listen TCP Port 5001
	listener, err := net.Listen(t.ServerNetworkType, t.ServerPort)
	if err != nil {
		log.Fatalf("======> Failed to listen: %v", err)
	}

	//Starts the gRPC Server
	s := grpc.NewServer()
	pb.RegisterCryptoServiceServer(s, server)

	//Show some message to displat that server is online
	log.Printf("======> The GO gRPC SERVER is UP!")
	log.Printf("======> Server listening at %v", listener.Addr())

	//Enables accepting incoming connections in the Listen port
	return s.Serve(listener)
}

func (server *CryptoServiceServer) CreateNewCrypto(ctx context.Context, request *pb.NewCryptoRequest) (*pb.NewCryptoResponse, error) {
	log.Printf("======> Received from client: New Crypto")
	log.Printf("======> NAME: %v", request.GetName())
	log.Printf("======> TOKEN: %v", request.GetToken())

	crypto_created := repo.AddCrypto(server.conn.DB, m.CryptoCurrency{
		Name:  request.GetName(),
		Token: request.GetToken(),
		Votes: 0,
	})

	fmt.Print("======> Persisted in the Database")
	log.Printf("======> ID: %v", &crypto_created.Id)
	log.Printf("======> NAME: %v", &crypto_created.Name)
	log.Printf("======> TOKEN: %v", &crypto_created.Token)
	log.Printf("======> VOTES: %v", &crypto_created.Votes)

	return &pb.NewCryptoResponse{
		Crypto: &pb.Crypto{
			Id:    int32(crypto_created.Id),
			Name:  crypto_created.Name,
			Token: crypto_created.Token,
			Votes: int32(crypto_created.Votes),
		},
	}, nil
}

func (server *CryptoServiceServer) GetCryptos(ctx context.Context, request *pb.ListCryptosRequest) (*pb.ListCryptosResponse, error) {
	cryptos := repo.FindAllCryptosSortedByTopVotes(server.conn.DB)
	fmt.Print(cryptos)

	var cryptos_list *pb.ListCryptosResponse = &pb.ListCryptosResponse{}

	for _, obj := range cryptos {
		crypto := &pb.Crypto{
			Id:    int32(obj.Id),
			Name:  obj.Name,
			Token: obj.Token,
			Votes: int32(obj.Votes),
		}
		cryptos_list.Cryptos = append(cryptos_list.Cryptos, crypto)
	}

	return cryptos_list, nil
}

func main() {
	var cryptoServer *CryptoServiceServer = NewCryptoServiceServer()
	cryptoServer.conn = repo.DatabaseInit()
	fmt.Println("======> Databased Ready: ", cryptoServer.conn.DB.Stats())
	//repo.BuildDatabase(cryptoServer.conn.DB, true)
	defer cryptoServer.conn.DB.Close()

	if err := cryptoServer.Run(); err != nil {
		log.Fatalf("======> Failed to serve: \n%v", err)
	}
}

func RepositoryTest() {
	//Connect to the Database
	// klever := repo.DatabaseInit()
	// fmt.Println("======> Databased Ready: ", klever.DB.Stats())

	//Function to Create the main Tables and Populate it with some Data
	//repo.BuildDatabase(klever.DB, true)

	//Some Query Tests
	// fmt.Println("======> Query Result:", repo.FindAllCryptos(klever.DB))
	// fmt.Println("======> Query Result:", repo.FindAllCryptosSortedByName(klever.DB))
	// fmt.Println("======> Query Result:", repo.FindAllCryptosSortedByToken(klever.DB))
	// fmt.Println("======> Query Result:\n\n", repo.FindAllCryptosSortedByTopVotes(klever.DB))
	// fmt.Println("======> Query Result:\n\n", repo.FindAllCryptosSortedByLeastVotes(klever.DB))
	// fmt.Println("======> Query Result:", repo.FindCryptoById(klever.DB, 1))
	// fmt.Println("======> Query Result:", repo.FindCryptoByName(klever.DB, "Klever"))
	// fmt.Println("======> Query Result:", repo.FindCryptoByToken(klever.DB, "BTC"))

	//repo.AddCrypto(klever.DB, m.CryptoCurrency{Name: "Teste Crypto", Token: "TST"})
	//repo.RemoveCryptoById(klever.DB, 8)
	//repo.CleanTable(klever.DB, "cryptoCurrencies")
	//repo.ReseedTable(klever.DB, "cryptoCurrencies")
	//repo.DropTable(klever.DB, "cryptoCurrencies")
	//repo.ExecuteCustomStatement(klever.DB, t.DropAllTables)
	//repo.ExecuteCustomStatement(klever.DB, t.RatingTableQuery)
	// repo.UpvoteCryptoById(klever.DB, 7)
	// repo.DownvoteCryptoById(klever.DB, 3)
}
