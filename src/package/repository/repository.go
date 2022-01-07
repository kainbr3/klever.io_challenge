package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	m "github.com/kainbr3/klever.io_challenge/package/model"
	t "github.com/kainbr3/klever.io_challenge/package/tool"
	_ "github.com/mattn/go-sqlite3" //Manually force Import SQLite3 Package to use its driver
)

type KleverDB struct {
	DB *sql.DB
}

func DatabaseInit() *KleverDB {
	//Display some info
	fmt.Println("======> Driver Type: ", t.Driver)
	fmt.Println("======> Database Name: ", t.Database)

	//Open database connection
	database, err := sql.Open(t.Driver, t.Database)

	//Test if the Database file is accessible
	if err != nil {
		log.Fatalf("======> Could not connect to the database %v", err)
	}
	//show some info
	fmt.Println("======> Database Found. Ready to connect!")

	return &KleverDB{
		DB: database,
	}
}

func BuildDatabase(database *sql.DB, seedData bool) {
	//Separeted statement and error variables to be reused, as GO cannot run multiples CreateTable Commands in the same query
	var statement *sql.Stmt
	var err error

	//Variable to store all the SQL Queries that will be used
	var queries []string

	//Concact all the queries that will be executed
	queries = append(queries, t.DropCryptoTable, t.CreateCryptoTableQuery)

	//Append the Data Seed Query. Use it to populate the empty database setting true as args to BuildDatabase Function
	if seedData {
		queries = append(queries, t.SeedCryptoDataQuery)
	}

	//Ignore the index and use only values from the range iteration
	for _, value := range queries {
		//Execute each SQL Query in the string array
		statement, err = database.Prepare(value)

		if err != nil {
			log.Fatalf("Could not prepare statement: %v \n %v", value, err)
		}

		//Execute query
		statement.Exec()
	}
}

func FindAllCryptos(database *sql.DB) []m.CryptoCurrency {
	//Creates a new CRYPTOCURRENCIES arrray to store the results from database query
	cryptos := []m.CryptoCurrency{}

	//Execute query
	rows, err := database.Query(t.SelectAllCryptosQuery)

	if err != nil {
		fmt.Println("Error while executing query: ", t.SelectAllCryptosQuery, " ERRROR: ", err)
	}

	//Variables to be used durring the rresult iteration
	var id int
	var name string
	var token string
	var votes int

	//Datarows iteration to create the object list
	for rows.Next() {
		rows.Scan(&id, &name, &token, &votes)
		cryptos = append(cryptos, m.CryptoCurrency{
			Id:    id,
			Name:  name,
			Token: token,
			Votes: votes,
		})
	}

	return cryptos
}

func FindAllCryptosSortedByName(database *sql.DB) []m.CryptoCurrency {
	//Creates a new CRYPTOCURRENCIES arrray to store the results from database query
	cryptos := []m.CryptoCurrency{}

	//Execute query
	rows, err := database.Query(t.SelectAllCryptosSortedByNameQuery)

	if err != nil {
		fmt.Println("Error while executing query: ", t.SelectAllCryptosSortedByNameQuery, " ERRROR: ", err)
	}

	//Variables to be used durring the rresult iteration
	var id int
	var name string
	var token string
	var votes int

	//Datarows iteration to create the object list
	for rows.Next() {
		rows.Scan(&id, &name, &token, &votes)
		cryptos = append(cryptos, m.CryptoCurrency{
			Id:    id,
			Name:  name,
			Token: token,
			Votes: votes,
		})
	}

	return cryptos
}

func FindAllCryptosSortedByToken(database *sql.DB) []m.CryptoCurrency {
	//Creates a new CRYPTOCURRENCIES arrray to store the results from database query
	cryptos := []m.CryptoCurrency{}

	//Execute query
	rows, err := database.Query(t.SelectAllCryptosSortedByTokenQuery)

	if err != nil {
		fmt.Println("Error while executing query: ", t.SelectAllCryptosSortedByTokenQuery, " ERRROR: ", err)
	}

	//Variables to be used durring the rresult iteration
	var id int
	var name string
	var token string
	var votes int

	//Datarows iteration to create the object list
	for rows.Next() {
		rows.Scan(&id, &name, &token, &votes)
		cryptos = append(cryptos, m.CryptoCurrency{
			Id:    id,
			Name:  name,
			Token: token,
			Votes: votes,
		})
	}

	return cryptos
}

func FindAllCryptosSortedByLeastVotes(database *sql.DB) []m.CryptoCurrency {
	//Creates a new CRYPTOCURRENCIES arrray to store the results from database query
	cryptos := []m.CryptoCurrency{}

	//Execute query
	rows, err := database.Query(t.SelectAllCryptosSortedByLeastVotesQuery)

	if err != nil {
		fmt.Println("Error while executing query: ", t.SelectAllCryptosSortedByLeastVotesQuery, " ERRROR: ", err)
	}

	//Variables to be used durring the rresult iteration
	var id int
	var name string
	var token string
	var votes int

	//Datarows iteration to create the object list
	for rows.Next() {
		rows.Scan(&id, &name, &token, &votes)
		cryptos = append(cryptos, m.CryptoCurrency{
			Id:    id,
			Name:  name,
			Token: token,
			Votes: votes,
		})
	}

	return cryptos
}

func FindAllCryptosSortedByTopVotes(database *sql.DB) []m.CryptoCurrency {
	//Creates a new CRYPTOCURRENCIES arrray to store the results from database query
	cryptos := []m.CryptoCurrency{}

	//Execute query
	rows, err := database.Query(t.SelectAllCryptosSortedByTopVotesQuery)

	if err != nil {
		fmt.Println("Error while executing query: ", t.SelectAllCryptosSortedByTopVotesQuery, " ERRROR: ", err)
	}

	//Variables to be used durring the rresult iteration
	var id int
	var name string
	var token string
	var votes int

	//Datarows iteration to create the object list
	for rows.Next() {
		rows.Scan(&id, &name, &token, &votes)
		cryptos = append(cryptos, m.CryptoCurrency{
			Id:    id,
			Name:  name,
			Token: token,
			Votes: votes,
		})
	}

	return cryptos
}

//Function to return a CRYPTO by ID
func FindCryptoById(database *sql.DB, cryptoId int) m.CryptoCurrency {
	//Creates a new CRYPTOCURRENCIES arrray to store the results from database query
	crypto := m.CryptoCurrency{}

	//Execute query
	rows, err := database.Query(t.SelectCryptoByIdQuery, cryptoId)

	if err != nil {
		fmt.Println("Error while executing query: ", t.SelectCryptoByIdQuery, cryptoId, " ERRROR: ", err)
	}

	//Variables to be used durring the rresult iteration
	var id int
	var name string
	var token string
	var votes int

	//Datarows iteration to create the object list
	for rows.Next() {
		rows.Scan(&id, &name, &token, &votes)
		crypto = m.CryptoCurrency{
			Id:    id,
			Name:  name,
			Token: token,
			Votes: votes,
		}
	}

	return crypto
}

//Function to return a CRYPTO by NAME. PS: Its Case Sensitive
func FindCryptoByName(database *sql.DB, cryptoName string) m.CryptoCurrency {
	//Creates a new CRYPTOCURRENCIES arrray to store the results from database query
	crypto := m.CryptoCurrency{}

	//Execute query
	rows, err := database.Query(fmt.Sprintf(t.SelectCryptoByNameQuery, cryptoName))

	if err != nil {
		fmt.Println("Error while executing query: ", t.SelectCryptoByNameQuery, cryptoName, " ERRROR: ", err)
	}

	//Variables to be used durring the rresult iteration
	var id int
	var name string
	var token string
	var votes int

	//Datarows iteration to create the object list
	for rows.Next() {
		rows.Scan(&id, &name, &token, &votes)
		crypto = m.CryptoCurrency{
			Id:    id,
			Name:  name,
			Token: token,
			Votes: votes,
		}
	}

	return crypto
}

//Function to return a CRYPTO by TOKEN. PS: Its Case Sensitive
func FindCryptoByToken(database *sql.DB, cryptoToken string) m.CryptoCurrency {
	//Creates a new CRYPTOCURRENCIES arrray to store the results from database query
	crypto := m.CryptoCurrency{}

	//Execute query
	rows, err := database.Query(fmt.Sprintf(t.SelectCryptoByTokenQuery, cryptoToken))

	if err != nil {
		fmt.Println("Error while executing query: ", t.SelectCryptoByTokenQuery, cryptoToken, " ERRROR: ", err)
	}

	//Variables to be used durring the rresult iteration
	var id int
	var name string
	var token string
	var votes int

	//Datarows iteration to create the object list
	for rows.Next() {
		rows.Scan(&id, &name, &token, &votes)
		crypto = m.CryptoCurrency{
			Id:    id,
			Name:  name,
			Token: token,
			Votes: votes,
		}
	}

	return crypto
}

func AddCrypto(database *sql.DB, crypto m.CryptoCurrency) {
	//Prepare the statement
	stm, err := database.Prepare(t.InsertCrypto)

	if err != nil {
		fmt.Printf("Could not prepare statement: %v", err)
	}

	//Execute query
	// stm.Exec(crypto.Name, crypto.Token, rand.Intn(99))
	stm.Exec(crypto.Name, crypto.Token, time.Now().Second())
}

func RemoveCryptoById(database *sql.DB, cryptoId int) {
	//Prepare the statement
	stm, err := database.Prepare(t.DeleteCryptoById)

	if err != nil {
		fmt.Printf("Could not prepare statement: %v", err)
	}

	//Execute query
	stm.Exec(cryptoId)
}

func UpvoteCryptoById(database *sql.DB, cryptoId int) {
	//Prepare the statement
	stm, err := database.Prepare(t.UpvoteCryptoQuery)

	if err != nil {
		fmt.Printf("Could not prepare statement: %v", err)
	}

	//Execute query
	stm.Exec(cryptoId)
}

func DownvoteCryptoById(database *sql.DB, cryptoId int) {
	//Prepare the statement
	stm, err := database.Prepare(t.DownvoteCryptoQuery)

	if err != nil {
		fmt.Printf("Could not prepare statement: %v", err)
	}

	//Execute query
	stm.Exec(cryptoId)
}

func ReseedTable(database *sql.DB, tableName string) {
	//Prepare the statement
	stm, err := database.Prepare(fmt.Sprintf(t.ReseedTable, tableName, tableName))

	if err != nil {
		fmt.Printf("Could not prepare statement: %v", err)
	}

	//Execute query
	stm.Exec()
}

func CleanTable(database *sql.DB, tableName string) {
	//Prepare the statement
	stm, err := database.Prepare(fmt.Sprintf(t.DeleteAllFromTable, tableName))

	if err != nil {
		fmt.Printf("Could not prepare statement: %v", err)
	}

	//Execute query
	stm.Exec()
}

func DropTable(database *sql.DB, tableName string) {
	//Prepare the statement
	stm, err := database.Prepare(fmt.Sprintf(t.DropTable, tableName))

	if err != nil {
		fmt.Printf("Could not prepare statement: %v", err)
	}

	//Execute query
	stm.Exec()
}

func ExecuteCustomStatement(database *sql.DB, statement string) {
	//Prepare the statement
	stm, err := database.Prepare(statement)

	if err != nil {
		fmt.Printf("Could not prepare statement: %v", err)
	}

	//Execute query
	stm.Exec()
}
