package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/codeedu/codebank/application/grpc/server"
	"github.com/codeedu/codebank/data/protocols"
	"github.com/codeedu/codebank/data/usecase"
	"github.com/codeedu/codebank/infrastructure/db/repository"
	"github.com/codeedu/codebank/infrastructure/kafka"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	db := setupDb()
	defer db.Close()
	producer := setupKafkaProducer()
	processTransactionUseCase := setupTransactionUseCase(db, producer)
	serveGrpc(processTransactionUseCase)
}

func setupTransactionUseCase(db *sql.DB, producer protocols.KafkaProducer) usecase.UseCaseTransaction {
	transactionRepository := repository.NewTransactionRepositoryDb(db)
	useCase := usecase.NewUseCaseTransaction(transactionRepository)
	useCase.KafkaProducer = producer
	useCase.TransactionTopic = os.Getenv("KAFKA_TRANSACTIONS_TOPIC")
	return useCase
}

func setupKafkaProducer() *kafka.KafkaProducer {
	producer := kafka.NewKafkaProducer()
	producer.SetupProducer(os.Getenv("KAFKA_BOOTSTRAP_SERVERS"))
	return producer
}

func setupDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("error connection to database")
	}
	return db
}

func serveGrpc(processTransactionUseCase usecase.UseCaseTransaction) {
	grpcServer := server.NewGRPCServer()
	grpcServer.ProcessTransactionUseCase = processTransactionUseCase
	grpcServer.Serve()
}
