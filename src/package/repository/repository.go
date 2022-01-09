package repository

import (
	"database/sql"
	"fmt"
	"log"

	m "github.com/kainbr3/klever.io_challenge/package/model"
	t "github.com/kainbr3/klever.io_challenge/package/tool"
	_ "github.com/mattn/go-sqlite3" //Manually force Import SQLite3 Package to use its driver
)

type Klever struct {
	DB *sql.DB
}

func DatabaseInit() *Klever {
	//Display some info
	log.Println("======> Driver Type: ", t.Driver)
	log.Println("======> Database Name: ", t.Database)
	fmt.Print("\n\n")

	//Open database connection
	database, err := sql.Open(t.Driver, t.Database)

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
	rows, err := database.Query(t.SelectAllCryptosQuery)

	if err != nil {
		log.Println("======> Error while executing query: \n", t.SelectAllCryptosQuery, "\n======> ERRROR: ", err)
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
		log.Println("======> Error while executing query: \n", t.SelectAllCryptosSortedByNameQuery, "\n======> ERRROR: ", err)
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
		log.Println("======> Error while executing query: \n", t.SelectAllCryptosSortedByTokenQuery, "\n======> ERRROR: ", err)
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
		log.Println("======> Error while executing query: \n", t.SelectAllCryptosSortedByLeastVotesQuery, "\n======> ERRROR: ", err)
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
		log.Println("======> Error while executing query: \n", t.SelectAllCryptosSortedByTopVotesQuery, "\n======> ERRROR: ", err)
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
		log.Println("======> Error while executing query: \n", t.SelectCryptoByIdQuery, cryptoId, "\n======> ERRROR: ", err)
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
		log.Println("======> Error while executing query: \n", t.SelectCryptoByNameQuery, cryptoName, "\n======> ERRROR: ", err)
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
		log.Println("======> Error while executing query: \n", t.SelectCryptoByTokenQuery, cryptoToken, "\n======> ERRROR: ", err)
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

func AddCrypto(database *sql.DB, crypto m.CryptoCurrency) *m.CryptoCurrency {
	//Prepare the statement
	stm, err := database.Prepare(t.InsertCrypto)
	if err != nil {
		log.Printf("======> Could not prepare statement: \n%v", err)
	}

	//Execute query
	//stm.Exec(crypto.Name, crypto.Token, time.Now().Second()) //Random votes generated for test purposes
	result, _ := stm.Exec(crypto.Name, crypto.Token, crypto.Votes)

	//Get the if from the record
	id, _ := result.LastInsertId()

	//Creates a instance of CryptoCurrency to return the Crypto Added with all Properties filled
	var crypto_created = FindCryptoById(database, int(id))

	//Retorna a Crypto
	return &crypto_created
}

func RemoveCryptoById(database *sql.DB, cryptoId int) {
	//Prepare the statement
	stm, err := database.Prepare(t.DeleteCryptoById)

	if err != nil {
		log.Printf("======> Could not prepare statement: \n%v", err)
	}

	//Execute query
	stm.Exec(cryptoId)
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
