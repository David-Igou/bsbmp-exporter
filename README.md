# bsbmp-exporter

This is in its POC stages

Exports Temperature, Pressure (Pa and MmHg), relative humidity, and altitude.


# Running 


```shell
$ ./bsbmp-exporter -h
Usage of ./bsbmp-exporter:
  -address uint
    	i2c address of BME/BMP sensor - Verify with i2cdetect -y [bus] (default 118)
  -bus int
    	Bus of i2c interface (ie 1 for /dev/i2c-1) (default 1)
  -model string
    	Model of probe [list options] (default "BME280")
  -port string
    	Address to listen on for web interface and telemetry. (default "9123")
  -web.telemetry-path string
    	Path under which to expose metrics. (default "/metrics")
```


