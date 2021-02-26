package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/weekndCN/rw-app/handler"
	"github.com/weekndCN/rw-app/store/auth"
	"github.com/weekndCN/rw-app/store/dbtest"
	"github.com/weekndCN/rw-app/store/file"
)

// graceWait is the time duration
// which server wait for existing connections to finish
var graceWait time.Duration

func main() {
	// inital log
	initLogging()
	// server set up
	flag.DurationVar(&graceWait, "graceful-timeout", time.Second*5, "server wait for existing connections to finish")
	flag.Parse()
	// test database connect
	db, _ := dbtest.Open()
	svc := handler.New(auth.New(db), file.New(db))

	//Http server
	server := &http.Server{
		Addr:         ":9090",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      handler.Handler(svc),
	}

	logrus.WithFields(
		logrus.Fields{
			"Addr":         server.Addr,
			"WriteTimeout": server.WriteTimeout,
			"ReadTimeout":  server.ReadTimeout,
			"IdleTimeout":  server.IdleTimeout,
		},
	).Infoln("starting the http server")

	// run server in a goroutine, make it unblock
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// close channel
	closeC := make(chan os.Signal, 1)

	// only work via control+ C
	signal.Notify(closeC, os.Interrupt)

	// block until server receive signal
	<-closeC

	// graceful shutdown with a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), graceWait)
	defer cancel()
	// if no connection, it will be unblock,
	// otherwise will wait util the  timeout deadline(graceWait time.Duration)
	server.Shutdown(ctx)
	log.Println("server is shutting down")
	os.Exit(0)
}

func initLogging() {
	level := os.Getenv("RWPLUS_LOG_LEVEL")
	switch level {
	case "Debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "Trace":
		logrus.SetLevel(logrus.TraceLevel)
	case "Text":
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors:   true,
			DisableColors: false,
		})
	default:
		logrus.SetFormatter(&logrus.JSONFormatter{
			PrettyPrint: true,
		})
	}
}
