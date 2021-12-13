package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/torvald2/solanamon/config"
	"github.com/torvald2/solanamon/handlers"
)

func main() {
	conf := config.GetConfig()
	wait := time.Second * 2
	router := handlers.NewRouter()
	srv := &http.Server{
		Addr:    ":" + conf.TCP_Port,
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
	// Listen for os sygnals
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	//Grasefun shutdown
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	os.Exit(0)

}
