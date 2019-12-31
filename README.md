# ppv-crypto

This is a research project that aims to discover how to charge for a pay per view event using cryptocurrencies.

## Cryptocurrencies

Since there are a LOT of cryptocurrencies the main focus is to cover different use cases. As December 2019 there are 3 cryptocurrencies that I'm currently interested in researching:

- ETH: With smart contracts, Ethereum seems to be the best crypto currency for handling buying a Pay Per View Event. EOS is also interesting in the same regard but I've chosen ETH because it seemed to have the best development environment, and arguably the best integration with GO.
- BTC: Bitcoin is not the best cryptocurrency, still it is the most well know so I could not leave it out.
- XMR: Monero is a truly fungible, privacy centered cryptocurrency. Due to its privacy features, it is interesting how to identify payments.

## Dependencies

### Node

Install nodejs and download the following dependencies

```sh
npm install -g solc truffle @vue/cli-service
cd crypto-frontend
npm install
```

### Golang

Install GolangCI by following the method that suit you best listed in [https://github.com/golangci/golangci-lint#install](https://github.com/golangci/golangci-lint#install)

To install `abi` command we'll have to download go-ethereum and compile it

```sh
go get github.com/ethereum/go-ethereum
cd $GOPATH/src/github.com/ethereum/go-ethereum/
make
make geth
make devtools
```

### Docker

Docker is needed to spin up a container with a Postgres database.

## Compile the backend server

Run `make` on the project folder

## Set up the database

In the root folder run

```sh
docker-compose up
```

Open up another terminal and run

```sh
# get the container id
docker ps
docker exec -ti <container_id> bash
#inside the container
psql -U cryptoc
```

Now the following commands are executed inside the `psql` interpreter

```sh
# I always like to recreate the database with the proper encoding
\c postgres
DROP DATABASE cryptoc;
# from the file "schema.sql" located in the root folder of the project
CREATE DATABASE cryptoc WITH OWNER 'cryptoc' ENCODING 'UTF8';
\c cryptoc
# paste the whole content of the schema.sql file
```
