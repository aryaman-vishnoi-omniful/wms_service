package appinit

import (
	"context"
	"fmt"

	// "log"
	kafka_producer "wms_service/kafka"
	"wms_service/postgres"
	"wms_service/redis"

	// "oms_service/database"

	// "oms_service/orders"
	// "oms_service/orders/listners"
	// "oms_service/orders/services"
	// "oms_service/redis"
	// "os"
	"time"

	// "github.com/joho/godotenv"
	// "github.com/joho/godotenv"
	"github.com/omniful/go_commons/config"
	"github.com/omniful/go_commons/log"

	"github.com/omniful/go_commons/kafka"

	// "github.com/omniful/go_commons/kafka"
	opostgres "github.com/omniful/go_commons/db/sql/postgres"
	goredis "github.com/omniful/go_commons/redis"
	// "github.com/omniful/go_commons/sqs"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

// type contextKey string

// const DBKey contextKey = "mongoDB"

// var DB *mongo.Client
// func init() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		fmt.Println("unable to lload")
// 	}

// }

func Initialize(ctx context.Context) {
	InitializeRedis(ctx)
	InitializeDB(ctx)
	InitializeKafka(ctx)
	// InitializeSQS(ctx)
	// return ctx
}

// godotenv.load()

func InitializeDB(ctx context.Context) {
	maxOpenConnections := config.GetInt(ctx, "postgresql.maxOpenConns")
	maxIdleConnections := config.GetInt(ctx, "postgresql.maxIdleConns")

	database := config.GetString(ctx, "postgresql.database")
	connIdleTimeout := 10 * time.Minute

	// Read Write endpoint config
	mysqlWriteServer := config.GetString(ctx, "postgresql.master.host")
	mysqlWritePort := config.GetString(ctx, "postgresql.master.port")
	mysqlWritePassword := config.GetString(ctx, "postgresql.master.password")
	mysqlWriterUsername := config.GetString(ctx, "postgresql.master.username")

	// Fetch Read endpoint config
	//mysqlReadServers := config.GetString(ctx, "postgresql.slaves.hosts")
	//mysqlReadPort := config.GetString(ctx, "postgresql.slaves.port")
	//mysqlReadPassword := config.GetString(ctx, "postgresql.slaves.password")
	//mysqlReadUsername := config.GetString(ctx, "postgresql.slaves.username")

	debugMode := config.GetBool(ctx, "postgresql.debugMode")

	// Master config i.e. - Write endpoint
	masterConfig := opostgres.DBConfig{
		Host:               mysqlWriteServer,
		Port:               mysqlWritePort,
		Username:           mysqlWriterUsername,
		Password:           mysqlWritePassword,
		Dbname:             database,
		MaxOpenConnections: maxOpenConnections,
		MaxIdleConnections: maxIdleConnections,
		ConnMaxLifetime:    connIdleTimeout,
		DebugMode:          debugMode,
	}

	// Slave config i.e. - array with read endpoints
	slavesConfig := make([]opostgres.DBConfig, 0)
	//for _, host := range strings.Split(mysqlReadServers, ",") {
	//	slaveConfig := opostgres.DBConfig{
	//		Host:               host,
	//		Port:               mysqlReadPort,
	//		Username:           mysqlReadUsername,
	//		Password:           mysqlReadPassword,
	//		Dbname:             database,
	//		MaxOpenConnections: maxOpenConnections,
	//		MaxIdleConnections: maxIdleConnections,
	//		ConnMaxLifetime:    connIdleTimeout,
	//		DebugMode:          debugMode,
	//	}
	//	slavesConfig = append(slavesConfig, slaveConfig)
	//}

	db := opostgres.InitializeDBInstance(masterConfig, &slavesConfig)
	log.InfofWithContext(ctx, "Initialized Postgres DB client")
	// database.SetCluster(db)
	postgres.SetCluster(db)
}

func InitializeKafka(ctx context.Context) {
	kafkaBrokers := config.GetStringSlice(ctx, "onlineKafka.brokers")
	kafkaClientID := config.GetString(ctx, "onlineKafka.clientId")
	kafkaVersion := config.GetString(ctx, "onlineKafka.version")
	producer := kafka.NewProducer(
		kafka.WithBrokers(kafkaBrokers),
		kafka.WithClientID(kafkaClientID),
		kafka.WithKafkaVersion(kafkaVersion),
	)
	// fmt.Println("Initialized kafka producer")
	log.Printf("Initialized Kafka Producer")
	kafka_producer.Set(producer)

}

func InitializeRedis(ctx context.Context) {
	redis_client := goredis.NewClient(&goredis.Config{
		ClusterMode: config.GetBool(ctx, "redis.clusterMode"),
		Hosts:       []string{config.GetString(ctx, "redis.hosts")},
		DB:          config.GetUint(ctx, "redis.db"),
	})
	fmt.Println("Initialized redis client")

	redis.SetClient(redis_client)

}

// func InitializeSQS(ctx context.Context) {
// 	SQSconfig := sqs.GetSQSConfig(ctx, false, "order", "eu-north-1", os.Getenv("AWS_ACCOUNT"), "")
// 	queue_url, err := sqs.GetUrl(ctx, SQSconfig, "samplequeue.fifo")
// 	if err != nil {
// 		log.Fatal("cant get url")
// 	}
// 	// log.Printf("Successfully initialized SQS. Queue URL: %s", *queue_url)
// 	Queue_instance, err := sqs.NewFifoQueue(ctx, "samplequeue.fifo", SQSconfig)
// 	if err != nil {
// 		log.Fatal("cant create queue instance")
// 	}
// 	// fmt.Println(Queue_instance,*Queue_instance.Url)
// 	services.SetProducer(ctx, Queue_instance)
// 	go listners.StartConsume(*queue_url, ctx)

// }
