package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"sync/atomic"
	"time"
)

type key int

const (
	requestIDKey key = 0
)

var (
	// default version value
	version = "0.0"

	// Whether or not to print the current version of the program
	printVersion = false

	// listen address
	listenAddr string

	// the directory to read from; e.g. /path/to/web/server
	serverPath string

	// health status
	healthy int32
)

func init() {

	flag.BoolVar(&printVersion, "version", false,
		"Print the current version of this program and exit.")

	flag.StringVar(&listenAddr, "listen-addr", "5000",
		"Port for the server to listen on;")

	flag.StringVar(&serverPath, "server-path", "/var/www/html/",
		"Directory to search for website HTML files.")
}

//
// Program main
//
func main() {

	flag.Parse()

	// if requested, go ahead and print the version; afterwards exit the
	// program, since this is all done
	if printVersion {
		fmt.Println("libreserver v" + version)
		os.Exit(0)
	}

	// safety check the requested port
	portRegex := regexp.MustCompile("^[0-9]+$")
	if portRegex.FindString(listenAddr) == "" {
		fmt.Println("Invalid port requested... " + listenAddr)
		os.Exit(1)
	}

	// safety check the server path
	// TODO: use this variable somewhere
	if serverPath == "" {
		fmt.Println("Invalid server path requested!")
		os.Exit(1)
	}

	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Println("Server is starting...")

	router := http.NewServeMux()
	router.Handle("/", index())
	router.Handle("/metrics", index())
	router.Handle("/healthz", healthz())

	nextRequestID := func() string {
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}

	server := &http.Server{
		Addr:         ":" + listenAddr,
		Handler:      tracing(nextRequestID)(logging(logger)(router)),
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		logger.Println("Server is shutting down...")
		atomic.StoreInt32(&healthy, 0)

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
		}
		close(done)
	}()

	logger.Println("Server is ready to handle requests at", listenAddr)
	atomic.StoreInt32(&healthy, 1)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %s: %v\n", listenAddr, err)
	}

	<-done
	logger.Println("Server stopped")
}
