package main

import (
	"context"
	"fmt"
	"log"
	"time"

	t "github.com/kainbr3/klever.io_challenge/package/tool"
	pb "github.com/kainbr3/klever.io_challenge/protobuf"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	//Starts the connection with gRPC Server
	connection, err := grpc.Dial(t.ClientAdress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("======> Could not connect: %v", err)
	}
	defer connection.Close() //Closes the connection at the end of the function scope

	//Creates a new instance of gRPC Client
	client := pb.NewCryptoServiceClient(connection)

	//Get the Context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//Make the Request and store the Response
	response, err := client.CreateNewCrypto(ctx, &pb.NewCryptoRequest{
		Name:  "2021 Fim do Mundo",
		Token: "21F",
	})
	if err != nil {
		log.Fatalf("======> Could not create new crypto: %v", err)
	}

	//Log in the console the Response Values of the new CryptoCurrency Created
	// 	log.Printf(`======> Received from server:
	// ======> Crypto from a Database
	// ======> ID: %d
	// ======> NAME: %s
	// ======> TOKEN: %s
	// ======> VOTES: %d
	// 	`,
	// 		response.GetCrypto().Id,
	// 		response.Crypto.GetName(),
	// 		response.Crypto.GetToken(),
	// 		response.Crypto.GetVotes(),
	// 	)

	log.Print("OLHAR AQUI ***********************************")
	fmt.Print(response)
	// fmt.Print(response.Crypto)
	// fmt.Print(response.Crypto.GetName())
	// fmt.Print(response.Crypto.Name)
	// fmt.Print(response.GetCrypto())
	// fmt.Print(response.GetCrypto().Name)
	// fmt.Print(response.GetCrypto().GetName())
	//log.Print("**********************************************")

	// params := &pb.ListCryptosRequest{}
	// response2, err := client.GetCryptos(ctx, params)
	// if err != nil {
	// 	log.Fatalf("======> Could not retrieve cryptos: \n%v", err)
	// }
	// log.Print("\n======> CRYPTO LIST: \n")
	// fmt.Printf("======> request.GetCryptos(): %v\n", response2.Cryptos)
}
