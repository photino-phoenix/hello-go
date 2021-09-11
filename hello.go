package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/log"
)

var addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")

var (
	numSayHello = promauto.NewCounter(prometheus.CounterOpts{
		Name: "hello_go_num_say_hello_total",
		Help: "The total number of saying hello",
	})
)

func sayHelloLoop() {
	go func() {
		for {
			numSayHello.Inc()

			var m = &dto.Metric{}
			if err := numSayHello.Write(m); err != nil {
				log.Error(err)
			}
			fmt.Printf("Hello, Prometheus #%d \n", (int)(m.Counter.GetValue()))

			time.Sleep(2 * time.Second)
		}
	}()
}

func main() {
	sayHelloLoop()

	flag.Parse()
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(*addr, nil)
}
