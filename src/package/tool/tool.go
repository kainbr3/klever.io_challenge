package tool

import "fmt"

//Server and Database configuring variables
var (
	ServerPort        string = ":50051"
	ServerNetworkType string = "tcp"
	ClientAdress      string = "127.0.0.1:50051"
)

//Database variables
var (
	Driver           string = "sqlserver"
	server           string = "localhost"
	port             int    = 1433
	user             string = "sa"
	password         string = "123qwe!@#"
	ConnectionString string = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", server, user, password, port)
)

//Utility Queries
var (
	SetDatabase        string = "use kleverchallenge "
	DeleteAllFromTable string = "DELETE FROM %s"
	ReseedTable        string = "UPDATE sqlite_sequence SET seq = (SELECT COUNT(*) FROM %s) WHERE name = '%s'"
	DropTable          string = "DROP TABLE %s;"
)

//Repository Queries
var (
	SelectAllCryptosQuery                   string = "SELECT * FROM cryptoCurrencies"
	SelectAllCryptosSortedByNameQuery       string = "SELECT * FROM cryptoCurrencies ORDER BY name"
	SelectAllCryptosSortedByTokenQuery      string = "SELECT * FROM cryptoCurrencies ORDER BY token"
	SelectAllCryptosSortedByLeastVotesQuery string = "SELECT * FROM cryptoCurrencies ORDER BY votes"
	SelectAllCryptosSortedByTopVotesQuery   string = "SELECT * FROM cryptoCurrencies ORDER BY votes DESC"
	SelectCryptoByIdQuery                   string = "SELECT * FROM cryptoCurrencies WHERE id = %d"
	SelectCryptoByNameQuery                 string = "SELECT * FROM cryptoCurrencies WHERE name = '%s'"
	SelectCryptoByTokenQuery                string = "SELECT * FROM cryptoCurrencies WHERE token = '%s'"
	InsertCrypto                            string = "INSERT INTO cryptoCurrencies (name, token, votes) OUTPUT Inserted.id, Inserted.name, Inserted.token, Inserted.votes VALUES (%s, %s, %d);"
	UpdateCryptoQuery                       string = "UPDATE cryptoCurrencies SET name = %s, token = %s, votes = %d WHERE id = %d"
	DeleteCryptoById                        string = "DELETE FROM cryptoCurrencies WHERE id = %d"
	UpvoteCryptoQuery                       string = "UPDATE cryptoCurrencies SET votes = votes + 1 WHERE id = %d"
	DownvoteCryptoQuery                     string = "UPDATE cryptoCurrencies SET votes = votes - 1 WHERE id = %d"
)

//Database table creation, seeding and utility squeries
var (
	CreateDatabaseQuery    string = "IF NOT EXISTS(SELECT * FROM sys.databases WHERE name = 'kleverchallenge') BEGIN CREATE DATABASE kleverchallenge; END"
	CreateCryptoTableQuery string = "IF NOT EXISTS (SELECT * FROM SYSOBJECTS WHERE name =  'cryptoCurrencies') BEGIN CREATE TABLE cryptoCurrencies (id int IDENTITY(1,1) NOT NULL PRIMARY KEY, name varchar(255) NOT NULL, token varchar(255) NOT NULL, votes int NOT NULL); END"
	DeleteOldDataQuery     string = "DELETE FROM cryptoCurrencies"
	ReseedIdentityQuery    string = "UPDATE sqlite_sequence SET seq = (SELECT COUNT(*) FROM cryptoCurrencies) WHERE name = 'cryptoCurrencies'"
	SeedCryptoDataQuery    string = "INSERT INTO cryptoCurrencies (name, token, votes) VALUES ('TetherUS', 'USDT', 47), ('Bitcoin', 'BTC', 41), ('Ethereum', 'ETH', 38), ('TRON', 'TRX', 12), ('Klever', 'KLV', 45), ('Devikins', 'DVK', 22), ('Axie Infinity', 'AXS', 7);"
	DropCryptoTable        string = "DROP TABLE cryptoCurrencies"
)
