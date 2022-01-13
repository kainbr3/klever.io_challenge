package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	m "github.com/kainbr3/klever.io_challenge/src/package/model"
	repo "github.com/kainbr3/klever.io_challenge/src/package/repository"
	t "github.com/kainbr3/klever.io_challenge/src/package/tool"
	pb "github.com/kainbr3/klever.io_challenge/src/protobuf"
	grpc "google.golang.org/grpc"
	ref "google.golang.org/grpc/reflection"
)

//Function to return the Server Instance
func NewCryptoServiceServer() *CryptoServiceServer {
	return &CryptoServiceServer{}
}

//Struc to store the Server Instance
type CryptoServiceServer struct {
	conn *repo.Klever
	pb.UnimplementedCryptoServiceServer
}

//Function to Start the Server
func (server *CryptoServiceServer) Run() error {
	//Starts to Listen TCP Port 5001
	listener, err := net.Listen(t.ServerNetworkType, t.ServerPort)
	if err != nil {
		log.Fatalf("======> Failed to listen: %v", err)
	}

	//Starts the gRPC Server
	s := grpc.NewServer()
	ref.Register(s)
	pb.RegisterCryptoServiceServer(s, server)

	//Show some message to displat that server is online
	log.Println("======> The GO gRPC SERVER is UP!")
	log.Printf("======> Server listening at %v", listener.Addr())
	fmt.Print("\n\n")

	//Enables accepting incoming connections in the Listen port
	return s.Serve(listener)
}

//Function to Insert a new Cryptocurrency in the database
func (server *CryptoServiceServer) CreateNewCrypto(ctx context.Context, request *pb.CreateNewCryptoRequest) (*pb.CreateNewCryptoResponse, error) {
	//Show the Response Parameters Received
	log.Println("======> Received from client: New Crypto")
	log.Printf("======> NAME: %v", request.GetName())
	log.Printf("======> TOKEN: %v", request.GetToken())
	fmt.Print("\n\n")

	//Insert the new Crypto in the Database
	crypto_created, err := repo.AddCrypto(server.conn.DB, ctx, m.CryptoCurrency{
		Name:  request.GetName(),
		Token: request.GetToken(),
		Votes: 0,
	})
	if err != nil {
		log.Printf("Error while saving crypto: \n%v", err)

		//Return the Response with the Null Crypto and the Errors
		return nil, err
	}

	//Print Crypto Details if success
	log.Println("======> Persisted in the Database")
	log.Printf("======> ID: %v", crypto_created.Id)
	log.Printf("======> NAME: %v", crypto_created.Name)
	log.Printf("======> TOKEN: %v", crypto_created.Token)
	log.Printf("======> VOTES: %v", crypto_created.Votes)
	fmt.Print("\n\n")

	//Return the Response with the New Crypto Added
	return &pb.CreateNewCryptoResponse{
		Crypto: &pb.Crypto{
			Id:    int32(crypto_created.Id),
			Name:  crypto_created.Name,
			Token: crypto_created.Token,
			Votes: int32(crypto_created.Votes),
		},
	}, nil
}

//Function to Update a Crypto
func (server *CryptoServiceServer) UpdateCrypto(ctx context.Context, request *pb.UpdateCryptoRequest) (*pb.UpdateCryptoResponse, error) {
	//Get the Crypto ID from the request
	cryptoId := int(request.GetCrypto().GetId())

	//Try to locate it
	crypto_old, err := repo.FindCryptoById(server.conn.DB, ctx, cryptoId)
	if err != nil {
		log.Printf("======> Could not locate Crypto ID = %d \n%v", cryptoId, err)

		//Return the Response with the Null Crypto and the Errors
		return nil, err
	}

	//Print Crypto Details with the Current State
	log.Println("======> Received from client: Update Crypto")
	log.Println("======> Old state")
	log.Printf("======> ID: %v", crypto_old.Id)
	log.Printf("======> NAME: %v", crypto_old.Name)
	log.Printf("======> TOKEN: %v", crypto_old.Token)
	log.Printf("======> VOTES: %v", crypto_old.Votes)
	fmt.Print("\n\n")

	//Update Crypto
	crypto_updated, err := repo.UpdateCrypto(server.conn.DB, ctx, m.CryptoCurrency{
		Id:    int(request.GetCrypto().GetId()),
		Name:  request.GetCrypto().GetName(),
		Token: request.GetCrypto().GetToken(),
		Votes: int(request.GetCrypto().GetVotes()),
	})
	if err != nil {
		log.Printf("Error while updating crypto: \n%v", err)

		//Return the Response with the Null Crypto and the Errors
		return nil, err
	}

	//Print Crypto Details with the New State if success
	log.Println("======> New state")
	log.Printf("======> ID: %v", crypto_updated.Id)
	log.Printf("======> NAME: %v", crypto_updated.Name)
	log.Printf("======> TOKEN: %v", crypto_updated.Token)
	log.Printf("======> VOTES: %v", crypto_updated.Votes)
	fmt.Print("\n\n")

	//Return the Response with the Updated Crypto Added
	return &pb.UpdateCryptoResponse{
		Crypto: &pb.Crypto{
			Id:    int32(crypto_updated.Id),
			Name:  crypto_updated.Name,
			Token: crypto_updated.Token,
			Votes: int32(crypto_updated.Votes),
		},
	}, nil
}

//Function to Find a Crypto By Its ID
func (server *CryptoServiceServer) GetCryptoById(ctx context.Context, request *pb.GetCryptoByIdRequest) (*pb.GetCryptoByIdResponse, error) {
	//Get the Crypto ID from the request
	cryptoId := int(request.GetId())

	//Try to locate it
	crypto_found, err := repo.FindCryptoById(server.conn.DB, ctx, cryptoId)
	if err != nil {
		log.Printf("======> Could not locate Crypto ID = %d \n%v", cryptoId, err)

		//Return the Response with the Null Crypto and the Errors
		return nil, err
	}

	//Print Crypto Details if success
	log.Print("======> Crypto Found By Id: ", request.GetId())
	fmt.Print("\n\n")

	//Return the Response with the Crypto Found
	return &pb.GetCryptoByIdResponse{
		Crypto: &pb.Crypto{
			Id:    int32(crypto_found.Id),
			Name:  crypto_found.Name,
			Token: crypto_found.Token,
			Votes: int32(crypto_found.Votes),
		},
	}, nil
}

//Function to Add a Vote to Given Crypto
func (server *CryptoServiceServer) UpvoteCrypto(ctx context.Context, request *pb.UpvoteCryptoRequest) (*pb.EmptyResponse, error) {
	//Get the Crypto ID from the request
	cryptoId := int(request.GetId())

	//Update the Crypto
	err := repo.UpvoteCryptoById(server.conn.DB, ctx, cryptoId)
	if err != nil {
		log.Printf("======> Could not locate Crypto ID = %d \n%v", cryptoId, err)

		//Return the Response with the Null Crypto and the Errors
		return nil, err
	}

	//Print the confirmation message if succeed
	log.Print("======> Upvote Registred to Crypto Id: ", cryptoId)
	fmt.Print("\n\n")

	//Return the Response with no Errors
	return &pb.EmptyResponse{}, nil
}

//Function to Subtract a Vote to Given Crypto
func (server *CryptoServiceServer) DownvoteCrypto(ctx context.Context, request *pb.DownvoteCryptoRequest) (*pb.EmptyResponse, error) {
	//Get the Crypto ID from the request
	cryptoId := int(request.GetId())

	//Update the Crypto
	err := repo.DownvoteCryptoById(server.conn.DB, ctx, cryptoId)
	if err != nil {
		log.Printf("======> Could not locate Crypto ID = %d \n%v", cryptoId, err)

		//Return the Response with the Null Crypto and the Errors
		return nil, err
	}

	//Print the confirmation message if succeed
	log.Print("======> Downvote Registred to Crypto Id: ", cryptoId)
	fmt.Print("\n\n")

	//Return the Response with no Errors
	return &pb.EmptyResponse{}, nil
}

//Function to Delete a Cryptocurrency from the database
func (server *CryptoServiceServer) DeleteCrypto(ctx context.Context, request *pb.DeleteCryptoRequest) (*pb.EmptyResponse, error) {
	//Get the Crypto ID from the request
	cryptoId := int(request.GetId())

	//Delete the Crypto
	err := repo.RemoveCryptoById(server.conn.DB, ctx, cryptoId)
	if err != nil {
		return nil, err
	}

	//Print the confirmation message if succeed
	log.Print("======> Crypto Deleted By Id: ", request.GetId())
	fmt.Print("\n\n")
	return &pb.EmptyResponse{}, nil
}

//Function to Stream a Crypto
func (server *CryptoServiceServer) ObserveCrypto(request *pb.ObserveCryptoRequest, streaming pb.CryptoService_ObserveCryptoServer) error {
	log.Printf("Streaming Crypto ID : %d", request.Id)

	//Find the Crypto by ID
	crypto_found, err := repo.FindCryptoById(server.conn.DB, context.Background(), int(request.Id))

	//Starts the Streaming it to the Client
	streaming.Send(&pb.ObserveCryptoResponse{
		Crypto: &pb.Crypto{
			Id:    int32(crypto_found.Id),
			Name:  crypto_found.Name,
			Token: crypto_found.Token,
			Votes: int32(crypto_found.Votes),
		},
	})

	return err
}

//Funct to Retrieve a Crypto List
func (server *CryptoServiceServer) ListCryptos(ctx context.Context, request *pb.ListCryptosRequest) (*pb.ListCryptosResponse, error) {
	//Get the Sort Parameter Result and Sets it to UPPERCASE to prevente Case-Sensitivy errors
	sortParameter := strings.ToLower(request.Sortparam)

	//Create a Crypto Array to Store the Result List
	cryptos, err := repo.FindAllCryptos(server.conn.DB, ctx, sortParameter)
	if err != nil {
		log.Println("Error while retrieving list ", err)
		return nil, err
	}

	if (sortParameter != "name" && sortParameter != "token") || (sortParameter != "votes" && sortParameter == "") {
		sortParameter = "ID"
	}

	//Print some info text
	log.Print("======> Crypto List from Database Sorted by ", strings.ToUpper(sortParameter), "\n")
	log.Println(cryptos)
	fmt.Print("\n\n")

	//Create a new instance of Crypto Array to store the List Result
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

	//Return lpopulated list
	return cryptos_list, nil
}

//Function to Start the Server
func Run() {
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
