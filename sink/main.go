package main

import (
	"encoding/json"
	"net"
	"net/http"
	"sync"

	"flag"

	"github.com/fernet/fernet-go"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	//"github.com/labstack/echo"
	//"github.com/fernet/fernet-go"
	//"github.com/labstack/gommon/log"
)

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

type Nodes struct {
	nodes []*Node
	sync.Mutex
}

var (
	nodeman  *Nodes
	datapool DataPool
	pdaAddr  *string
	laddr    *string
	enckey   = "f-Fd2vxmXx_3xO6Y1vdZ0eCnRN0ocBJ_cfpxC8W2JII="
)

func main() {
	// Echo instance
	pdaAddr = flag.String("pda", "127.0.0.1:9091", "personal digital assistant address")
	laddr = flag.String("local", "0.0.0.0:9092", "local addr")
	flag.Parse()
	nodeman = new(Nodes)
	e := echo.New()
	datapool.Data = make([]*SenData, 0)

	e.POST("/node/register", register)
	e.POST("/data/upload", upload)
	e.GET("/battery/charge", charge)

	// Start server
	e.Logger.Fatal(e.Start(*laddr))
}

type SenData struct {
	T string
	V string
}

type DataPool struct {
	Data []*SenData
	sync.Mutex
}

type Ret struct {
	Code int
	Data interface{}
}

func upload(c echo.Context) error {
	d := new(SenData)
	if err := c.Bind(d); err != nil {
		return err
	}
	datapool.Lock()
	defer datapool.Unlock()
	datapool.Data = append(datapool.Data, d)
	log.Infof("data path:%v", d.T)
	log.Infof("receive data:%v", d)
	return nil
}

func charge(c echo.Context) error {
	conn, err := net.Dial("tcp", *pdaAddr)
	if err != nil {
		log.Errorf("Send data error:%v", err)
		return err
	}
	defer conn.Close()
	datapool.Lock()
	olddata := datapool.Data
	datapool.Data = make([]*SenData, 0)
	datapool.Unlock()
	payload, err := json.Marshal(olddata)
	if err != nil {
		log.Errorf("charge json encode:%v", err)
		return err
	}
	k := fernet.MustDecodeKeys(enckey)
	tok, err := fernet.EncryptAndSign(payload, k[0])
	if err != nil {
		log.Errorf("encrypt failed:%v", err)
		return err
	}

	_, err = conn.Write(tok)
	//_, err = conn.Write(payload)
	if err != nil {
		log.Errorf("sending data failed:%v", err)
	}
	log.Info("sending data successfully")
	return nil
}

// Handler
func register(c echo.Context) error {
	req := new(Node)
	if err := c.Bind(req); err != nil {
		return err
	}

	res := make([]*Node, 0)

	nodeman.Lock()
	defer nodeman.Unlock()

	exist := false
	for sidx, v := range nodeman.nodes {
		if v.ID == req.ID {
			nodeman.nodes[sidx] = req
			exist = true
		} else {
			if isCloser(req, v) {
				res = append(res, v)
			}
		}
		//if v.Loc.X < req.Loc.X && v.Loc.Y < req.Loc.Y && v.Loc.Z < req.Loc.Z {
	}

	//TODO: sort res
	for i := 0; i < len(res); i++ {
		for j := i + 1; j < len(res); j++ {
			if isSmaller(res[j], res[i], req) {
				res[i], res[j] = res[j], res[i]
			}
		}
	}
	if !exist {
		nodeman.nodes = append(nodeman.nodes, req)
		log.Infof("Register Sensor:%v", req)
	}
	return c.JSON(http.StatusOK, res)
}

func isSmaller(a, b, sensor *Node) bool {
	if distance(sensor, a) < distance(sensor, b) {
		return true
	}
	return false
}

func isCloser(sensor, relay *Node) bool {
	// 0.
	sink := &Node{Loc: Location{0, 0, 0}}
	if distance(relay, sink) >= distance(sensor, sink) {
		return false
	}
	// 1. replay-sink distance < sensor-sink distance
	/*if distance(relay, sink) < distance(sensor, sink) {
		return true
	}*/
	// 1. sensor-sink distance > sensor-relay distance
	if distance(sensor, sink) <= distance(sensor, relay) {
		return false
	}
	return true
}

func distance(a, b *Node) int {
	xd := a.Loc.X - b.Loc.X
	yd := a.Loc.Y - b.Loc.Y
	zd := a.Loc.Z - b.Loc.Z
	return xd*xd + yd*yd + zd*zd
}
