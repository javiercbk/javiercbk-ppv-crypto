package main

//go:generate sqlboiler psql

import (
	"database/sql"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/javiercbk/ppv-crypto/server/cryptocurrency/eth"
	"github.com/javiercbk/ppv-crypto/server/http"
)

const (
	defaultLogFilePath       = "stdout"
	defaultJWTSecret         = "ppvcrypto"
	defaultAddress           = "0.0.0.0:9000"
	defaultDBHost            = "127.0.0.1"
	defaultDBName            = "cryptoc"
	defaultDBUser            = "cryptoc"
	defaultDBPass            = "cryptoc"
	defaultETHHost           = "http://127.0.0.1:8545"
	defaultETHPrivateKeyPath = "eth/keys/ppv_private_key_1"
)

func main() {
	var logFilePath, address, jwtSecret, dbName, dbHost, dbUser, dbPass, ethHost, ethPrivateKeyPath string
	flag.StringVar(&logFilePath, "l", defaultLogFilePath, "the log file location")
	flag.StringVar(&address, "a", defaultAddress, "the http server address")
	flag.StringVar(&jwtSecret, "jwt", defaultJWTSecret, "the jwt secret")
	flag.StringVar(&dbName, "dbn", defaultDBName, "the database name")
	flag.StringVar(&dbHost, "dbh", defaultDBHost, "the database host")
	flag.StringVar(&dbUser, "dbu", defaultDBUser, "the database user")
	flag.StringVar(&dbPass, "dbp", defaultDBPass, "the database password")
	flag.StringVar(&ethHost, "ethHost", defaultETHHost, "the eth network host to connect")
	flag.StringVar(&ethPrivateKeyPath, "ethPrivate", defaultETHPrivateKeyPath, "the file path to the eth private key")
	flag.Parse()
	var logWritter io.Writer
	if logFilePath != "stdout" {
		logFile, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			fmt.Printf("error opening lof file: %s", err)
			os.Exit(1)
		}
		defer logFile.Close()
		logWritter = logFile
	} else {
		logWritter = os.Stdout
	}
	logger := log.New(logWritter, "applog: ", log.Lshortfile|log.LstdFlags)
	privateKeyFile, err := os.Open(ethPrivateKeyPath)
	if err != nil {
		logger.Printf("error opening eth private key: %v", err)
		os.Exit(1)
	}
	privateKeyByte, err := ioutil.ReadAll(privateKeyFile)
	if err != nil {
		logger.Printf("error reading eth private key: %v", err)
		os.Exit(1)
	}
	ethClient, err := ethclient.Dial(ethHost)
	if err != nil {
		logger.Printf("error connecting to eth network: %v", err)
		os.Exit(1)
	}
	contractDeployer, err := eth.NewSmartContractDeployer(ethClient, logger, privateKeyByte)
	if err != nil {
		logger.Printf("failed to create contract deployer: %v", err)
		os.Exit(1)
	}
	db, err := connectPostgres(dbName, dbHost, dbUser, dbPass, logger)
	if err != nil {
		logger.Printf("error connecting to postgres: %s", err)
		os.Exit(1)
	}
	cnf := http.Config{
		Address:   address,
		JWTSecret: jwtSecret,
	}
	logger.Printf("server is initializing\n")
	err = http.Serve(cnf, logger, db, contractDeployer)
	if err != nil {
		logger.Fatalf("could not start server %s\n", err)
	}
}

func connectPostgres(dbName, dbHost, dbUser, dbPass string, logger *log.Logger) (*sql.DB, error) {
	logger.Printf("connecting to postgres server\n")
	postgresOpts := fmt.Sprintf("dbname=%s host=%s user=%s password=%s sslmode=disable", dbName, dbHost, dbUser, dbPass)
	db, err := sql.Open("postgres", postgresOpts)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
