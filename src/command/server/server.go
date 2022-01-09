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
	log.Println("======> The GO gRPC SERVER is UP!")
	log.Printf("======> Server listening at %v", listener.Addr())
	fmt.Print("\n\n")

	//Enables accepting incoming connections in the Listen port
	return s.Serve(listener)
}

func (server *CryptoServiceServer) CreateNewCrypto(ctx context.Context, request *pb.CreateNewCryptoRequest) (*pb.CreateNewCryptoResponse, error) {
	log.Println("======> Received from client: New Crypto")
	log.Printf("======> NAME: %v", request.GetName())
	log.Printf("======> TOKEN: %v", request.GetToken())
	fmt.Print("\n\n")

	crypto_created := repo.AddCrypto(server.conn.DB, m.CryptoCurrency{
		Name:  request.GetName(),
		Token: request.GetToken(),
		Votes: 0,
	})

	log.Println("======> Persisted in the Database")
	log.Printf("======> ID: %v", crypto_created.Id)
	log.Printf("======> NAME: %v", crypto_created.Name)
	log.Printf("======> TOKEN: %v", crypto_created.Token)
	log.Printf("======> VOTES: %v", crypto_created.Votes)
	fmt.Print("\n\n")

	return &pb.CreateNewCryptoResponse{
		Crypto: &pb.Crypto{
			Id:    int32(crypto_created.Id),
			Name:  crypto_created.Name,
			Token: crypto_created.Token,
			Votes: int32(crypto_created.Votes),
		},
	}, nil
}

func (server *CryptoServiceServer) ListCryptos(ctx context.Context, request *pb.ListCryptosRequest) (*pb.ListCryptosResponse, error) {
	var cryptos []m.CryptoCurrency

	switch request.Sortparam {
	case "name":
		cryptos = repo.FindAllCryptosSortedByName(server.conn.DB)
	case "token":
		cryptos = repo.FindAllCryptosSortedByToken(server.conn.DB)
	case "votes":
		cryptos = repo.FindAllCryptosSortedByTopVotes(server.conn.DB)
	default:
		cryptos = repo.FindAllCryptos(server.conn.DB)
	}

	log.Print("======> Crypto List from Database\n")
	log.Println(cryptos)
	fmt.Print("\n\n")

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

func (server *CryptoServiceServer) DeleteCrypto(ctx context.Context, request *pb.DeleteCryptoRequest) (*pb.EmptyResponse, error) {
	repo.RemoveCryptoById(server.conn.DB, int(request.GetId()))
	log.Print("======> Crypto Deleted By Id: ", request.GetId())
	fmt.Print("\n\n")
	return &pb.EmptyResponse{}, nil
}

func (server *CryptoServiceServer) UpdateCrypto(ctx context.Context, request *pb.UpdateCryptoRequest) (*pb.UpdateCryptoResponse, error) {
	crypto_old := repo.FindCryptoById(server.conn.DB, int(request.GetCrypto().GetId()))

	log.Println("======> Received from client: Update Crypto")
	log.Println("======> Old state")
	log.Printf("======> ID: %v", crypto_old.Id)
	log.Printf("======> NAME: %v", crypto_old.Name)
	log.Printf("======> TOKEN: %v", crypto_old.Token)
	log.Printf("======> VOTES: %v", crypto_old.Votes)
	fmt.Print("\n\n")

	crypto_updated := repo.UpdateCrypto(server.conn.DB, m.CryptoCurrency{
		Id:    int(request.GetCrypto().GetId()),
		Name:  request.GetCrypto().GetName(),
		Token: request.GetCrypto().GetToken(),
		Votes: int(request.GetCrypto().GetVotes()),
	})

	log.Println("======> New state")
	log.Printf("======> ID: %v", crypto_updated.Id)
	log.Printf("======> NAME: %v", crypto_updated.Name)
	log.Printf("======> TOKEN: %v", crypto_updated.Token)
	log.Printf("======> VOTES: %v", crypto_updated.Votes)
	fmt.Print("\n\n")

	return &pb.UpdateCryptoResponse{
		Crypto: &pb.Crypto{
			Id:    int32(crypto_updated.Id),
			Name:  crypto_updated.Name,
			Token: crypto_updated.Token,
			Votes: int32(crypto_updated.Votes),
		},
	}, nil
}

func (server *CryptoServiceServer) UpvoteCrypto(ctx context.Context, request *pb.UpvoteCryptoRequest) (*pb.EmptyResponse, error) {
	repo.UpvoteCryptoById(server.conn.DB, int(request.GetId()))
	log.Print("======> Upvote Registred to Crypto Id: ", request.GetId())
	fmt.Print("\n\n")
	return &pb.EmptyResponse{}, nil
}

func (server *CryptoServiceServer) DownvoteCrypto(ctx context.Context, request *pb.DownvoteCryptoRequest) (*pb.EmptyResponse, error) {
	repo.DownvoteCryptoById(server.conn.DB, int(request.GetId()))
	log.Print("======> Downvote Registred to Crypto Id: ", request.GetId())
	fmt.Print("\n\n")
	return &pb.EmptyResponse{}, nil
}

func (server *CryptoServiceServer) GetCryptoById(ctx context.Context, request *pb.GetCryptoByIdRequest) (*pb.GetCryptoByIdResponse, error) {
	crypto_found := repo.FindCryptoById(server.conn.DB, int(request.GetId()))
	log.Print("======> Crypto Found By Id: ", request.GetId())
	fmt.Print("\n\n")
	return &pb.GetCryptoByIdResponse{
		Crypto: &pb.Crypto{
			Id:    int32(crypto_found.Id),
			Name:  crypto_found.Name,
			Token: crypto_found.Token,
			Votes: int32(crypto_found.Votes),
		},
	}, nil
}

// func (server *CryptoServiceServer) ObserveCrypto(ctx context.Context, request *pb.ObserveCryptoRequest) (pb.CryptoService_ObserveCryptoClient, error) {
// 	return nil, nil
// }

func main() {
	fmt.Print("\n\n")
	log.Println("======> STARTING THE gRPC SERVER")
	var cryptoServer *CryptoServiceServer = NewCryptoServiceServer()
	cryptoServer.conn = repo.DatabaseInit()
	log.Println("======> Databased Ready: \n", cryptoServer.conn.DB.Stats())
	fmt.Print("\n\n")

	//repo.BuildDatabase(cryptoServer.conn.DB, true)
	defer cryptoServer.conn.DB.Close()

	if err := cryptoServer.Run(); err != nil {
		log.Fatalf("======> Failed to serve: \n%v", err)
	}
}
