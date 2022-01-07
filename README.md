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
