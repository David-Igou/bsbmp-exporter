package main

import (
	"net/http"
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	client "github.com/david-igou/bsbmp-exporter/services"
	"github.com/david-igou/bsbmp-exporter/collectors"
)

var (
	listenAddress	= flag.String("port", "9123", "Address to listen on for web interface and telemetry.")
	metricPath	= flag.String("web.telemetry-path", "/metrics", "Path under which to expose metrics.")
	bus		= flag.Int("bus", 1, "Bus of i2c interface (ie 1 for /dev/i2c-1)")
	address		= flag.Uint("address", 0x76, "i2c address of BME/BMP sensor - Verify with i2cdetect -y [bus]")
	model		= flag.String("model", "BME280", "Model of probe [list options]")
)


func main() {
	flag.Parse()

	//Create a new instance of the foocollector and 
	//register it with the prometheus client.
	//case model numbers
	foo := collectors.NewBsbmpCollector(client.Sensor{I2c: *bus, Model: *model})

	prometheus.MustRegister(foo)

	//This section will start the HTTP server and expose
	//any metrics on the /metrics endpoint.
	http.Handle(*metricPath, promhttp.Handler())
	log.Info("Beginning to serve on port ", *listenAddress)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v",*listenAddress), nil))
}
