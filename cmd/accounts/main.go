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
	"github.com/vectorhacker/bank/internal/pkg/accounts"
	"github.com/vectorhacker/bank/internal/pkg/events"
	domain "github.com/vectorhacker/bank/internal/pkg/events/accounts"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	command "github.com/vectorhacker/bank/internal/pkg/accounts/command"
	pb "github.com/vectorhacker/bank/pb/accounts"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	connection = flag.String("db", os.Getenv("DB_CONNECTION"), "db database connection string")
	brokers    = flag.String("brokers", os.Getenv("KAFKA_PEERS"), "The Kafka brokers to connect to, as a comma separated list")
	consul     = flag.String("consul", os.Getenv("CONSUL_ADDR"), "The consul address to connect to")
	bind       = flag.String("bind", os.Getenv("BIND_ADDR"), "bind to address")
	id         = flag.String("id", os.Getenv("SERVICE_ID"), "The service id")
)

func serviceRegistration(address string, port int, client *api.Client) error {
	registration := &api.AgentServiceRegistration{
		ID:      *id,
		Name:    "accounts.AccountsCommand",
		Tags:    []string{"grpc", "command-side"},
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

	if err := db.AutoMigrate(&accounts.Account{}).Error; err != nil {
		log.Fatal(err)
	}

	// create event dispatcher
	dispatcher, err := events.NewDispatcher(
		brokerList,
		"accounts",
		events.NewJSONSerializer(),
	)
	if err != nil {
		log.Fatal(err)
	}

	// create event consumer
	consumer := events.NewConsumer(
		brokerList,
		[]string{"accounts"},
		"accounts-command",
		events.NewJSONSerializer(
			&domain.AccountCreated{},
			&domain.AccountCredited{},
			&domain.AccountDebited{},
		),
		accounts.NewUpdateHandler(db),
	)

	// register grpc service
	svc := command.New(db, dispatcher)
	pb.RegisterAccountsCommandServer(server, svc)
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
