![Be Klever!.](https://pbs.twimg.com/profile_banners/1389545109973200896/1640750305/1500x500 "Klever Logo")

# Klever Technical Challenge

## About the challenge
*The Technical Challenge consists of creating an API with Golang using gRPC with stream pipes that exposes an Upvote service endpoints. The API will provide the user an interface to upvote or downvote a known list of the main Cryptocurrencies (Bitcoin, ethereum, litecoin, etc..).*

## Technical requirements:
*(Keep the code in Github)*  

* The API must have a read, insert, delete and update interfaces.
* The API must have a method that stream a live update of the current sum of the votes from a given Cryptocurrency
* The API must guarantee the typing of user inputs. If an input is expected as a string, it can only be received as a string.
* The API must contain unit test of methods it uses
* You can choose the database but the structs used with it should support Marshal/Unmarshal with bson, json and struct

*Extra:*
 * Deliver the whole solution running in some free cloud service
 * Job to take snapshots of the votes every hour and plot a graph

## Project Detais
| Type  | Detail |
| --------------- |:-------------:|
| Server Language | Go            |
| Client Language | Go            |
| Type            | gRPC + API    |
| Database        | SqLite        |
| Frontend        | .Net6 Web MVC |
| Cloud Host      |     -         |

## Installation
 *Project requirement:* 
 * [GO (Golang)](https://go.dev/) to run server and client.
 * [.Net 6 (Dotnet 6)](https://dotnet.microsoft.com/en-us/download/dotnet/6.0) to run the Additional Frontend.

## Plugins
| Plugin  | Readme |
| --------------- |:-------------:|
| gRPC Web UI | https://github.com/fullstorydev/grpcui |
| Go SQLite3  | https://github.com/mattn/go-sqlite3    |

##### Protobuf File NEEDED OPTIONS
 ```sh
 Golang Version must add this
 option go_package = "github.com/kainbr3/klever.io_challenge/protobuf;protobuf"; 
 
 C# Version must add this
 option csharp_namespace = "KleverGrpcClient";
 ```

##### gRPC Files Generation: (Golang)
 ```sh
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protobuf/service.proto
```

*service.pb.go => Responsible to Serialize and Deserialize the messages defined in service definitions
service.grpc.pb.go => Contains the auto generated Client and Server Code that we need to implement in our own Client and Server programs*

##### gRPC Files Generation: (C# - CSharp)
 ```sh
 Point the service.proto in the src/protobuf fold and it will be generated when you run/build the donet project
 
 If you want to use the manual command to generate the files, use this
protoc --proto_path=. --csharp_out=library=service_pb,binary:protobuf --grpc-web_out=import_style=commonjs,mode=grpcwebtext:. protobuf/service.proto
```

##### Dedepndencies and Package Go (Golang)
*From \SRC Folder:*
 ```sh
go mod tidy
```

##### Dedepndencies and Package C# (CSharp)
*From /SRC/FRONTEND Folder:*
 ```sh
dotnet restore
dotnet build
```

## Starting the Server
*From /ROOT Folder:*
```sh
go run main.go
```
server ip: Port
grpcui -plaintext localhost:50051

## Starting the CLient
*From /ROOT Folder:*
```sh
go run src/command/server/client.go
```

## Starting the FRONT END
*From /SRC/FRONTEND Folder:*
```sh
dotnet run
<Command with Hot Reload Support>
dotnet watch run
```

## Project Structure
```
📦klever.io_challenge
 ┣ 📂.git
 ┣ 📂.vscode
 ┣ 📂src
 ┃ ┣ 📂command
 ┃ ┃ ┣ 📂client
 ┃ ┃ ┃ ┗ 📜client.go
 ┃ ┃ ┗ 📂server
 ┃ ┃ ┃ ┗ 📜server.go
 ┃ ┣ 📂frontend
 ┃ ┃ ┗ 📂kleverchallenge
 ┃ ┃ ┃ ┣ 📂Controllers
 ┃ ┃ ┃ ┃ ┣ 📜CryptoController.cs
 ┃ ┃ ┃ ┃ ┗ 📜HomeController.cs
 ┃ ┃ ┃ ┣ 📂Models
 ┃ ┃ ┃ ┃ ┣ 📜CryptoViewModel.cs
 ┃ ┃ ┃ ┃ ┗ 📜ErrorViewModel.cs
 ┃ ┃ ┃ ┣ 📂Properties
 ┃ ┃ ┃ ┃ ┗ 📜launchSettings.json
 ┃ ┃ ┃ ┣ 📂Views
 ┃ ┃ ┃ ┃ ┣ 📂Crypto
 ┃ ┃ ┃ ┃ ┃ ┣ 📜Add.cshtml
 ┃ ┃ ┃ ┃ ┃ ┣ 📜Delete.cshtml
 ┃ ┃ ┃ ┃ ┃ ┣ 📜List.cshtml
 ┃ ┃ ┃ ┃ ┃ ┗ 📜Update.cshtml
 ┃ ┃ ┃ ┃ ┣ 📂Home
 ┃ ┃ ┃ ┃ ┃ ┗ 📜Index.cshtml
 ┃ ┃ ┃ ┃ ┣ 📂Shared
 ┃ ┃ ┃ ┃ ┃ ┣ 📜Error.cshtml
 ┃ ┃ ┃ ┃ ┃ ┣ 📜_Layout.cshtml
 ┃ ┃ ┃ ┃ ┃ ┣ 📜_Layout.cshtml.css
 ┃ ┃ ┃ ┃ ┃ ┗ 📜_ValidationScriptsPartial.cshtml
 ┃ ┃ ┃ ┃ ┣ 📜_ViewImports.cshtml
 ┃ ┃ ┃ ┃ ┗ 📜_ViewStart.cshtml
 ┃ ┃ ┃ ┣ 📂wwwroot
 ┃ ┃ ┃ ┃ ┣ 📂css
 ┃ ┃ ┃ ┃ ┃ ┗ 📜site.css
 ┃ ┃ ┃ ┃ ┣ 📂img
 ┃ ┃ ┃ ┃ ┃ ┣ 📜404.png
 ┃ ┃ ┃ ┃ ┃ ┣ 📜AXS.png
 ┃ ┃ ┃ ┃ ┃ ┣ 📜BTC.png
 ┃ ┃ ┃ ┃ ┃ ┣ 📜DVK.png
 ┃ ┃ ┃ ┃ ┃ ┣ 📜ETH.png
 ┃ ┃ ┃ ┃ ┃ ┣ 📜KLV.png
 ┃ ┃ ┃ ┃ ┃ ┣ 📜TRX.png
 ┃ ┃ ┃ ┃ ┃ ┗ 📜USDT.png
 ┃ ┃ ┃ ┃ ┣ 📂js
 ┃ ┃ ┃ ┃ ┃ ┗ 📜site.js
 ┃ ┃ ┃ ┃ ┣ 📂lib
 ┃ ┃ ┃ ┃ ┗ 📜favicon.ico
 ┃ ┃ ┃ ┣ 📜appsettings.Development.json
 ┃ ┃ ┃ ┣ 📜appsettings.json
 ┃ ┃ ┃ ┣ 📜kleverchallenge.csproj
 ┃ ┃ ┃ ┗ 📜Program.cs
 ┃ ┣ 📂infra
 ┃ ┃ ┗ 📂database
 ┃ ┃ ┃ ┗ 📜kleverchallenge.db
 ┃ ┣ 📂package
 ┃ ┃ ┣ 📂model
 ┃ ┃ ┃ ┗ 📜model.go
 ┃ ┃ ┣ 📂repository
 ┃ ┃ ┃ ┗ 📜repository.go
 ┃ ┃ ┗ 📂tool
 ┃ ┃ ┃ ┗ 📜tool.go
 ┃ ┗ 📂protobuf
 ┃ ┃ ┣ 📜service.pb.go
 ┃ ┃ ┣ 📜service.proto
 ┃ ┃ ┗ 📜service_grpc.pb.go
 ┣ 📜.gitignore
 ┣ 📜go.mod
 ┣ 📜go.sum
 ┣ 📜LICENSE
 ┣ 📜main.go
 ┣ 📜Makefile
 ┗ 📜README.md
 ```

  ## TODO List
- [x] Create Server and Client Go Applications
- [x] Organize all the reusable and shared data in Tool Package
- [x] Create and Configure a Database
- [x] Create a Repository Package to handle the Database Persistence
- [x] Create the Protobuf (Contract) file and generate the PB and gRPC Service
- [x] Create a Frontend Application to consume the gRPC Services
- [ ] Create New gRPC Services, including a Streaming to show Crypto Votes in Real Time
- [ ] Add more views in the Frontend Application to iterate with the new gRPC Services
- [ ] Create the Unit Tests
- [ ] Test the Marshal/Unmarshal support
- [ ] Deploy the Solution in a free Coud Service
- [ ] Configure the Cloud Host Machine Snapshots
- [ ] Create a Graph to show Host Snapshots