package main

import (
	"context"
	"fmt"
	"log"
	"time"

	t "github.com/kainbr3/klever.io_challenge/src/package/tool"
	pb "github.com/kainbr3/klever.io_challenge/src/protobuf"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Print("\n\n")
	log.Println("======> STARTING THE gRPC CLIENT")
	//Starts the connection with gRPC Server
	connection, err := grpc.Dial(t.ClientAdress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("======> Could not connect: %v", err)
	}
	defer connection.Close() //Closes the connection at the end of the function scope

	//Creates a new instance of gRPC Client
	client := pb.NewCryptoServiceClient(connection)

	//Get the Context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// //Make the Request and store the Response
	responseListCryptos, err := client.ListCryptos(ctx, &pb.ListCryptosRequest{
		Sortparam: "name",
	})
	if err != nil {
		log.Fatalf("======> Could not retrieve cryptos: \n%v", err)
	}
	log.Print("======> CRYPTO LIST SORTED BY NAME: \n")
	log.Printf("======> request.GetCryptos(): %v", responseListCryptos.Cryptos)
	fmt.Print("\n\n")

	//Make the Request and store the Response
	responseNewCrypto, err := client.CreateNewCrypto(ctx, &pb.CreateNewCryptoRequest{
		Name:  "2021 Fim do Mundo",
		Token: "21F",
	})
	if err != nil {
		log.Fatalf("======> Could not create new crypto: %v", err)
	}

	//Log in the console the Response Values of the new CryptoCurrency Created
	log.Printf(`
======> Received from server/database:
======> Crypto Created
======> ID: %d
======> NAME: %s
======> TOKEN: %s
======> VOTES: %d
	`,
		responseNewCrypto.GetCrypto().Id,
		responseNewCrypto.Crypto.GetName(),
		responseNewCrypto.Crypto.GetToken(),
		responseNewCrypto.Crypto.GetVotes(),
	)

	//Make the Request and store the Response
	responseNewCrypto.GetCrypto().Name = responseNewCrypto.GetCrypto().GetName() + " UPDATED!"
	responseUpdateCrypto, err := client.UpdateCrypto(ctx, &pb.UpdateCryptoRequest{
		Crypto: responseNewCrypto.GetCrypto(),
	})
	if err != nil {
		log.Fatalf("======> Could not update crypto: %v", err)
	}

	//Log in the console the Response Values of the CryptoCurrency Updated
	log.Printf(`
======> Received from server/database:
======> Crypto Updated
======> ID: %d
======> NAME: %s
======> TOKEN: %s
======> VOTES: %d
	`,
		responseUpdateCrypto.GetCrypto().Id,
		responseUpdateCrypto.Crypto.GetName(),
		responseUpdateCrypto.Crypto.GetToken(),
		responseUpdateCrypto.Crypto.GetVotes(),
	)

	//Log in the console the Response Values of the CryptoCurrency Before Updated Votes
	log.Printf(`
======> Votes Before Update:
======> NAME: %s
======> VOTES: %d
	`,
		responseUpdateCrypto.Crypto.GetName(),
		responseUpdateCrypto.Crypto.GetVotes(),
	)
	//Make the Request and store the Response
	_, err = client.UpvoteCrypto(ctx, &pb.UpvoteCryptoRequest{
		Id: responseUpdateCrypto.GetCrypto().GetId(),
	})
	if err != nil {
		log.Fatalf("======> Could not add vote to crypto: %v", err)
	}

	//Log in the console the Response Values of the CryptoCurrency Votes Updated
	responseGetById, err := client.GetCryptoById(ctx, &pb.GetCryptoByIdRequest{
		Id: responseUpdateCrypto.Crypto.GetId(),
	})
	if err != nil {
		log.Fatalf("======> Could not retrieve crypto: %v", err)
	}

	log.Printf(`
======> Received from server/database:
======> Crypto Upvote Updated
======> NAME: %s
======> VOTES: %d
	`,
		responseGetById.Crypto.GetName(),
		responseGetById.Crypto.GetVotes(),
	)

	//Make the Request and store the Response
	_, err = client.DownvoteCrypto(ctx, &pb.DownvoteCryptoRequest{
		Id: responseUpdateCrypto.GetCrypto().GetId(),
	})
	if err != nil {
		log.Fatalf("======> Could not add vote to crypto: %v", err)
	}

	//Log in the console the Response Values of the CryptoCurrency Votes Updated
	responseGetById, err = client.GetCryptoById(ctx, &pb.GetCryptoByIdRequest{
		Id: responseUpdateCrypto.Crypto.GetId(),
	})
	if err != nil {
		log.Fatalf("======> Could not retrieve crypto: %v", err)
	}

	log.Printf(`
======> Received from server/database:
======> Crypto Downvote Updated
======> NAME: %s
======> VOTES: %d
`,
		responseGetById.Crypto.GetName(),
		responseGetById.Crypto.GetVotes(),
	)
	fmt.Print("\n\n")

	_, err = client.DeleteCrypto(ctx, &pb.DeleteCryptoRequest{
		Id: responseGetById.Crypto.GetId(),
	})
	if err != nil {
		log.Fatalf("======> Could not delete crypto: %v", err)
	}
	log.Print("Crypto Deleted Id: ", responseGetById.Crypto.GetId())
	fmt.Print("\n\n")
}
