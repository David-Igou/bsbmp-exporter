package main

import (
  "net/http"

  log "github.com/Sirupsen/logrus"
  "github.com/prometheus/client_golang/prometheus"
  "github.com/prometheus/client_golang/prometheus/promhttp"

  client "github.com/david-igou/bsbmp-exporter/services"
  "github.com/david-igou/bsbmp-exporter/collectors"

)

func main() {
// add flags:
// i2c address
// bmp models
// port

//  i2c, ierr := i2c.NewI2C(0x76, 1)
//  if err != nil {
//	lg.Fatal(err)
// }
//  defer i2c.Close()

 // TODO Make this a case

  // sensor, err := bsbmp.NewBMP(bsbmp.BMP180, i2c) // signature=0x55
  //sensor, err := bsbmp.NewBMP(bsbmp.BMP280, i2c) // signature=0x58
//  sensor, err := bsbmp.NewBMP(bsbmp.BME280, i2c) // signature=0x60
  // sensor, err := bsbmp.NewBMP(bsbmp.BMP388, i2c) // signature=0x50



  //Create a new instance of the foocollector and 
  //register it with the prometheus client.
  foo := collectors.NewBsbmpCollector(client.Sensor{I2c: 1, Model: "BME280"})
  prometheus.MustRegister(foo)

  //This section will start the HTTP server and expose
  //any metrics on the /metrics endpoint.
  http.Handle("/metrics", promhttp.Handler())
  log.Info("Beginning to serve on port :8080")
  log.Fatal(http.ListenAndServe(":8080", nil))


}
