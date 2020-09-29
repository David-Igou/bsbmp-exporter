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

type Response struct {
	TemperatureC float32
	PressurePa float32
	PressureMmHg float32
	HumidityRH float32
	AltitudeM float32
}

func (c Sensor) Poll() (*Response, error) {

	resp := Response{}

	i2c, err := i2c.NewI2C(0x76, c.I2c)
	if err != nil {
		log.Fatal(err)
	}
	defer i2c.Close()

	// Make this a case, hard code for now.
	sensor, err := bsbmp.NewBMP(bsbmp.BME280, i2c) // signature=0x60

	// Read Temperature in Celcius
	resp.TemperatureC, err = sensor.ReadTemperatureC(bsbmp.ACCURACY_STANDARD)
	if err != nil {
		log.Fatal(err)
	}

	// Read atmospheric pressure in pascal
	resp.PressurePa, err = sensor.ReadPressurePa(bsbmp.ACCURACY_STANDARD)
	if err != nil {
		log.Fatal(err)
	}

	// Read atmospheric pressure in mmHg
	resp.PressureMmHg, err = sensor.ReadPressureMmHg(bsbmp.ACCURACY_STANDARD)
	if err != nil {
		log.Fatal(err)
	}

	// Read humidity in RH
	_, resp.HumidityRH, err = sensor.ReadHumidityRH(bsbmp.ACCURACY_STANDARD)
	if err != nil {
		log.Fatal(err)
	}

	// Read atmospheric altitude in meters above sea level, if we assume
	// that pressure at see level is equal to 101325 Pa.
	resp.AltitudeM, err = sensor.ReadAltitude(bsbmp.ACCURACY_STANDARD)
	if err != nil {
		log.Fatal(err)
	}

	return &resp, nil
}
