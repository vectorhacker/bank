package bedrock

import (
	"encoding/base64"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	"crypto/rand"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

const (
	// ServerAddrEnvName defines an environment variable name which sets
	// the full IP:Port pair for the grpc server
	ServerAddrEnvName = "SERVER_ADDR"

	// ServerPortEnvName defines an environment variable name which sets
	// the server Port for the grpc server
	ServerPortEnvName = "SERVER_PORT"

	// ServerIPEnvName defines an environment variable name which sets
	// the server IP address for the grpc server
	ServerIPEnvName = "SERVER_IP"

	// ServiceIDEnvName defines an environment variable name which sets
	// the id for services located in this instance
	ServiceIDEnvName = "SERVICE_ID"

	// ServiceTagsEnvName defines an environment variable name which sets
	// the tags for servcies located in this instance
	ServiceTagsEnvName = "SERIVCE_TAGS"

	ServiceGRPCUseTLSEnvName = "SERVICE_USE_TLS"
)

type Server struct {
	server *grpc.Server
	stop   chan struct{}
}

// Init initializes the server
func Init(server *grpc.Server) *Server {

	return &Server{
		server: server,
	}
}

// Stop preforms a graceful stop
func (s *Server) Stop() {
	s.stop <- struct{}{}
}

// Run registers and runs the server
func (s *Server) Run() error {

	config := api.DefaultConfig()

	client, err := api.NewClient(config)
	if err != nil {
		return err
	}

	// generate random id
	c := 10
	b := make([]byte, c)
	_, err = rand.Read(b)
	if err != nil {
		return err
	}

	id := base64.RawURLEncoding.EncodeToString(b)

	registration := api.AgentServiceRegistration{
		Address: "",
		Port:    5000,
		Tags:    []string{"grpc"},
		ID:      id,
		Check: &api.AgentServiceCheck{
			CheckID:  id,
			Interval: "10s",
			// GRPCUseTLS: false,
		},
	}
	if addr := os.Getenv(ServerAddrEnvName); addr != "" {
		ipport := strings.Split(addr, ":")

		ip := ipport[0]
		port, err := strconv.Atoi(ipport[1])
		if err != nil {
			return err
		}

		registration.Address = ip
		registration.Port = port
	}

	if ip := os.Getenv(ServerIPEnvName); ip != "" {
		registration.Address = ip
	}

	if p := os.Getenv(ServerPortEnvName); p != "" {
		port, err := strconv.Atoi(p)
		if err != nil {
			return err
		}
		registration.Port = port
	}

	if id := os.Getenv(ServiceIDEnvName); id != "" {
		registration.ID = id
	}

	if t := os.Getenv(ServiceTagsEnvName); t != "" {
		tags := strings.Split(t, ",")

		registration.Tags = append(registration.Tags, tags...)
	}

	if use := os.Getenv(ServiceGRPCUseTLSEnvName); use == "true" {
		registration.Check.GRPCUseTLS = true
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", registration.Address, registration.Port))
	if err != nil {
		return err
	}

	registration.Check.TCP = lis.Addr().String()
	for service := range s.server.GetServiceInfo() {
		registration.Name = service

		if err := client.Agent().ServiceRegister(&registration); err != nil {
			return err
		}
	}

	errChan := make(chan error, 1)

	go func() {
		errChan <- s.server.Serve(lis)
	}()

	// for some reason, this doesn't really work
	defer client.Agent().ServiceDeregister(registration.ID)

	// start server
	select {
	case err := <-errChan:
		if err != nil {
			return err
		}
		return nil
	case <-s.stop:
		s.server.GracefulStop()
		return nil
	}
}
