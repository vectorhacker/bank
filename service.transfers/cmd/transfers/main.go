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

	accountspb "github.com/vectorhacker/bank/service.accounts/pb"
	"github.com/vectorhacker/bank/service.transfers/internal/pkg/transfers"
	pb "github.com/vectorhacker/bank/service.transfers/pb"
	td "github.com/vectorhacker/bank/service.transfers/pkg/events"
	"google.golang.org/grpc"
)

var (
	connection = flag.String("db", os.Getenv("DB_CONNECTION"), "db database connection string")
	brokers    = flag.String("brokers", os.Getenv("KAFKA_PEERS"), "The Kafka brokers to connect to, as a comma separated list")
	dial       = flag.String("grpc", os.Getenv("GRPC_DIAL"), "The grpc connection string")
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

	if err := db.AutoMigrate(&transfers.Transfer{}).Error; err != nil {
		log.Fatal(err)
	}

	cc, err := grpc.Dial(*dial, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := accountspb.NewAccountsCommandClient(cc)

	// create event dispatcher
	dispatcher, err := events.NewKafkaDispatcher(
		brokerList,
		"transfers",
		events.NewJSONSerializer(),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer dispatcher.Close()

	consumer := events.NewConsumer(
		brokerList,
		[]string{"transfers"},
		"transfers-saga",
		events.NewJSONSerializer(
			&td.TransferBegun{},
			&td.TransferCompleted{},
			&td.TransferCreditAccountBegun{},
			&td.TransferCreditCompleted{},
			&td.TransferCreditFailed{},
			&td.TransferDebitAccountBegun{},
			&td.TransferDebitCompleted{},
			&td.TransferDebitFailed{},
			&td.TransferFailed{},
		),
		transfers.NewExecutor(db, dispatcher, client),
	)
	defer consumer.Stop()

	server := grpc.NewServer()

	svc := transfers.NewService(dispatcher)
	pb.RegisterTransfersServer(server, svc)

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
