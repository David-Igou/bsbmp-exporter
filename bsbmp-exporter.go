package main

import (
	"fmt"
	"github.com/david-igou/bsbmp-exporter/collectors"
	client "github.com/david-igou/bsbmp-exporter/services"
	"github.com/namsral/flag"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
)

var (
	port       string
	metricPath string
	bus        int
	address    string
	model      string
)

func main() {
	flag.StringVar(&port, "port", "9756", "Address to listen on for web interface and telemetry. Also can use environment variable PORT")
	flag.StringVar(&metricPath, "metricspath", "/metrics", "Path under which to expose metrics. Also can use environment variable METRICSPATH ")
	flag.IntVar(&bus, "bus", 1, "Bus of i2c interface (ie 1 for /dev/i2c-1). Also can use environment variable BUS")
	flag.StringVar(&address, "address", "0x76", "i2c address of BME/BMP sensor - Verify with i2cdetect -y [bus] - Also can use environment variable ADDRESS")
	flag.StringVar(&model, "model", "BME280", "Model of probe - Current supported models: [bmp180, bme280, bmp280, bmp388] - Also can use environment variable MODEL")
	flag.Parse()

	// Clean up address input, since not everyone prefixes with 0x/x
	if address[0:2] == "0x" {
		address = address[2:]
	} else if address[0:1] == "x" {
		address = address[1:]
	}

	address64, _ := strconv.ParseUint(address, 16, 8)
	address8 := uint8(address64)

	log.Info("Using address ", address8)

	//case model numbers
	foo := collectors.NewBsbmpCollector(client.Sensor{})
	switch strings.ToLower(model) {
	case "bmp180":
		foo = collectors.NewBsbmpCollector(client.Sensor{Address: address8, I2c: bus, Model: "bmp180"})
	case "bmp280":
		foo = collectors.NewBsbmpCollector(client.Sensor{Address: address8, I2c: bus, Model: "bmp280"})
	case "bme280":
		foo = collectors.NewBsbmpCollector(client.Sensor{Address: address8, I2c: bus, Model: "bme280"})
	case "bmp388":
		foo = collectors.NewBsbmpCollector(client.Sensor{Address: address8, I2c: bus, Model: "bmp388"})
	default:
		log.Fatal("Invalid model!")
	}
	prometheus.MustRegister(foo)
	//This section will start the HTTP server and expose
	//any metrics on the /metrics endpoint.
	http.Handle(metricPath, promhttp.Handler())
	log.Info("Beginning to serve on port ", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
