package tool

//Server Default Port
var ServerPort = ":50051"

//Server Newwork Type
var ServerNetworkType = "tcp"

//Client Default Adress
var ClientAdress = "localhost:50051"

//SQLite Provider
var Driver = "sqlite3"

//SQLite Database File
var Database = "./infra/database/kleverchallenge.db"

//Query to create the CRYPTOCURRENCIES table
var CreateCryptoTableQuery = `
CREATE TABLE IF NOT EXISTS "cryptoCurrencies" (
	"id"	INTEGER NOT NULL UNIQUE,
	"name"	TEXT NOT NULL,
	"token"	TEXT NOT NULL,
	"votes" INTEGER NOT NULL,
	PRIMARY KEY("id" AUTOINCREMENT)
);
`

//Delete all data in CRYPTOCURRENCIES before start SEEDING
var DeleteOldDataQuery = "DELETE FROM cryptoCurrencies"

//Reset the Identity Sequence number count
var ReseedIdentityQuery = "UPDATE sqlite_sequence SET seq = (SELECT COUNT(*) FROM cryptoCurrencies) WHERE name = 'cryptoCurrencies'"

//Query to SEED some DATA in the CRYPTOCURRENCIES table
var SeedCryptoDataQuery = `
INSERT INTO cryptoCurrencies (name, token, votes) values 
	('TetherUS', 'USDT', 47), ('Bitcoin', 'BTC', 41), 
	('Ethereum', 'ETH', 38), ('TRON', 'TRX', 12), 
	('Klever', 'KLV', 45), ('Devikins', 'DVK', 22), 
	('Axie Infinity', 'AXS', 7)
; 
`

//Querries to Drop CRYPTOCURRENCY table
var DropCryptoTable = "DROP TABLE 'cryptoCurrencies';"

//RRepository Queries -> CryptoCurrency Table
var SelectAllCryptosQuery = "SELECT * FROM cryptoCurrencies"

var SelectAllCryptosSortedByNameQuery = "SELECT * FROM cryptoCurrencies ORDER BY name"

var SelectAllCryptosSortedByTokenQuery = "SELECT * FROM cryptoCurrencies ORDER BY token"

var SelectAllCryptosSortedByLeastVotesQuery = "SELECT * FROM cryptoCurrencies ORDER BY votes"

var SelectAllCryptosSortedByTopVotesQuery = "SELECT * FROM cryptoCurrencies ORDER BY votes DESC"

var SelectCryptoByIdQuery = "SELECT * FROM cryptoCurrencies WHERE id = ?"

var SelectCryptoByNameQuery = "SELECT * FROM cryptoCurrencies WHERE name = '%s'"

var SelectCryptoByTokenQuery = "SELECT * FROM cryptoCurrencies WHERE token = '%s'"

var InsertCrypto = "INSERT INTO cryptoCurrencies (name, token, votes) values (?, ?, ?); select last_insert_rowid() id;"

var UpdateCryptoQuery = "UPDATE cryptoCurrencies SET name = ?, token = ?, votes = ? WHERE id = ?"

var DeleteCryptoById = "DELETE FROM cryptoCurrencies WHERE id = ?"

var UpvoteCryptoQuery = "UPDATE cryptoCurrencies SET votes = votes + 1 WHERE id = ?"

var DownvoteCryptoQuery = "UPDATE cryptoCurrencies SET votes = votes - 1 WHERE id = ?"

//Other Queries to be used when added new tables
var DeleteAllFromTable = "DELETE FROM %s"

var ReseedTable = "UPDATE sqlite_sequence SET seq = (SELECT COUNT(*) FROM %s) WHERE name = '%s'"

var DropTable = "DROP TABLE %s;"
