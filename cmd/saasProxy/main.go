package main

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"runtime/trace"
	"saasProxy/internal/pkg/saasProxy"
	"syscall"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func main() {
	log.Info("saasProxy is starting.")

	// TODO: pass rootCtx and rootCancel around the app
	rootCtx, rootCancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	var exitStatus int

	// Create an instance of your Config struct
	var config saasProxy.Configuration

	config = saasProxy.LoadConfiguration(config)
	router := buildRouter(config)
	server := buildServer(config, router)
	exitStatus = listenAndServe(server, exitStatus)

	select {
	case <-rootCtx.Done():
		log.Info("saasProxy is done!")
		exitStatus = 0
	}

	shutdown(rootCtx, rootCancel)
	log.Exit(exitStatus)
}

func listenAndServe(server *http.Server, exitStatus int) int {
	err := server.ListenAndServe() // Run the http server
	exitStatus = handleListenAndServeError(err, exitStatus)
	return exitStatus
}

func handleListenAndServeError(err error, exitStatus int) int {
	if err != nil {
		exitStatus = 1
		log.WithFields(log.Fields{
			"exitStatus":  exitStatus,
			"err.Error()": err.Error(),
		}).Error("saasProxy encountered an error setting up web server!")
	}
	return exitStatus
}

func buildServer(config saasProxy.Configuration, router *http.ServeMux) *http.Server {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Port),
		Handler: router,
	}
	log.WithFields(log.Fields{
		"port": fmt.Sprintf(":%d", config.Port),
	}).Info("saasProxy is now listening for requests...")
	return server
}

func buildRouter(config saasProxy.Configuration) *http.ServeMux {
	router := config.ToServeMux()
	log.Info("saasProxy handlers have been instantiated!")
	return router
}

func shutdown(ctx context.Context, cancel context.CancelFunc) {
	cancel()
	trace.Stop()
	ctx.Done()
}
