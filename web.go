package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func web() {
	n := negroni.Classic()
	root := mux.NewRouter()
	hub = newHub()
	go hub.run()

	root.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	root.HandleFunc("/stats", StatsHandler).Methods("GET")
	root.HandleFunc("/roads", RoadsHandler).Methods("GET")
	root.HandleFunc("/roads/{rd}/stop", RoadStop).Methods("GET")
	root.HandleFunc("/roads/{rd}/go", RoadGo).Methods("GET")
	root.HandleFunc("/roads/{rd}/vehicle/{veh}/stop", StopVehicle).Methods("GET")
	root.HandleFunc("/roads/{rd}/vehicle/{veh}/go", GoVehicle).Methods("GET")
	root.HandleFunc("/roads/{rd}/vehicle/{veh}", VehicleHandler).Methods("GET")
	root.HandleFunc("/roads/{rd}/vehicles", VehiclesOnRoad).Methods("GET")

	n.UseHandler(root)
	n.Run(":" + fmt.Sprint(3060))

}
