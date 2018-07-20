package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/vectorhacker/bank/core/bedrock"
	"github.com/vectorhacker/bank/core/events"
	"github.com/vectorhacker/bank/service.accounts/internal/pkg/command"
	"github.com/vectorhacker/bank/service.accounts/internal/pkg/models"
	pb "github.com/vectorhacker/bank/service.accounts/pb"
	ad "github.com/vectorhacker/bank/service.accounts/pkg/events"
	"google.golang.org/grpc"
)

var (
	connection = flag.String("db", os.Getenv("DB_CONNECTION"), "db database connection string")
	brokers    = flag.String("brokers", os.Getenv("KAFKA_PEERS"), "The Kafka brokers to connect to, as a comma separated list")
)

func main() {

	flag.Parse()

	if *connection == "" {
		log.Fatal("expected database connection")
	}

	if *brokers == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	brokerList := strings.Split(*brokers, ",")
	log.Printf("Kafka brokers: %s", strings.Join(brokerList, ", "))

	db, err := gorm.Open("postgres", *connection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.AutoMigrate(&models.Account{}).Error; err != nil {
		log.Fatal(err)
	}

	// create event dispatcher
	dispatcher, err := events.NewKafkaDispatcher(
		brokerList,
		"accounts",
		events.NewJSONSerializer(),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer dispatcher.Close()

	consumer := events.NewConsumer(
		brokerList,
		[]string{"accounts"},
		"accounts-command",
		events.NewJSONSerializer(
			&ad.AccountCreated{},
			&ad.AccountClosed{},
			&ad.AccountCredited{},
			&ad.AccountDebited{},
		),
		models.NewUpdateHandler(db),
	)
	defer consumer.Stop()

	server := grpc.NewServer()

	svc := command.New(db, dispatcher)
	pb.RegisterAccountsCommandServer(server, svc)

	// initialize server
	s := bedrock.Init(server)

	errChan := make(chan error, 1)
	signalChan := make(chan os.Signal)

	signal.Notify(signalChan, os.Interrupt, os.Kill)

	// start server and other services
	go func() {
		errChan <- s.Run()
	}()

	go func() {
		errChan <- consumer.Start()
	}()

	for {
		select {
		case err := <-errChan:
			if err != nil {
				log.Fatal(err)
			}
		case sig := <-signalChan:
			log.Println(sig)
			s.Stop()
			os.Exit(0)
		}
	}

}
