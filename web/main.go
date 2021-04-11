package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/luanpontolio/restaurant-kitchen-go/core/restaurant"
)

const dbsource = "data/restaurant.db"

func takeCook(ctx context.Context, service restaurant.RestaurantService, score int64) *restaurant.Cook {
	cook, err := service.GetCookByScore(ctx, score)
	if err != nil {
		fmt.Printf("failed GetCookByScore: %v", err)
		return nil
	}

	return cook
}

func makeOrder(ctx context.Context, service restaurant.RestaurantService, cook *restaurant.Cook, order *restaurant.Order) {
	// start to preparer the order
	service.UpdateCook(ctx, cook.ID.String(), cook.Score, 1)
	service.UpdateOrder(ctx, order.ID.String(), order.Plate, order.Score, 1)
	// time to cook
	time.Sleep(time.Duration(order.Score) * time.Millisecond)
	// finished the order
	service.UpdateCook(ctx, cook.ID.String(), cook.Score, 0)
	service.UpdateOrder(ctx, order.ID.String(), order.Plate, order.Score, 2)
}

func deliveryOrder(ctx context.Context, service restaurant.RestaurantService, order *restaurant.Order) {
	service.UpdateOrderHash(ctx, order.ID.String())
}

func runRestaurantWorker(ctx context.Context, service restaurant.RestaurantService, log log.Logger) {
	level.Info(log).Log("msg", "restaurant start")

	orders, err := service.GetAllOrder(ctx, 0, false)
	if err != nil {
		level.Error(log).Log("error", "something wrong on worker %v", err)
	}

	defer level.Info(log).Log("msg", "processed orders...")
	for i := 0; i < len(orders); i++ {
		go func(index int) {
			cook := takeCook(ctx, service, orders[index].Score)
			if cook == nil {
				level.Info(log).Log("not found", "not found the cook to make a order: %v", orders[index].ID)
				return
			}

			makeOrder(ctx, service, cook, orders[index])
			deliveryOrder(ctx, service, orders[index])
		}(i)
	}
}

func main() {
	var httpAddr = flag.String("http", ":5000", "http listen address")
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "restaurant",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	var db *sql.DB
	{
		var err error

		db, err = sql.Open("sqlite3", dbsource)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}

	}

	flag.Parse()
	ctx := context.Background()
	var srv restaurant.RestaurantService
	{
		repository := restaurant.NewRepo(db, logger)

		srv = restaurant.NewService(repository, logger)
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoints := restaurant.MakeEndpoints(srv)

	go func() {
		fmt.Println("listening on port", *httpAddr)
		handler := restaurant.NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	go runRestaurantWorker(ctx, srv, logger)

	level.Error(logger).Log("exit", <-errs)
}
