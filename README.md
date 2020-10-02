# bsbmp-exporter
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/david-igou/bsbmp-exporter/push_latest?style=plastic)
[![Go Report Card](https://goreportcard.com/badge/github.com/david-igou/bsbmp-exporter)](https://goreportcard.com/report/github.com/david-igou/bsbmp-exporter) 
[![MIT License](http://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)



This exporter is made to export readings from Bosch Sensortec BMP180, BMP280, BME280, and BMP388 sensors into a format that can scraped by [Prometheus](https://prometheus.io)

Currently it exports temperature in Celsius, Percent humidity, Pressure in mmHg and Pascal, and Altitude in Meters above sea level.

# Usage

## Command line:

```shell
./bsbmp-exporter -h
Usage of ./bsbmp-exporter:
  -address="0x76": i2c address of BME/BMP sensor - Verify with i2cdetect -y [bus] - Also can use environment variable MODEL
  -bus=1: Bus of i2c interface (ie 1 for /dev/i2c-1). Also can use environment variable BUS
  -metricspath="/metrics": Path under which to expose metrics. Also can use environment variable METRICSPATH 
  -model="BME280": Model of probe - Current supported models: [bmp180, bme280, bmp280, bmp388] - Also can use environment variable MODEL
  -port="9756": Address to listen on for web interface and telemetry. Also can use environment variable PORT
```

## Docker run

```shell
# Currently privileged is required to access /dev/i2c
docker run -p 9756:9756 -v /dev:/dev --privileged quay.io/igou/bsbmp-exporter:latest
```

# Installation

![image](https://raw.github.com/david-igou/bsbmp-exporter/master/docs/bme280-pizero.jpg)

To install the sensors onto a Pi, honestly for each model I just googled "[Model number] Raspberry Pi"

Enable the i2c interface via `raspi-config`

To get the address of the device:

```shell
# ls /dev/i2c* #Note device number
/dev/i2c-1
# i2cdetect -y 1 #Because /dev/i2c-1
     0  1  2  3  4  5  6  7  8  9  a  b  c  d  e  f
00:          -- -- -- -- -- -- -- -- -- -- -- -- -- 
10: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- 
20: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- 
30: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- 
40: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- 
50: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- 
60: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- 
70: -- -- -- -- -- -- 76 --                         
```

This has not (yet!) been tested on Orange/Banana Pis.

# Thanks

This exporter heavily relys on [d2r2/go-i2c](https://github.com/d2r2/go-i2c) and [d2r2/go-bsbmp](https://github.com/d2r2/go-bsbmp) to talk to the i2c bus and wouldn't have been possible for me without them as I am no expert on that subject.
