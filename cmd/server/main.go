package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"food-delivery-backend/infra/kafka"
	"food-delivery-backend/infra/postgres"
	"food-delivery-backend/infra/redis"
	"food-delivery-backend/internal/app"
	grpcclient "food-delivery-backend/internal/grpc/client"
	"food-delivery-backend/internal/router"
	"food-delivery-backend/pkg/config"
	"food-delivery-backend/pkg/logger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	log, err := logger.New(cfg.App.LogLevel)
	if err != nil {
		panic(err)
	}
	defer log.Sync()

	db, err := postgres.NewPostgresDB(cfg)
	if err != nil {
		log.Fatal("db")
	}
	_ = postgres.RunMigrations(db.DB, "./migrations")
	rdb, err := redis.NewRedisClient(cfg)
	if err != nil {
		log.Fatal("redis")
	}
	kw, err := kafka.NewProducer(cfg)
	if err != nil {
		log.Fatal("kafka")
	}
	oc, err := grpcclient.NewOrderServiceClient(cfg)
	if err != nil {
		log.Fatal("grpc")
	}

	deps := &app.Container{Config: cfg, Logger: log, DB: db, Redis: rdb, KafkaWriter: kw, OrderClient: oc}
	eng := router.NewRouter(deps)
	srv := &http.Server{Addr: ":" + cfg.App.Port, Handler: eng}
	go func() { _ = srv.ListenAndServe() }()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_ = srv.Shutdown(ctx)
	_ = kw.Close()
	_ = db.Close()
	_ = rdb.Close()
	_ = oc.Close()
}
