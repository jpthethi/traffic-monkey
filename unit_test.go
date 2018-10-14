package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"traffic-monkey/udp"

	"github.com/gorilla/mux"
)

var m *mux.Router
var req *http.Request
var err error
var respRec *httptest.ResponseRecorder
var offline bool

func TestMain(mn *testing.M) {
	udpConn = udp.Start()
	defer udpConn.Close()

	//mux router with added question routes
	m = mux.NewRouter()

	//The response recorder used to record HTTP responses
	respRec = httptest.NewRecorder()
	mn.Run()
}

func TestGetWs(t *testing.T) {
	req, err = http.NewRequest("GET", "/ws", nil)
	if err != nil {
		t.Fatal("Creating 'GET /ws' request failed!")
	}

}

func TestGetStats(t *testing.T) {
	req, err = http.NewRequest("GET", "/stats", nil)
	if err != nil {
		t.Fatal("Creating 'GET /stats' request failed!")
	}
}

func TestGetRoads(t *testing.T) {
	req, err = http.NewRequest("GET", "/roads", nil)
	if err != nil {
		t.Fatal("Creating 'GET /roads' request failed!")
	}
}

func TestCreateIntersection(t *testing.T) {
	silkboard = &Intersection{}
	setup(silkboard)
}

func TestGenerateTraffic(t *testing.T) {
	var officeHours Traffic
	officeHours.Generate(silkboard)
}

func TestUUID(t *testing.T) {
	fmt.Println(uuid())
}
