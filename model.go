package main

// Traffic where it comes from
type Traffic struct {
	Name string
}

// Intersection is the crossing
type Intersection struct {
	Roads []*Road `json:"roads"`
}

// Road of an intersection from where vehicles pass
type Road struct {
	ID          string     `json:"id"`
	Name        string     `json:"Name"`
	Color       string     `json:"color"`
	Orientation int        `json:"orientation"`
	Vehicles    []*Vehicle `json:"-"`
	Go          bool       `json:"signal"`
}

// Vehicle on a road
type Vehicle struct {
	Plate                string
	Route                *Road
	intersection         *Intersection
	Distance             int
	velocity             int
	Stop                 bool
	Brake                bool
	WaitTime             int
	VehicleInFront       *Vehicle
	lastReportedDistance int
}

// VehicleInfo struct
type VehicleInfo struct {
	Plate    string `json:"plate"`
	Road     string `json:"road"`
	Distance int    `json:"distance"`
	Velocity int    `json:"-"`
	InFront  string `json:"-"`
	Stopped  bool   `json:"-"`
}

// RoadByID gives road by ID
func (intersection *Intersection) RoadByID(RoadID string) *Road {
	for i := 0; i < len(intersection.Roads); i++ {
		if intersection.Roads[i].ID == RoadID {
			return intersection.Roads[i]
		}
	}
	return nil
}

// ResetLastReported resets reporting
func (intersection *Intersection) ResetLastReported() {
	for i := 0; i < len(intersection.Roads); i++ {
		for j := 0; j < len(intersection.Roads[i].Vehicles); j++ {
			intersection.Roads[i].Vehicles[j].lastReportedDistance = 0
		}
	}
}

// TrafficLight stopping or allowing vehicles to pass
type TrafficLight struct {
	from *Road
}

// VehicleInfo compute from vehicle
func (v *Vehicle) VehicleInfo() VehicleInfo {
	veh := VehicleInfo{Plate: v.Plate, Distance: v.Distance, Road: v.Route.Name, Velocity: v.velocity, InFront: ""}
	veh.Stopped = v.Stop
	if v.VehicleInFront != nil {
		veh.InFront = v.VehicleInFront.Plate
	}

	return veh
}

// Crossed the road ?
func (v *Vehicle) Crossed() bool {
	return v.Distance > 50
}

// Velocity changing method
func (v *Vehicle) Velocity(newVelocity int) {
	v.velocity = newVelocity
}

// DistanceTo next vehicle
func (v *Vehicle) DistanceTo(v2 *Vehicle) int {
	if v2 == nil {
		return 1000
	}
	diff := v2.Distance - v.Distance
	if diff < 0 {
		diff = diff * -1
	}
	return diff
}

// WaitTime total
func (intersection *Intersection) WaitTime() int {
	waitTime := 0
	for i := 0; i < len(intersection.Roads); i++ {
		waitTime += intersection.Roads[i].WaitTime()
	}
	return waitTime
}

// WaitTime total
func (road *Road) WaitTime() int {
	waitTime := 0
	for i := 0; i < len(road.Vehicles); i++ {
		waitTime += road.Vehicles[i].WaitTime
	}
	return waitTime
}
