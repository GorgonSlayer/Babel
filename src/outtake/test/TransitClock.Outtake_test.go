package test

import (
	"encoding/json"
	"io"
	"net/http"
	"radiola.co.nz/babel/src/model/outtakeRequest"
	"radiola.co.nz/babel/src/outtake"
	"radiola.co.nz/babel/src/util/logger"
	"strconv"
	"strings"
	"testing"
	"time"
)

var host = "localhost"
var port = 3000
var key = "magicKey"
var agency = "agencyName"

/** This is "RoundTripper" in the HTTP client. We have it implement RoundTrip to replace the RoundTrip inside http.Client. **/
type RoundTripperFunc func(*http.Request) (*http.Response, error)

/** This lets us implement the round trip function with each test. In turn, we can test the response to different HTTP calls. **/
func (fn RoundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return fn(r)
}

/** Successfully testing Generate TransitClock request. **/
func TestGenerateTransitClockRequestSuccess(t *testing.T) {
	logger := logger.NewLogger(false, "test.log")
	tco := outtake.NewTransitClockOuttake(host, port, key, agency, logger)
	req, err := tco.GenerateTransitClockRequest()
	if err != nil {
		t.Errorf("Something has gone wrong in GenerateTransitClockRequest() \n %v \n", err)
	}
	s := []string{"https://", host, "/api/v1/key/", key, "/agency/", agency, "/command/pushAvl"}
	compositeUrl := strings.Join(s, "")
	if req.URL.String() != compositeUrl {
		t.Logf("Host: %s", req.URL.String())
		t.Logf("compositeURL: %s", compositeUrl)
		t.Errorf("The Raw Query did not match the compositeUrl format used for TransitClock")
	}
}

/** Generate an invalid TransitClock request. **/
func TestGenerateURLParams(t *testing.T) {
	logger := logger.NewLogger(false, "test.log")
	tco := outtake.NewTransitClockOuttake(host, port, key, agency, logger)
	req, err := tco.GenerateTransitClockRequest()
	if err != nil {
		t.Errorf("Something has gone wrong in GenerateTransitClockRequest() \n %v \n", err)
	}
	tce := outtakeRequest.TransitClockEvent{
		VehicleId: "123456789",
		Time:      time.Now().Unix(),
		Lat:       -37.739345,
		Lon:       176.14524799999998,
		Speed:     50,
		Heading:   162,
		Door:      0,
		DriverId:  "987654321",
	}
	tco.GenerateURLParams(req, tce)
	s := []string{"door=0&driverId=987654321&h=162&lat=-37.739345&lon=176.145248&s=50.000000&t=", strconv.FormatInt(time.Now().Unix(), 10), "&v=123456789"}
	fullURL := strings.Join(s, "")
	if req.URL.RawQuery != fullURL {
		t.Logf("RawQuery: %s", req.URL.RawQuery)
		t.Logf("FullQuery: %s", fullURL)
		t.Errorf("Formatting is incorrect for GenerateURLParams")
	}
}

/** Flush Data to TransitClock (return a false response) **/
func TestFlushDataToTransitClock(t *testing.T) {
	client := &http.Client{
		Transport: RoundTripperFunc(func(r *http.Request) (*http.Response, error) {
			// Assert on request attributes
			// Return a response or error you want
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`{"success":true,"message":"AVL processed"}`)),
			}, nil
		}),
	}
	logger := logger.NewLogger(false, "test.log")
	tco := outtake.NewTransitClockOuttake(host, port, key, agency, logger)
	req, err := tco.GenerateTransitClockRequest()
	if err != nil {
		t.Errorf("Something has gone wrong in GenerateTransitClockRequest() \n %v \n", err)
	}
	tce := outtakeRequest.TransitClockEvent{
		VehicleId: "123456789",
		Time:      time.Now().Unix(),
		Lat:       -37.739345,
		Lon:       176.14524799999998,
		Speed:     50,
		Heading:   162,
		Door:      0,
		DriverId:  "987654321",
	}
	tco.GenerateURLParams(req, tce)
	res, err := tco.FlushDataToTransitClock(client, req)
	if err != nil {
		t.Errorf("An error occurred during the Flush Data to TransitClock call.")
	}
	var tcr outtakeRequest.TransitClockResponse
	err = json.NewDecoder(res.Body).Decode(&tcr)
	if err != nil {
		t.Errorf("Error during decoding of JSON. Error: %v", err.Error())
	}
	t.Logf("Trasnit Clock response: %+v", tcr)
	if tcr.Message != "AVL processed" {
		t.Errorf("Message in TransitClock response is not what we expected. Message: \n %s", tcr.Message)
	} else if tcr.Success != true {
		t.Errorf("TransitClock Response failed for some reason. \n Success: %t", tcr.Success)
	}
}
