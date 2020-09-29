package client

import (
//	"bytes"
//	"encoding/json"
//	"fmt"
//	"io/ioutil"
//	"net/http"
//	"net/url"
	"log"
  "github.com/d2r2/go-i2c"
  "github.com/d2r2/go-bsbmp"
)

type Sensor struct {
	I2c int
	Model string
}

func (c Sensor) Poll() int {

	i2c, err := i2c.NewI2C(0x76, c.I2c)
	if err != nil {
		log.Fatal(err)
	}
	defer i2c.Close()
	// Make this a case
	sensor, err := bsbmp.NewBMP(bsbmp.BME280, i2c) // signature=0x60


	t, err := sensor.ReadTemperatureC(bsbmp.ACCURACY_STANDARD)
	if err != nil {
		log.Fatal(err)
	}

	return int(t)
}
