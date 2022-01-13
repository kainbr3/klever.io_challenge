package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-playground/validator/v10"
	m "github.com/kainbr3/klever.io_challenge/src/package/model"
	t "github.com/kainbr3/klever.io_challenge/src/package/tool"
)

//Struct to store the Database connection to share it across the application
type Klever struct {
	DB *sql.DB
}

//Function to Open the Database Connection and store it in the Klever Struct
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

//Function to Build Database, Create Tables and Data Seeding
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

//Function to Add a new Cryptocurrency
func AddCrypto(database *sql.DB, ctx context.Context, crypto m.CryptoCurrency) (*m.CryptoCurrency, error) {
	//Creates a new Validator
	validate := validator.New()

	//Validate all struct fields
	err := validate.Struct(crypto)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			log.Println("======> Error ", err)
		}
		log.Print("\n\n")
		return &crypto, err
	}

	//Sets a 5 second timeout to prevent being stuck
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	//Prepare query
	query := fmt.Sprintf(t.SetDatabase+t.InsertCrypto, crypto.Name, crypto.Token, crypto.Votes)

	//Begin transaction
	transaction, err := database.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("======> Error starting new Transaction: \n%v", transaction)
		return &crypto, err
	}

	//Temporary variables to store the RowScan Result
	var name, token string
	var id, votes int

	//Execute query and scan the values from the result
	result := transaction.QueryRowContext(ctx, query).Scan(&id, &name, &token, &votes)
	if result != nil {
		transaction.Rollback()
		log.Printf("======> Error while executing statement: \n%s", query)
		return &crypto, err
	}

	//Creates an instance of CryptoCurrency to return the Crypto Added with all Properties fulfilled
	var crypto_created = m.CryptoCurrency{
		Id:    id,
		Name:  name,
		Token: token,
		Votes: votes,
	}

	//Commit the transaction
	err = transaction.Commit()
	if err != nil {
		transaction.Rollback()
		log.Printf("======> Error while commiting changes to database: \n%v", err)
		return &crypto, err
	}

	//Return the new Crypto
	return &crypto_created, nil
}

func UpdateCrypto(database *sql.DB, ctx context.Context, crypto m.CryptoCurrency) (*m.CryptoCurrency, error) {
	//Creates a new Validator
	validate := validator.New()

	//Validate all struct fields
	err := validate.Struct(crypto)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			log.Println("======> Error ", err)
		}
		log.Print("\n\n")
		return nil, err
	}

	//Sets a 5 second timeout to prevent being stuck
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	//Prepare query
	query := fmt.Sprintf(t.SetDatabase+t.UpdateCryptoQuery, crypto.Name, crypto.Token, crypto.Votes, crypto.Id)

	//Begin transaction
	transaction, err := database.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("======> Error starting new Transaction")
		return nil, err
	}

	//Temporary variables to store the RowScan Result
	var name, token string
	var id, votes int

	//Execute query and scan the values from the result
	result := transaction.QueryRowContext(ctx, query).Scan(&id, &name, &token, &votes)
	if result != nil {
		transaction.Rollback()
		log.Printf("======> Error while executing statement: \n%s", query)
		return nil, err
	}

	//Commit the transaction
	err = transaction.Commit()
	if err != nil {
		transaction.Rollback()
		log.Printf("======> Error while commiting changes to database: \n%v", err)
		return nil, err
	}

	//Creates an instance of CryptoCurrency to return the Crypto Added with all Properties fulfilled
	var crypto_updated = m.CryptoCurrency{
		Id:    id,
		Name:  name,
		Token: token,
		Votes: votes,
	}

	//Return the updated Crypto
	return &crypto_updated, nil
}

// //Function to return a CRYPTO by ID
func FindCryptoById(database *sql.DB, ctx context.Context, cryptoId int) (m.CryptoCurrency, error) {
	//Creates an instance of CryptoCurrency to return the Crypto Found with all Properties fulfilled
	crypto := m.CryptoCurrency{}

	//Validate the ID Parameter
	if cryptoId == 0 {
		return crypto, errors.New("invalid value for Id")
	}

	//Sets a 5 second timeout to prevent being stuck
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	//Prepare query
	query := fmt.Sprintf(t.SetDatabase+t.SelectCryptoByIdQuery, cryptoId)

	//Begin transaction
	transaction, err := database.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("======> Error starting new Transaction")
		return crypto, err
	}

	//Temporary variables to store the RowScan Result
	var name, token string
	var id, votes int

	//Execute query and scan the values from the result
	result := transaction.QueryRowContext(ctx, query).Scan(&id, &name, &token, &votes)
	if result != nil {
		transaction.Rollback()
		log.Printf("======> Error while executing statement: \n%s", query)
		return crypto, err
	}

	//Commit the transaction
	err = transaction.Commit()
	if err != nil {
		transaction.Rollback()
		log.Printf("======> Error while commiting changes to database: \n%v", err)
		return crypto, err
	}

	//Creates an instance of CryptoCurrency to return the Crypto Added with all Properties fulfilled
	crypto = m.CryptoCurrency{
		Id:    id,
		Name:  name,
		Token: token,
		Votes: votes,
	}

	return crypto, nil
}

//Function to Add a Vote to Given Crypto
func UpvoteCryptoById(database *sql.DB, ctx context.Context, cryptoId int) error {
	//Validate the ID Parameter
	if cryptoId == 0 {
		return errors.New("invalid value for Id")
	}

	//Sets a 5 second timeout to prevent being stuck
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	//Prepare query
	query := fmt.Sprintf(t.SetDatabase+t.UpvoteCryptoQuery, cryptoId)

	//Begin transaction
	transaction, err := database.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("======> Error starting new Transaction")
		return err
	}

	//Execute query and scan the values from the result
	result := transaction.QueryRowContext(ctx, query)
	if result != nil {
		transaction.Rollback()
		log.Printf("======> Error while executing statement: \n%s", query)
		return err
	}

	//Commit the transaction
	err = transaction.Commit()
	if err != nil {
		transaction.Rollback()
		log.Printf("======> Error while commiting changes to database: \n%v", err)
		return err
	}

	return nil
}

//Function to Subtract a Vote to Given Crypto
func DownvoteCryptoById(database *sql.DB, ctx context.Context, cryptoId int) error {
	//Validate the ID Parameter
	if cryptoId == 0 {
		return errors.New("invalid value for Id")
	}

	//Sets a 5 second timeout to prevent being stuck
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	//Prepare query
	query := fmt.Sprintf(t.SetDatabase+t.DownvoteCryptoQuery, cryptoId)

	//Begin transaction
	transaction, err := database.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("======> Error starting new Transaction")
		return err
	}

	//Execute query and scan the values from the result
	result := transaction.QueryRowContext(ctx, query)
	if result != nil {
		transaction.Rollback()
		log.Printf("======> Error while executing statement: \n%s", query)
		return err
	}

	//Commit the transaction
	err = transaction.Commit()
	if err != nil {
		transaction.Rollback()
		log.Printf("======> Error while commiting changes to database: \n%v", err)
		return err
	}

	return nil
}

//Function to Delete a Crypto by its ID
func RemoveCryptoById(database *sql.DB, ctx context.Context, cryptoId int) error {
	//Validate the ID Parameter
	if cryptoId == 0 {
		return errors.New("invalid value for Id")
	}

	//Sets a 5 second timeout to prevent being stuck
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	//Prepare query
	query := fmt.Sprintf(t.SetDatabase+t.DeleteCryptoById, cryptoId)

	//Begin transaction
	transaction, err := database.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("======> Error starting new Transaction")
		return err
	}

	var idDeleted int
	//Execute query and scan the values from the result
	result := transaction.QueryRowContext(ctx, query)
	err = result.Scan(&idDeleted)
	if err != nil {
		transaction.Rollback()
		log.Printf("======> Error while executing statement: \n%s", query)
		return err
	}

	if idDeleted != cryptoId {
		transaction.Rollback()
		var messageError = fmt.Sprintf("======> Error while deleting Crypto ID %d", cryptoId)
		err = errors.New(messageError)
		log.Println(messageError)
		return err
	}

	//Commit the transaction
	err = transaction.Commit()
	if err != nil {
		transaction.Rollback()
		log.Printf("======> Error while commiting changes to database: \n%v", err)
		return err
	}

	return nil
}

//Function to Get All Cryptos according to given parameter
func FindAllCryptos(database *sql.DB, ctx context.Context, sortParameter string) ([]m.CryptoCurrency, error) {
	//Creates a new CRYPTOCURRENCIES arrray to store the results from database query
	cryptos := []m.CryptoCurrency{}

	//Query variable to store the specific query according to Parameter
	var query string

	//Sets a 5 second timeout to prevent being stuck
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	//Check the Parameter to Sort the Result
	switch sortParameter {
	case "name":
		query = t.SetDatabase + t.SelectAllCryptosSortedByNameQuery
	case "token":
		query = t.SetDatabase + t.SelectAllCryptosSortedByTokenQuery
	case "votes":
		query = t.SetDatabase + t.SelectAllCryptosSortedByTopVotesQuery
	default:
		query = t.SetDatabase + t.SelectAllCryptosQuery
	}

	//Begin transaction
	transaction, err := database.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("======> Error starting new Transaction")
		return nil, err
	}

	//Execute query and scan the values from the result
	rows, err := transaction.QueryContext(ctx, query)
	if err != nil {
		transaction.Rollback()
		log.Printf("======> Error while executing statement: \n%s", query)
		return nil, err
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

	//Commit the transaction
	err = transaction.Commit()
	if err != nil {
		transaction.Rollback()
		log.Printf("======> Error while commiting changes to database: \n%v", err)
		return nil, err
	}

	return cryptos, nil
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
