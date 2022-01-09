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
| ------------- |:-------------:|
| Language      | Go            |
| Database      | SqLite        |
| Type          | gRPC + API    |


## Installation
 *Project requirement:* 
 * [GO (Golang)](https://go.dev/) to run server and client.
 * [.Net 6 (Dotnet 6)](https://dotnet.microsoft.com/en-us/download/dotnet/6.0) to run the Additional Frontend.

##### Protobuf File NEEDED OPTIONS
 ```
 Golang Version must add this
 option go_package = "github.com/kainbr3/klever.io_challenge/protobuf;protobuf"; 
 
 C# Version must add this
 option csharp_namespace = "KleverGrpcClient";
 ```

##### gRPC Files Generation: (Golang)
 ```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protobuf/service.proto
```

*service.pb.go => Responsible to Serialize and Deserialize the messages defined in service definitions
service.grpc.pb.go => Contains the auto generated Client and Server Code that we need to implement in our own Client and Server programs*

##### gRPC Files Generation: (C# - CSharp)
 ```
protoc --proto_path=. --csharp_out=library=service_pb,binary:protobuf --grpc-web_out=import_style=commonjs,mode=grpcwebtext:. protobuf/service.proto
```

##### Dedepndencies and Package Go (Golang)
*From \SRC Folder:*
 ```
go mod tidy
```

##### Dedepndencies and Package C# (CSharp)
*From \SRC\FRONTEND Folder:*
 ```
dotnet restore
dotnet build
```

## Starting the Server
*From \SRC Folder:*
```
go run command/server/server.go
```

## Starting the CLient
*From \SRC Folder:*
```
go run command/server/client.go
```

## Starting the FRONT END
*From \SRC|FRONTEND Folder:*
```
dotnet run
<Command with Hot Reload Support>
dotnet watch run
```

## Project Structure
```
📦klever-challenge
 ┣ 📂src
 ┃ ┣ 📂command
 ┃ ┃ ┣ 📂client
 ┃ ┃ ┃ ┗ 📜client.go
 ┃ ┃ ┗ 📂server
 ┃ ┃ ┃ ┗ 📜server.go
 ┃ ┃ ┣ 📂frontend
 ┃ ┃ ┃ ┗ 📂kleverchallenge
 ┃ ┃ ┃   ┣ 📂Controllers
 ┃ ┃ ┃   ┃ ┣ 📜CryptoController.cs
 ┃ ┃ ┃   ┃ ┗ 📜HomeController.cs
 ┃ ┃ ┃   ┣ 📂Models
 ┃ ┃ ┃   ┃ ┗ 📜ErrorViewModel.cs
 ┃ ┃ ┃   ┣ 📂Properties
 ┃ ┃ ┃   ┃ ┗ 📜launchSettings.json
 ┃ ┃ ┃   ┣ 📂protobuf
 ┃ ┃ ┃   ┃ ┗ 📜service.proto
 ┃ ┃ ┃   ┣ 📂Views
 ┃ ┃ ┃   ┃ ┣ 📂Crypto
 ┃ ┃ ┃   ┃ ┃ ┣ 📜Add.cshtml
 ┃ ┃ ┃   ┃ ┃ ┣ 📜Delete.cshtml
 ┃ ┃ ┃   ┃ ┃ ┣ 📜List.cshtml
 ┃ ┃ ┃   ┃ ┃ ┗ 📜Update.cshtml
 ┃ ┃ ┃   ┃ ┣ 📂Home
 ┃ ┃ ┃   ┃ ┃ ┗ 📜Index.cshtml
 ┃ ┃ ┃   ┃ ┣ 📂Shared
 ┃ ┃ ┃   ┃ ┃ ┣ 📜Error.cshtml
 ┃ ┃ ┃   ┃ ┃ ┣ 📜_Layout.cshtml
 ┃ ┃ ┃   ┃ ┃ ┣ 📜_Layout.cshtml.css
 ┃ ┃ ┃   ┃ ┃ ┗ 📜_ValidationScriptsPartial.cshtml
 ┃ ┃ ┃   ┃ ┣ 📜_ViewImports.cshtml
 ┃ ┃ ┃   ┃ ┗ 📜_ViewStart.cshtml
 ┃ ┃ ┃   ┣ 📂wwwroot
 ┃ ┃ ┃   ┃ ┣ 📂css
 ┃ ┃ ┃   ┃ ┃ ┗ 📜site.css
 ┃ ┃ ┃   ┃ ┣ 📂img
 ┃ ┃ ┃   ┃ ┃ ┣ 📜AXS.png
 ┃ ┃ ┃   ┃ ┃ ┣ 📜BTC.png
 ┃ ┃ ┃   ┃ ┃ ┣ 📜DVK.png
 ┃ ┃ ┃   ┃ ┃ ┣ 📜ETH.png
 ┃ ┃ ┃   ┃ ┃ ┣ 📜KLV.png
 ┃ ┃ ┃   ┃ ┃ ┣ 📜TRX.png
 ┃ ┃ ┃   ┃ ┃ ┗ 📜USDT.png
 ┃ ┃ ┃   ┃ ┣ 📂js
 ┃ ┃ ┃   ┃ ┃ ┗ 📜site.js
 ┃ ┃ ┃   ┃ ┣ 📂lib
 ┃ ┃ ┃   ┃ ┗ 📜favicon.ico
 ┃ ┃ ┃   ┣ 📜appsettings.Development.json
 ┃ ┃ ┃   ┣ 📜appsettings.json
 ┃ ┃ ┃   ┣ 📜kleverchallenge.csproj
 ┃ ┃ ┃   ┗ 📜Program.cs
 ┃ ┣ 📂infra
 ┃ ┃ ┗ 📂database
 ┃ ┃ ┃ ┗ 📜kleverchallenge.db
 ┃ ┣ 📂package
 ┃ ┃ ┣ 📂api
 ┃ ┃ ┃ ┗ 📜api.go
 ┃ ┃ ┣ 📂model
 ┃ ┃ ┃ ┗ 📜model.go
 ┃ ┃ ┣ 📂repository
 ┃ ┃ ┃ ┗ 📜repository.go
 ┃ ┃ ┗ 📂tool
 ┃ ┃ ┃ ┗ 📜tool.go
 ┃ ┣ 📂protobuf
 ┃ ┃ ┣ 📜service.pb.go
 ┃ ┃ ┣ 📜service.proto
 ┃ ┃ ┗ 📜service_grpc.pb.go
 ┃ ┣ 📜go.mod
 ┃ ┗ 📜go.sum
 ┣ 📜.gitignore
 ┣ 📜LICENSE
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