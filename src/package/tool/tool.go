package tool

//Server and Database configuring variables
var (
	ServerPort        string = ":50051"
	ServerNetworkType string = "tcp"
	ClientAdress      string = "localhost:50051"
	Driver            string = "sqlite3"
	Database          string = "./src/infra/database/kleverchallenge.db"
)

//Utility Queries
var (
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
	SelectCryptoByIdQuery                   string = "SELECT * FROM cryptoCurrencies WHERE id = ?"
	SelectCryptoByNameQuery                 string = "SELECT * FROM cryptoCurrencies WHERE name = '%s'"
	SelectCryptoByTokenQuery                string = "SELECT * FROM cryptoCurrencies WHERE token = '%s'"
	InsertCrypto                            string = "INSERT INTO cryptoCurrencies (name, token, votes) values (?, ?, ?);"
	UpdateCryptoQuery                       string = "UPDATE cryptoCurrencies SET name = ?, token = ?, votes = ? WHERE id = ?"
	DeleteCryptoById                        string = "DELETE FROM cryptoCurrencies WHERE id = ?"
	UpvoteCryptoQuery                       string = "UPDATE cryptoCurrencies SET votes = votes + 1 WHERE id = ?"
	DownvoteCryptoQuery                     string = "UPDATE cryptoCurrencies SET votes = votes - 1 WHERE id = ?"
)

//Database table creation, seeding and utility squeries
var (
	CreateCryptoTableQuery string = `
		CREATE TABLE IF NOT EXISTS "cryptoCurrencies" (
			"id"	INTEGER NOT NULL UNIQUE,
			"name"	TEXT NOT NULL,
			"token"	TEXT NOT NULL,
			"votes" INTEGER NOT NULL,
			PRIMARY KEY("id" AUTOINCREMENT)
		);
	`
	DeleteOldDataQuery string = "DELETE FROM cryptoCurrencies"

	ReseedIdentityQuery string = "UPDATE sqlite_sequence SET seq = (SELECT COUNT(*) FROM cryptoCurrencies) WHERE name = 'cryptoCurrencies'"

	SeedCryptoDataQuery string = `
		INSERT INTO cryptoCurrencies (name, token, votes) values 
			('TetherUS', 'USDT', 47), ('Bitcoin', 'BTC', 41), 
			('Ethereum', 'ETH', 38), ('TRON', 'TRX', 12), 
			('Klever', 'KLV', 45), ('Devikins', 'DVK', 22), 
			('Axie Infinity', 'AXS', 7)
		; 
	`
	DropCryptoTable string = "DROP TABLE 'cryptoCurrencies';"
)
