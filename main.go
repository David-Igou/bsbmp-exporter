package main

import (
	"net/http"
	"github.com/namsral/flag"
	"fmt"
	"strings"
	log "github.com/sirupsen/logrus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	client "github.com/david-igou/bsbmp-exporter/services"
	"github.com/david-igou/bsbmp-exporter/collectors"
)

var (
	port string
	metricPath string
	bus int
	address uint
	model string
)

func main() {
	flag.StringVar(&port, "port", "9123", "Address to listen on for web interface and telemetry.")
	flag.StringVar(&metricPath, "web.telemetry-path", "/metrics", "Path under which to expose metrics.")
	flag.IntVar(&bus, "bus", 1, "Bus of i2c interface (ie 1 for /dev/i2c-1)")
	flag.UintVar(&address, "address", 0x76, "i2c address of BME/BMP sensor - Verify with i2cdetect -y [bus]")
	flag.StringVar(&model, "model", "BME280", "Model of probe [list options]")
	flag.Parse()

	//case model numbers
	foo := collectors.NewBsbmpCollector(client.Sensor{})
	switch strings.ToLower(model) {
	case "bmp180":
		foo = collectors.NewBsbmpCollector(client.Sensor{I2c: bus, Model: "bmp180"})
	case "bmp280":
		foo = collectors.NewBsbmpCollector(client.Sensor{I2c: bus, Model: "bmp280"})
	case "bme280":
		foo = collectors.NewBsbmpCollector(client.Sensor{I2c: bus, Model: "bme280"})
	case "bmp388":
		foo = collectors.NewBsbmpCollector(client.Sensor{I2c: bus, Model: "bmp388"})
	default:
		log.Fatal("Invalid model!")
	}
	prometheus.MustRegister(foo)
	//This section will start the HTTP server and expose
	//any metrics on the /metrics endpoint.
	http.Handle(metricPath, promhttp.Handler())
	log.Info("Beginning to serve on port ", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v",port), nil))
}
