package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//VehiclesOnRoad how many on road
func VehiclesOnRoad(w http.ResponseWriter, req *http.Request) {
	rd := ExtractRoadID(w, req)
	x := len(rd.Vehicles)
	w.Write([]byte(strconv.Itoa(x)))
}

//ExtractRoadID from request context
func ExtractRoadID(w http.ResponseWriter, req *http.Request) *Road {
	vars := mux.Vars(req)
	RoadID, _ := (vars["rd"])
	rd := silkboard.RoadByID(RoadID)
	if rd == nil {
		panic("invalid roadid")
	}
	return rd
}

// RoadStop the signal
func RoadStop(w http.ResponseWriter, req *http.Request) {
	rd := ExtractRoadID(w, req)
	rd.Go = false
	s, _ := json.Marshal(rd)
	w.Write(s)
}

// RoadGo the signal
func RoadGo(w http.ResponseWriter, req *http.Request) {
	rd := ExtractRoadID(w, req)
	rd.Go = true
	s, _ := json.Marshal(rd)
	w.Write(s)
}

// RoadsHandler web api
func RoadsHandler(w http.ResponseWriter, req *http.Request) {
	s, _ := json.Marshal(silkboard.Roads)
	fmt.Println(string(s))
	w.Write(s)
}

// StatsHandler web api
func StatsHandler(w http.ResponseWriter, req *http.Request) {
	type Stats struct {
		WaitTime int `json:"wait_time"`
	}
	s := &Stats{WaitTime: silkboard.WaitTime()}
	js, _ := json.Marshal(s)
	w.Write(js)
}

// VehicleHandler web api
func VehicleHandler(w http.ResponseWriter, req *http.Request) {
	rd := ExtractRoadID(w, req)
	var found bool
	vars := mux.Vars(req)
	vehiclePlate := vars["veh"]
	for i := 0; i < len(rd.Vehicles); i++ {
		if rd.Vehicles[i].Plate == vehiclePlate {
			found = true
			s, _ := json.Marshal(rd.Vehicles[i].VehicleInfo())
			fmt.Println(string(s))
			w.Write(s)
		}
	}
	if !found {
		w.Write([]byte("not found"))
		s, _ := json.Marshal(silkboard)
		fmt.Println(s)
		w.Write(s)
	}
}

// StopVehicle api
func StopVehicle(w http.ResponseWriter, req *http.Request) {
	rd := ExtractRoadID(w, req)
	vars := mux.Vars(req)
	vehiclePlate, _ := vars["veh"]
	for i := 0; i < len(rd.Vehicles); i++ {
		if rd.Vehicles[i].Plate == vehiclePlate {
			rd.Vehicles[i].Brake = true
			w.Write([]byte("stopped "))
		}
	}
	w.Write([]byte("done"))
}

// GoVehicle API
func GoVehicle(w http.ResponseWriter, req *http.Request) {
	rd := ExtractRoadID(w, req)
	vars := mux.Vars(req)
	vehiclePlate := vars["veh"]
	for i := 0; i < len(rd.Vehicles); i++ {
		if rd.Vehicles[i].Plate == vehiclePlate {
			rd.Vehicles[i].Brake = false
			w.Write([]byte("go "))
		}
	}
	w.Write([]byte("done"))
}
