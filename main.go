package main

import (
	"fmt"
	"log"
	"math/big"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type EthereumClient interface {
	GetBalance(address string) (*big.Int, error)
}

type Client struct {
	URL string
}

var (
	clients []EthereumClient

	httpRequestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "Duration of HTTP requests.",
		Buckets: prometheus.DefBuckets,
	}, []string{"path"})
	balanceRetrievalDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "balance_retrieval_duration_seconds",
		Help:    "Duration of balance retrieval operations.",
		Buckets: prometheus.DefBuckets,
	})
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	startTime := time.Now()

	r.Use(middleware.Heartbeat("/"))

	r.Post("/health", func(w http.ResponseWriter, r *http.Request) {
		duration := time.Since(startTime)
		if duration.Seconds() > 10 {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			return
		}
	})

	// Prometheus Metrics
	r.Handle("/metrics", promhttp.Handler())

	// Load Clients
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("unable to load configs: %e", err)
		return
	}

	clients, err = initClients(config)
	if err != nil {
		log.Fatal(err)
		return
	}

	r.Get("/eth/balance/{address}", prometheusMiddleware(getBalance))

	fmt.Println("Serving at :3333")
	http.ListenAndServe(":3333", r)
}

func prometheusMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		timer := prometheus.NewTimer(httpRequestDuration.WithLabelValues(r.URL.Path))
		defer timer.ObserveDuration()

		next(w, r)
	}
}

func getBalance(w http.ResponseWriter, r *http.Request) {
	userAddress := chi.URLParam(r, "address")

	var (
		balance *big.Int
		err     error
	)

	startTime := time.Now()

	fmt.Println("userAddress: ", userAddress)

	for _, client := range clients {
		balance, err = client.GetBalance(userAddress)
		if err != nil {
			log.Printf("unable to get balance: %e", err)
			continue
		}

		fmt.Println("balance: ", balance)
		break
	}

	balanceRetrievalDuration.Observe(time.Since(startTime).Seconds())

	w.Write([]byte("Balance: " + balance.String()))

}
