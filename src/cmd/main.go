package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"court.com/src/internal/controller"
	"court.com/src/internal/repo"
	"court.com/src/internal/service"
	"court.com/src/pkg/mongodb"
	"court.com/src/pkg/redis"
	config "court.com/src/pkg/utils/conf"
	"github.com/gorilla/mux"
)

func main() {

	// init config
	conf, err := config.Init()
	if err != nil {
		log.Fatal("Error config", err)
	}

	// init pkg
	cli, err := mongodb.InitClient(conf.Repo.MongoDB)
	if err != nil {
		log.Fatal(err)
	}
	redisCli := redis.Init(conf.Repo.Redis)

	// init repo
	repo := repo.New(cli, redisCli)

	// init service
	placeSvc := service.New(repo)
	locSvc := service.NewLocationSvc(repo)

	ctrl := controller.Init(placeSvc, locSvc)

	// init router
	m := mux.NewRouter()
	m.HandleFunc("/place/{id}", ctrl.GetPlaceByIdHandler)
	m.HandleFunc("/gymnasium/{id}", ctrl.GetGymnasiumByID)

	Run(conf, m)
}

func Run(conf *config.Config, r *mux.Router) {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	srv := &http.Server{
		Addr: fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
