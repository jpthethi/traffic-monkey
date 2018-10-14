package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"time"
	"traffic-monkey/udp"
)

var hub *Hub
var silkboard *Intersection
var udpConn *net.UDPConn

func main() {
	udpConn = udp.Start()
	defer udpConn.Close()
	silkboard = &Intersection{}
	setup(silkboard)
	var officeHours Traffic
	go officeHours.Generate(silkboard)
	web()
}

func setup(crossroad *Intersection) {
	r1 := &Road{ID: "HSR", Name: "HSR", Color: "yellow"}
	r2 := &Road{ID: "ECity", Name: "ECity", Color: "blue"}
	r3 := &Road{ID: "Madivala", Name: "Madivala", Color: "pink"}
	r4 := &Road{ID: "BTM", Name: "BTM", Color: "orange"}

	crossroad.Roads = append(crossroad.Roads, r1, r2, r3, r4)

	for i := 0; i < len(crossroad.Roads); i++ {
		crossroad.Roads[i].Go = false
		crossroad.Roads[i].Orientation = (360 / len(crossroad.Roads)) * i
	}
}

// PutOnRoad to put vehicle at end of road
func (v *Vehicle) PutOnRoad() {
	v.Distance = -1 * (50 + rand.Intn(50))
	if len(v.Route.Vehicles) > 0 {
		v.VehicleInFront = v.Route.Vehicles[len(v.Route.Vehicles)-1]
		if v.VehicleInFront.Distance < 0 {
			v.Distance = v.VehicleInFront.Distance - (rand.Intn(50) + 10)
		}
	}
	v.Route.Vehicles = append(v.Route.Vehicles, v)

}

// OffRoad to put vehicle out of road
func (v *Vehicle) OffRoad() {
	for i, rcount, rlen := 0, 0, len(v.Route.Vehicles); i < rlen; i++ {
		j := i - rcount
		if v.Route.Vehicles[j].Plate == v.Plate {
			v.Route.Vehicles = append(v.Route.Vehicles[:j], v.Route.Vehicles[j+1:]...)
			rcount++
		}
	}
}

// Generate Traffic for intersection
func (t *Traffic) Generate(crossroad *Intersection) {
	for i := 0; i < len(crossroad.Roads); i++ {
		for j := 0; j < 10; j++ {
			var v Vehicle
			v.Plate = strconv.Itoa(j)
			v.velocity = 10
			v.Stop = false
			v.Route = crossroad.Roads[i]
			v.PutOnRoad()
			go v.Drive()
			rand.Seed(time.Now().UnixNano())
			w := rand.Int31n(100)
			time.Sleep(time.Duration(w) * time.Millisecond)
		}
	}
}

// Drive the vehicle
func (v *Vehicle) Drive() {
	if v.VehicleInFront != nil {
		v.velocity = v.VehicleInFront.velocity
	}

	for !v.Crossed() {
		travel := v.velocity * (1 + rand.Intn(15)/10)
		if v.VehicleInFront != nil && !v.VehicleInFront.Crossed() {
			gap := v.DistanceTo(v.VehicleInFront)
			//slowdown
			if gap <= v.velocity*2 {
				travel = v.velocity / 2
			}

			if gap <= v.velocity {
				travel = 0
			}

		}
		if v.Route.Go == false && v.Distance < -1*v.velocity && v.Distance > -2*v.velocity {
			v.Stop = true
		} else {
			if v.Stop == true {
				fmt.Println("Switching Stop")
				v.Stop = false
			}
		}

		if v.Stop == true || v.Brake == true {
			travel = 0
		}
		if travel == 0 {
			v.WaitTime++
		}
		v.Distance += travel
		if v.lastReportedDistance != v.Distance {
			vi := v.VehicleInfo()
			s, _ := json.Marshal(vi)

			//Send to UDP
			if udpConn != nil {
				udpConn.Write(append(s, '\n'))
			}
			if hub != nil {
				//Send to Websocket
				hub.broadcast <- s
			}
			v.lastReportedDistance = v.Distance
		}

		time.Sleep(1000 * time.Millisecond)
	}

	//if crossed remove from road
	v.OffRoad()
	v.PutOnRoad()
	go v.Drive()
}
