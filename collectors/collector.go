package collectors

import (
//	"log"
        "math/rand"
        "github.com/prometheus/client_golang/prometheus"
        bsbmp "github.com/david-igou/bsbmp-exporter/services"
)

var client bsbmp.Client

//Define a struct for you collector that contains pointers
//to prometheus descriptors for each metric you wish to expose.
//Note you can also include fields of other types if they provide utility
//but we just won't be exposing them as metrics.
type bsbmpCollector struct {
	Temperature *prometheus.Desc
}

//You must create a constructor for you collector that
//initializes every descriptor and returns a pointer to the collector
func NewBsbmpCollector(c bsbmp.Client) *bsbmpCollector {
	return &bsbmpCollector{
		Temperature: prometheus.NewDesc("Temperature",
			"The temperature",
			nil, nil,
		),
	}
}

//Each and every collector must implement the Describe function.
//It essentially writes all descriptors to the prometheus desc channel.
func (collector *bsbmpCollector) Describe(ch chan<- *prometheus.Desc) {

	//Update this section with the each metric you create for a given collector
	ch <- collector.Temperature
}

//Collect implements required collect function for all promehteus collectors
func (collector *bsbmpCollector) Collect(ch chan<- prometheus.Metric) {

	//Implement logic here to determine proper metric value to return to prometheus
	//for each descriptor or call other functions that do so.
	var metricValue float64
	resp := client.Poll() // I want the metrics
	if resp > 1 {
		metricValue = float64(rand.Intn(20) + 10)
	}

	//Write latest value for each metric in the prometheus metric channel.
	//Note that you can pass CounterValue, GaugeValue, or UntypedValue types here.
	ch <- prometheus.MustNewConstMetric(collector.Temperature, prometheus.CounterValue, metricValue)
}
