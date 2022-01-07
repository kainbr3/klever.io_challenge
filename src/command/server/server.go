package main

import (
	"fmt"

	repo "github.com/kainbr3/klever.io_challenge/package/repository"
)

func main() {
	//Display some info
	fmt.Println("======> The GO gRPC SERVER is UP!")

	//Connect to the Database
	klever := repo.DatabaseInit()
	fmt.Println("======> Databased Ready: ", klever.DB.Stats())

	//Function to Create the main Tables and Populate it with some Data
	repo.BuildDatabase(klever.DB, true)

	//Some Query Tests
	// fmt.Println("======> Query Result:", repo.FindAllCryptos(klever.DB))
	// fmt.Println("======> Query Result:", repo.FindAllCryptosSortedByName(klever.DB))
	// fmt.Println("======> Query Result:", repo.FindAllCryptosSortedByToken(klever.DB))
	// fmt.Println("======> Query Result:\n\n", repo.FindAllCryptosSortedByTopVotes(klever.DB))
	// fmt.Println("======> Query Result:\n\n", repo.FindAllCryptosSortedByLeastVotes(klever.DB))
	// fmt.Println("======> Query Result:", repo.FindCryptoById(klever.DB, 1))
	// fmt.Println("======> Query Result:", repo.FindCryptoByName(klever.DB, "Klever"))
	// fmt.Println("======> Query Result:", repo.FindCryptoByToken(klever.DB, "BTC"))

	//repo.AddCrypto(klever.DB, model.CryptoCurrency{Name: "Teste Crypto", Token: "TST"})
	//repo.RemoveCryptoById(klever.DB, 8)
	//repo.CleanTable(klever.DB, "cryptoCurrencies")
	//repo.ReseedTable(klever.DB, "cryptoCurrencies")
	//repo.DropTable(klever.DB, "cryptoCurrencies")
	//repo.ExecuteCustomStatement(klever.DB, tool.DropAllTables)
	//repo.ExecuteCustomStatement(klever.DB, tool.RatingTableQuery)
	// repo.UpvoteCryptoById(klever.DB, 7)
	// repo.DownvoteCryptoById(klever.DB, 3)
}
