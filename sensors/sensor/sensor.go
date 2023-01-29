package sensor

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync/atomic"
	"time"

	"github.com/labstack/gommon/log"
)

type Sensor interface {
	ForwardData(d SenData) error
	GenerateData() error
	StartDutyCycle()
	Register() error
	ID() string
	ReduceBattery(v int32) error
	ChargeBattery() error
	GetBattery() int32
}

type Location struct {
	X int
	Y int
	Z int
}

type Node struct {
	ID   string
	Addr string
	Port string
	Loc  Location
}

type RealSensor struct {
	self        Node
	Sink        string
	Interval    int
	Duration    int
	CloserNodes []Node
	DataSet     string
	alive       bool
	battery     *int32
}

var (
	httpclient          = createHTTPClient()
	ErrSleep            = errors.New("sensor is not alive")
	ErrForward          = errors.New("data forwarding failed")
	ErrNoBattery        = errors.New("Battery Exhausted")
	ErrLowBattery       = errors.New("Battery Low")
	ForwardCost   int32 = -2
	GenerateCost  int32 = -1
)

func createHTTPClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 20,
		},
		Timeout: time.Duration(3) * time.Second,
	}

	return client
}

// NewRealCensor Create new Sensor Instance
func NewRealCensor(self Node, sink, dataset string, interv int, dur int) *RealSensor {
	res := &RealSensor{
		self:     self,
		Sink:     sink,
		Interval: interv,
		Duration: dur,
		DataSet:  dataset,
		alive:    true,
		battery:  new(int32),
	}
	res.ChargeBattery()
	return res
}

func (s *RealSensor) ID() string {
	return s.self.ID
}

// ChargeBattery change battery level to 100
func (s *RealSensor) ChargeBattery() error {
	atomic.StoreInt32(s.battery, 100)
	if s.GetBattery() <= 0 && s.alive == false {
		s.alive = true
	}
	log.Info("battery charged")
	return nil
}

func (s *RealSensor) ReduceBattery(v int32) error {
	if atomic.LoadInt32(s.battery) <= 0 {
		return ErrNoBattery
	}
	atomic.AddInt32(s.battery, v)
	if s.GetBattery() < 10 {
		log.Warnf("battery level:%v", s.GetBattery())
	}
	return nil
}

func (s *RealSensor) GetBattery() int32 {
	return atomic.LoadInt32(s.battery)
}

// Register can register the sensor to the sink
func (s *RealSensor) Register() error {
	payload, _ := json.Marshal(s.self)
	req, err := http.NewRequest("POST", fmt.Sprintf("http://%s/node/register", s.Sink), bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := httpclient.Do(req)
	if err != nil {
		return err
	}

	// the sink will return all the sensors that are closer to the sink, the list will be save to CloserNodes
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	nodes := make([]Node, 0)
	err = json.Unmarshal(body, &nodes)
	if err != nil {
		return err
	}
	s.CloserNodes = nodes
	log.Infof("found closer sensor:%v", s.CloserNodes)
	return nil
}

// ForwardData can send data to the closer sensor
func (s *RealSensor) ForwardData(d SenData) error {
	if !s.alive {
		return ErrSleep
	}

	if s.GetBattery() < 10 {
		log.Warnf("battery level:%v, refuse to forward data", s.GetBattery())
		return ErrLowBattery
	}
	s.ReduceBattery(ForwardCost)
	//origint := d.T
	var err error
	// keep trying until success
	for _, v := range s.CloserNodes {
		log.Infof("forward data:%v to %v", d, v)
		//d.T = origint + "->" + v.ID
		err = s.senddata(d, fmt.Sprintf("%s:%s", v.Addr, v.Port))
		if err == nil {
			log.Info("forward success")
			return nil
		} else {
			log.Errorf("forward failed:%v", err)
		}
	}
	//the sensor will just try to send the data to the sink if there are no closer sensors
	log.Infof("no forward node avaliable, send data:%v to the sink:%v", d, s.Sink)
	err = s.senddata(d, s.Sink)
	if err != nil {
		log.Errorf("send data to sink failed:%v", err)
	} else {
		log.Infof("send data to sink successfully")
	}
	return nil
}

// senddata sendata to another sensor or the sink using http
func (s *RealSensor) senddata(d SenData, addr string) error {
	payload, _ := json.Marshal(d)
	req, err := http.NewRequest("POST", fmt.Sprintf("http://%s/data/upload", addr), bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := httpclient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return ErrForward
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Infof("send data result:%v", string(body))
	return nil
}

// GenerateData generate data from provided dataset files
func (s *RealSensor) GenerateData() error {
	dat, err := ioutil.ReadFile(s.DataSet)
	if err != nil {
		return err
	}
	dataset := strings.Split(string(dat), "\n")
	idx := -1
	go func() {
		for true {
			time.Sleep(time.Second)
			idx++
			s.ReduceBattery(GenerateCost)
			//TODO: add critial logic
			for !s.alive {
				time.Sleep(time.Second)
			}
			dvalue := ""
			//for dvalue == "" {
			dvalue = strings.Trim(dataset[idx%len(dataset)], "\t \n")
			//	idx++
			//}
			decodedata := make(map[string]interface{})
			err = json.Unmarshal([]byte(dvalue), &decodedata)
			if err != nil {
				log.Errorf("GenerateData:fail to docode data:%v", dvalue)
				continue
			}
			//only send the data when the data is not normal
			if decodedata["Condition"].(string) == "Normal" {
				log.Infof("normal found, no sending:%v", dvalue)
				continue
			}
			report := SenData{T: s.self.ID, V: dvalue}
			log.Infof("report:%v", report)
			s.ForwardData(report)
		}
	}()
	return nil
}

//StartDucyCycle start ducy cycle for the sensor
func (s *RealSensor) StartDutyCycle() {
	s.alive = true
	duration := time.Duration(s.Duration) * time.Second
	interval := time.Duration(s.Interval) * time.Second

	go func() {
		for true {
			time.Sleep(duration)
			s.alive = false
			log.Warn("sensor goes to sleep")
			time.Sleep(interval)
			// the sensor will keep sleeping if the battery level is 0
			if s.GetBattery() > 0 {
				s.alive = true
				log.Warn("sensor wakes up")
			}
		}
	}()
}
