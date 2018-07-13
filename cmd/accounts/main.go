package main

import (
	"context"
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
	"github.com/vectorhacker/bank/internal/pkg/events/domain"

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

	config := api.DefaultConfig()
	config.Address = *consul
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	service := &api.AgentServiceRegistration{
		ID:      *id,
		Name:    "accounts.AccountsCommand",
		Tags:    []string{"grpc", "command-side"},
		Address: addr[0],
		Port:    port,
	}

	err = client.Agent().ServiceRegister(service)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Agent().ServiceDeregister(*id)

	brokerList := strings.Split(*brokers, ",")
	log.Printf("Kafka brokers: %s", strings.Join(brokerList, ", "))

	server := grpc.NewServer()

	db, err := gorm.Open("postgres", *connection)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err := db.AutoMigrate(&accounts.Account{}).Error; err != nil {
		log.Fatal(err)
	}

	dispatcher, err := events.NewDispatcher(
		brokerList,
		"accounts",
		events.NewJSONSerializer(),
	)
	if err != nil {
		log.Fatal(err)
	}

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

	svc := command.New(db, dispatcher)
	pb.RegisterAccountsCommandServer(server, svc)
	reflection.Register(server)

	lis, err := net.Listen("tcp", *bind)
	if err != nil {
		log.Fatal(err)
	}

	errChan := make(chan error)
	signalChan := make(chan os.Signal)

	signal.Notify(signalChan, os.Kill, os.Interrupt)

	go func() {
		errChan <- server.Serve(lis)
	}()

	ctx := context.Background()
	ctx, close := context.WithCancel(ctx)
	defer close()

	go func() {
		errChan <- consumer.Start(ctx)
	}()

	select {
	case s := <-signalChan:
		log.Println("exiting", s)

		server.GracefulStop()
	case err := <-errChan:
		if err != nil {
			log.Fatal(err)
		}
	}
}
