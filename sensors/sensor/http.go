package sensor

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

var (
	s Sensor
)

type SenData struct {
	T string
	V string
}

type Ret struct {
	Code int
	Data interface{}
}

func Upload(c echo.Context) error {
	d := new(SenData)
	if err := c.Bind(d); err != nil {
		return err
	}
	d.T += fmt.Sprintf("-%s", s.ID())
	return s.ForwardData(*d)
}

func HttpCharge(c echo.Context) error {
	s.ChargeBattery()
	return nil
}

func StartSensor(id, ip string, port string, sink, dataset string, interv, dura int, x, y, z int) {
	self := Node{
		ID:   id,
		Addr: ip,
		Port: port,
		Loc: Location{
			X: x,
			Y: y,
			Z: z,
		},
	}
	//Create new sensor
	s = NewRealCensor(self, sink, dataset, interv, dura)
	err := s.Register()
	if err != nil {
		log.Fatalf("Register failed:%v", err)
	}
	//start ducy cycle for the sensor
	s.StartDutyCycle()
	//starting generate data
	err = s.GenerateData()
	if err != nil {
		log.Fatalf("load dataset failed:%v", err)
	}
	e := echo.New()
	e.HideBanner = true
	e.POST("/data/upload", Upload)
	e.GET("/sensor/charge", HttpCharge)
	listenaddr := fmt.Sprintf(":%s", port)
	log.Info("listen:", listenaddr)
	e.Logger.Fatal(e.Start(listenaddr))
}
