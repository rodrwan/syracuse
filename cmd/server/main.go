package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/jmoiron/sqlx"
	"github.com/rodrwan/syracuse"
	"github.com/rodrwan/syracuse/citizens"
	"github.com/rodrwan/syracuse/postgres"
	"google.golang.org/grpc"

	_ "github.com/lib/pq"
)

func main() {
	port := flag.Int64("port", 8001, "listening port")
	postgresDSN := flag.String("postgres-dsn", "postgres://locahost/sigiriya?sslmode=disable", "Postgres DSN")

	flag.Parse()

	db, err := sqlx.Connect("postgres", *postgresDSN)
	if err != nil {
		log.Fatalln(err)
	}

	srv := grpc.NewServer()

	citizens.RegisterCitizenshipServer(srv, &CitizensService{
		Citizens: &postgres.CitizensService{
			Store: db,
		},
	})

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("Starting Syracuse service...")
	log.Println(fmt.Sprintf("Syracuse service, Listening on: %d", *port))
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

// CitizensService ...
type CitizensService struct {
	Citizens syracuse.Citizens
}

// Get ...
func (cs *CitizensService) Get(ctx context.Context, gr *citizens.GetRequest) (*citizens.GetResponse, error) {
	return nil, nil
}

// Select ...
func (cs *CitizensService) Select(ctx context.Context, gr *citizens.SelectRequest) (*citizens.SelectResponse, error) {
	return nil, nil
}

// Create ...
func (cs *CitizensService) Create(ctx context.Context, gr *citizens.CreateRequest) (*citizens.CreateResponse, error) {
	return nil, nil
}

// Update ...
func (cs *CitizensService) Update(ctx context.Context, gr *citizens.UpdateRequest) (*citizens.UpdateResponse, error) {
	return nil, nil
}

// Delete ...
func (cs *CitizensService) Delete(ctx context.Context, gr *citizens.DeleteRequest) (*citizens.DeleteResponse, error) {
	return nil, nil
}
