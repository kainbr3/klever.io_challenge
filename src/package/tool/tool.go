package tool

import (
	"fmt"
	"os"
)

//Server and Database configuring variables
var (
	ServerPort        string = ":50051"
	ServerNetworkType string = "tcp"
	ClientAdress      string = "localhost:50051"
)

//Database variables
var (
	Driver   string = "sqlserver"
	server   string = os.Getenv("DATABASE_IP")
	port     int    = 1433
	user     string = "sa"
	password string = "123qwe!@#"
	// ConnectionString string = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d; initial catalog=kleverchallenge; integrated security=True;", server, user, password, port)
	ConnectionString string = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d; integrated security=False;", server, user, password, port)
)

//Utility Queries
var (
	SetDatabase        string = "use kleverchallenge "
	DeleteAllFromTable string = "DELETE FROM %s"
	ReseedTable        string = "UPDATE sqlite_sequence SET seq = (SELECT COUNT(*) FROM %s) WHERE name = '%s'"
	DropTable          string = "DROP TABLE %s;"
	SetWaitingTime     string = "WAITFOR DELAY '00:00:015';"
)

//Repository Queries
var (
	SelectAllCryptosQuery                   string = SetDatabase + "SELECT * FROM cryptoCurrencies"
	SelectAllCryptosSortedByNameQuery       string = SetDatabase + "SELECT * FROM cryptoCurrencies ORDER BY name"
	SelectAllCryptosSortedByTokenQuery      string = SetDatabase + "SELECT * FROM cryptoCurrencies ORDER BY token"
	SelectAllCryptosSortedByLeastVotesQuery string = SetDatabase + "SELECT * FROM cryptoCurrencies ORDER BY votes"
	SelectAllCryptosSortedByTopVotesQuery   string = SetDatabase + "SELECT * FROM cryptoCurrencies ORDER BY votes DESC"
	SelectCryptoByIdQuery                   string = SetDatabase + "SELECT * FROM cryptoCurrencies WHERE id = %d"
	SelectCryptoByNameQuery                 string = SetDatabase + "SELECT * FROM cryptoCurrencies WHERE name = '%s'"
	SelectCryptoByTokenQuery                string = SetDatabase + "SELECT * FROM cryptoCurrencies WHERE token = '%s'"
	InsertCryptoQuery                       string = SetDatabase + "INSERT INTO cryptoCurrencies (name, token, votes) OUTPUT Inserted.id, Inserted.name, Inserted.token, Inserted.votes VALUES ('%s', '%s', %d);"
	UpdateCryptoQuery                       string = SetDatabase + "UPDATE cryptoCurrencies SET name = '%s', token = '%s', votes = %d OUTPUT Inserted.id, Inserted.name, Inserted.token, Inserted.votes WHERE id = %d"
	UpvoteCryptoQuery                       string = SetDatabase + "UPDATE cryptoCurrencies SET votes = votes + 1 OUTPUT Inserted.id WHERE id = %d"
	DownvoteCryptoQuery                     string = SetDatabase + "UPDATE cryptoCurrencies SET votes = votes - 1 OUTPUT Inserted.id WHERE id = %d"
	DeleteCryptoById                        string = SetDatabase + "DELETE FROM cryptoCurrencies OUTPUT deleted.id WHERE id = %d"
)

//Database table creation, seeding and utility squeries
var (
	// CreateDatabaseQuery    string = "use master IF NOT EXISTS(SELECT * FROM sys.databases WHERE name = 'kleverchallenge') BEGIN CREATE DATABASE kleverchallenge; END"
	CreateDatabaseQuery    string = "CREATE DATABASE kleverchallenge"
	CreateCryptoTableQuery string = SetDatabase + "IF NOT EXISTS (SELECT * FROM SYSOBJECTS WHERE name =  'cryptoCurrencies') BEGIN CREATE TABLE cryptoCurrencies (id int IDENTITY(1,1) NOT NULL PRIMARY KEY, name varchar(255) NOT NULL, token varchar(255) NOT NULL, votes int NOT NULL); END"
	DeleteOldDataQuery     string = SetDatabase + "DELETE FROM cryptoCurrencies"
	ReseedIdentityQuery    string = SetDatabase + "UPDATE sqlite_sequence SET seq = (SELECT COUNT(*) FROM cryptoCurrencies) WHERE name = 'cryptoCurrencies'"
	SeedCryptoDataQuery    string = SetDatabase + "INSERT INTO cryptoCurrencies (name, token, votes) VALUES ('TetherUS', 'USDT', 47), ('Bitcoin', 'BTC', 41), ('Ethereum', 'ETH', 38), ('TRON', 'TRX', 12), ('Klever', 'KLV', 45), ('Devikins', 'DVK', 22), ('Axie Infinity', 'AXS', 7);"
	DropCryptoTable        string = SetDatabase + "DROP TABLE cryptoCurrencies"
)
