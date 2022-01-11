package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/denisenkom/go-mssqldb" //Manually force Import SQLite3 Package to use its driver
	m "github.com/kainbr3/klever.io_challenge/src/package/model"
	t "github.com/kainbr3/klever.io_challenge/src/package/tool"
	//_ "github.com/mattn/go-sqlite3" //Manually force Import SQLite3 Package to use its driver
)

type Klever struct {
	DB *sql.DB
}

func DatabaseInit() *Klever {
	//Display some info
	log.Println("======> Driver Type: ", t.Driver)
	log.Println("======> Database Name: ", t.ConnectionString)
	fmt.Print("\n\n")

	//Open database connection
	database, err := sql.Open(t.Driver, t.ConnectionString)

	//Test if the Database file is accessible
	if err != nil {
		log.Fatalf("======> Could not connect to the database \n%v", err)
	}
	//show some info
	log.Println("======> Database Found. Ready to connect!")
	fmt.Print("\n\n")

	return &Klever{
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
	queries = append(queries, t.SetDatabase+t.CreateDatabaseQuery, t.SetDatabase+t.DropCryptoTable, t.SetDatabase+t.CreateCryptoTableQuery)

	//Append the Data Seed Query. Use it to populate the empty database setting true as args to BuildDatabase Function
	if seedData {
		queries = append(queries, t.SetDatabase+t.SeedCryptoDataQuery)
	}

	//Ignore the index and use only values from the range iteration
	for _, value := range queries {
		//Execute each SQL Query in the string array
		statement, err = database.Prepare(value)

		if err != nil {
			log.Fatalf("======> Could not prepare statement: \n%v \n %v", value, err)
		}

		//Execute query
		statement.Exec()
	}
}

func FindAllCryptos(database *sql.DB) []m.CryptoCurrency {
	//Creates a new CRYPTOCURRENCIES arrray to store the results from database query
	cryptos := []m.CryptoCurrency{}

	//Execute query
	query := t.SetDatabase + t.SelectAllCryptosQuery
	rows, err := database.Query(query)

	if err != nil {
		log.Println("======> Error while executing query: \n", query, "\n======> ERRROR: ", err)
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

func FindAllCryptosSortedByName(database *sql.DB, ctx context.Context) []m.CryptoCurrency {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	//Creates a new CRYPTOCURRENCIES arrray to store the results from database query
	cryptos := []m.CryptoCurrency{}

	//Execute query
	query := t.SetDatabase + t.SelectAllCryptosSortedByNameQuery
	rows, err := database.QueryContext(ctx, query)

	if err != nil {
		log.Println("======> Error while executing query: \n", query, "\n======> ERRROR: ", err)
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
	query := t.SetDatabase + t.SelectAllCryptosSortedByTokenQuery
	rows, err := database.Query(query)

	if err != nil {
		log.Println("======> Error while executing query: \n", query, "\n======> ERRROR: ", err)
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
	query := t.SetDatabase + t.SelectAllCryptosSortedByLeastVotesQuery
	rows, err := database.Query(query)

	if err != nil {
		log.Println("======> Error while executing query: \n", query, "\n======> ERRROR: ", err)
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
	query := t.SetDatabase + t.SelectAllCryptosSortedByTopVotesQuery
	rows, err := database.Query(query)

	if err != nil {
		log.Println("======> Error while executing query: \n", query, "\n======> ERRROR: ", err)
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
	query := fmt.Sprintf(t.SetDatabase+t.SelectCryptoByIdQuery, cryptoId)
	rows, err := database.Query(query)

	if err != nil {
		log.Println("======> Error while executing query: \n", query, "\n======> ERRROR: ", err)
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
	query := fmt.Sprintf(t.SelectCryptoByNameQuery, cryptoName)
	rows, err := database.Query(query)

	if err != nil {
		log.Println("======> Error while executing query: \n", query, "\n======> ERRROR: ", err)
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
	query := fmt.Sprintf(t.SelectCryptoByTokenQuery, cryptoToken)
	rows, err := database.Query(query)

	if err != nil {
		log.Println("======> Error while executing query: \n", query, "\n======> ERRROR: ", err)
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

func AddCrypto(database *sql.DB, ctx context.Context, crypto m.CryptoCurrency) *m.CryptoCurrency {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	transaction, err := database.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("======> Error starting new Transaction")
	}

	test := fmt.Sprintf(t.SetDatabase+t.InsertCrypto, crypto.Name, crypto.Token, crypto.Votes)
	log.Println(test)

	//Prepare the statement
	result, err := transaction.ExecContext(ctx, fmt.Sprintf(t.SetDatabase+t.InsertCrypto, crypto.Name, crypto.Token, crypto.Votes))
	if err != nil {
		transaction.Rollback()
		log.Printf("======> Could not prepare statement: %v", err)
	}

	log.Print(result)

	//Execute query
	err = transaction.Commit()
	if err != nil {
		log.Printf("======> Error while commiting changes to database %v", err)
	}

	//Get the if from the record
	id, _ := result.LastInsertId()

	//Creates a instance of CryptoCurrency to return the Crypto Added with all Properties filled
	var crypto_created = FindCryptoById(database, int(id))

	//Retorna a Crypto
	return &crypto_created
}

func RemoveCryptoById(database *sql.DB, cryptoId int) {
	//Prepare the statement
	stm, err := database.Prepare(t.SetDatabase + t.DeleteCryptoById)

	if err != nil {
		log.Printf("======> Could not prepare statement: \n%v", err)
	}

	//Execute query
	stm.Exec(cryptoId)
}

func UpdateCrypto(database *sql.DB, crypto m.CryptoCurrency) *m.CryptoCurrency {
	//Prepare the statement
	stm, err := database.Prepare(t.UpdateCryptoQuery)
	if err != nil {
		log.Printf("======> Could not prepare statement: \n%v", err)
	}

	//Execute query
	stm.Exec(crypto.Name, crypto.Token, crypto.Votes, crypto.Id)

	//Creates a instance of CryptoCurrency to return the Crypto Added with all Properties filled
	var crypto_updated = FindCryptoById(database, crypto.Id)

	//Retorna a Crypto
	return &crypto_updated
}

func UpvoteCryptoById(database *sql.DB, cryptoId int) {
	//Prepare the statement
	stm, err := database.Prepare(t.UpvoteCryptoQuery)

	if err != nil {
		log.Printf("======> Could not prepare statement: \n%v", err)
	}

	//Execute query
	stm.Exec(cryptoId)
}

func DownvoteCryptoById(database *sql.DB, cryptoId int) {
	//Prepare the statement
	stm, err := database.Prepare(t.DownvoteCryptoQuery)

	if err != nil {
		log.Printf("======> Could not prepare statement: \n%v", err)
	}

	//Execute query
	stm.Exec(cryptoId)
}

func ReseedTable(database *sql.DB, tableName string) {
	//Prepare the statement
	stm, err := database.Prepare(fmt.Sprintf(t.ReseedTable, tableName, tableName))

	if err != nil {
		log.Printf("======> Could not prepare statement: \n%v", err)
	}

	//Execute query
	stm.Exec()
}

func CleanTable(database *sql.DB, tableName string) {
	//Prepare the statement
	stm, err := database.Prepare(fmt.Sprintf(t.DeleteAllFromTable, tableName))

	if err != nil {
		log.Printf("======> Could not prepare statement: \n%v", err)
	}

	//Execute query
	stm.Exec()
}

func DropTable(database *sql.DB, tableName string) {
	//Prepare the statement
	stm, err := database.Prepare(fmt.Sprintf(t.DropTable, tableName))

	if err != nil {
		log.Printf("======> Could not prepare statement: \n%v", err)
	}

	//Execute query
	stm.Exec()
}

func ExecuteCustomStatement(database *sql.DB, statement string) {
	//Prepare the statement
	stm, err := database.Prepare(statement)

	if err != nil {
		log.Printf("======> Could not prepare statement: \n%v", err)
	}

	//Execute query
	stm.Exec()
}
