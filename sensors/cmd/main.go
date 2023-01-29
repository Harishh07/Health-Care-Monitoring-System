package main

import (
	"flag"
	"sensors/sensor"
)

var (
	id      *string
	addr    *string
	port    *string
	sink    *string
	x       *int
	y       *int
	z       *int
	interv  *int
	dura    *int
	dataset *string
)

func init() {
	id = flag.String("id", "1", "sensor id")
	addr = flag.String("addr", "127.0.0.1", "listen ip address")
	port = flag.String("port", "1234", "listen port")
	sink = flag.String("sink", "127.0.0.1:9092", "sink address")
	dataset = flag.String("dataset", "data.txt", "dataset file")

	interv = flag.Int("interval", 10, "interval between working(seconds)")
	dura = flag.Int("duration", 5, "working duration(seconds)")
	x = flag.Int("x", 1, "coordinate x")
	y = flag.Int("y", 1, "coordinate y")
	z = flag.Int("z", 1, "coordinate z")
}

func main() {
	flag.Parse()
	sensor.StartSensor(*id, *addr, *port, *sink, *dataset, *interv, *dura, *x, *y, *z)
}
