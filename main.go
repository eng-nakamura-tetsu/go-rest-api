package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/eng-nakamura-tetsu/go-rest-api/config"
)

func run(ctx context.Context) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen port %d: %v", cfg.Port, err)
	}
	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("start with: %v", url)

	mux := NewMux()
	s := NewServer(l, mux)

	return s.Run(ctx)
}

func main() {
	if err := run(context.Background()); err != nil {
		fmt.Printf("failed to terminate server: %v", err)
	}
}
