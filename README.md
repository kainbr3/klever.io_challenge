![Be Klever!.](https://cdn.substack.com/image/fetch/w_1100,c_limit,f_auto,q_auto:good,fl_progressive:steep/https%3A%2F%2Fbucketeer-e05bbc84-baa3-437e-9518-adb32be77984.s3.amazonaws.com%2Fpublic%2Fimages%2F91702ea8-69da-4dab-a807-dd0325e836e9_1920x1080.png "Klever Logo")

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
