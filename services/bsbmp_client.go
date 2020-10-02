package client

import (
	"github.com/d2r2/go-bsbmp"
	"github.com/d2r2/go-i2c"
	logger "github.com/d2r2/go-logger"
	log "github.com/sirupsen/logrus"
)

type Sensor struct {
	I2c     int
	Model   string
	Address uint8
}

type Response struct {
	TemperatureC float32
	PressurePa   float32
	PressureMmHg float32
	HumidityRH   float32
	AltitudeM    float32
}

func (c Sensor) Poll() (*Response, error) {

	resp := Response{}

	i2c, err := i2c.NewI2C(c.Address, c.I2c)
	if err != nil {
		log.Fatal(err)
	}
	defer i2c.Close()

	// todo loglevel flag
	logger.ChangePackageLogLevel("i2c", logger.InfoLevel)
	logger.ChangePackageLogLevel("bsbmp", logger.InfoLevel)

	sensor, err := bsbmp.NewBMP(bsbmp.BMP388, i2c) // No default constructor, placeholder

	switch c.Model {
	case "bmp180":
		sensor, err = bsbmp.NewBMP(bsbmp.BMP180, i2c) // signature=0x55
	case "bme280":
		sensor, err = bsbmp.NewBMP(bsbmp.BME280, i2c) // signature=0x60
	case "bmp280":
		sensor, err = bsbmp.NewBMP(bsbmp.BMP280, i2c) // signature=0x58
	case "bmp388":
		sensor, err = bsbmp.NewBMP(bsbmp.BMP388, i2c) // signature=0x50
	default:
		log.Fatal("No model match!")
	}

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
