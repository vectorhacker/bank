package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"

	"github.com/hashicorp/consul/api"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/vectorhacker/bank/internal/pkg/events"
	td "github.com/vectorhacker/bank/internal/pkg/events/transfers"
	"github.com/vectorhacker/bank/internal/pkg/transfers"
	accountsPb "github.com/vectorhacker/bank/pb/accounts"
	pb "github.com/vectorhacker/bank/pb/transfers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	connection = flag.String("db", os.Getenv("DB_CONNECTION"), "db database connection string")
	brokers    = flag.String("brokers", os.Getenv("KAFKA_PEERS"), "The Kafka brokers to connect to, as a comma separated list")
	consul     = flag.String("consul", os.Getenv("CONSUL_ADDR"), "The consul address to connect to")
	bind       = flag.String("bind", os.Getenv("BIND_ADDR"), "bind to address")
	id         = flag.String("id", os.Getenv("SERVICE_ID"), "The service id")
	grpcDial   = flag.String("dial", os.Getenv("GRPC_DIAL"), "The grpc service endpoint")
)

func serviceRegistration(address string, port int, client *api.Client) error {
	registration := &api.AgentServiceRegistration{
		ID:      *id,
		Name:    "transfers.Transfers",
		Tags:    []string{"grpc", "sagas"},
		Address: address,
		Port:    port,
	}

	return client.Agent().ServiceRegister(registration)
}

func deregisterService(client *api.Client) {
	if err := client.Agent().ServiceDeregister(*id); err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()

	if *connection == "" {
		log.Fatal("expected database connection")
	}

	if *brokers == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *consul == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	addr := strings.Split(*bind, ":")
	port, err := strconv.Atoi(addr[1])
	if err != nil {
		log.Fatal(err)
	}

	brokerList := strings.Split(*brokers, ",")
	log.Printf("Kafka brokers: %s", strings.Join(brokerList, ", "))

	config := api.DefaultConfig()
	config.Address = *consul
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
	if err := serviceRegistration(addr[0], port, client); err != nil {
		log.Fatal(err)
	}
	defer deregisterService(client)

	server := grpc.NewServer()

	db, err := gorm.Open("postgres", *connection)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err := db.AutoMigrate(&transfers.Transfer{}).Error; err != nil {
		log.Fatal(err)
	}

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

	cc, err := grpc.Dial(*grpcDial, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	accounts := accountsPb.NewAccountsCommandClient(cc)

	// create event consumer
	consumer := events.NewConsumer(
		brokerList,
		[]string{"transfers"},
		"transfers",
		events.NewJSONSerializer(
			&td.TransferCreditCompleted{},
			&td.TransferDebitCompleted{},
			&td.TransferBegun{},
			&td.TransferCompleted{},
			&td.TransferCreditAccountBegun{},
			&td.TransferCreditFailed{},
			&td.TransferDebitAccountBegun{},
			&td.TransferDebitFailed{},
		),
		transfers.NewExecutor(db, dispatcher, accounts),
	)

	// register grpc service
	svc := transfers.NewService(dispatcher)
	pb.RegisterTransfersServer(server, svc)
	reflection.Register(server)

	errChan := make(chan error, 10)
	signalChan := make(chan os.Signal)

	signal.Notify(signalChan, os.Kill, os.Interrupt)

	// Start grpc server
	go func() {
		lis, err := net.Listen("tcp", *bind)
		if err != nil {
			log.Fatal(err)
		}

		defer server.GracefulStop()
		errChan <- server.Serve(lis)
	}()

	// Start kafka update consumer
	go func() {
		defer consumer.Stop()
		errChan <- consumer.Start()
	}()

	// listen for errors or signal
	select {
	case s := <-signalChan:
		// handle graceful exit
		log.Println("exiting", s)
		os.Exit(0)
	case err := <-errChan:
		if err != nil {
			log.Fatal(err)
		}
	}

}
