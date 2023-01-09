package outtake

import (
	"fmt"
	"net/http"
	"strings"

	"go.uber.org/zap"
	"radiola.co.nz/babel/src/model/outtakeRequest"
	"radiola.co.nz/babel/src/util/logger"
)

/**
	FleetPin Outtake API

	This provides us with our functions for inserting data into TransitClock.
**/

// ITransitClockOuttake /** This interface is exclusively to insert data into TransitClock. **/
type ITransitClockOuttake interface {
	GenerateTransitClockRequest() (*http.Request, error)
	GenerateURLParams(req *http.Request, tce outtakeRequest.TransitClockEvent)
	FlushDataToTransitClock(client *http.Client, req *http.Request) (*http.Response, error)
}

// TransitClockOuttake /** This Struct stores all the variables we need to access transitclock and to push data into it. **/
type TransitClockOuttake struct {
	TransitClockHost   string
	TransitClockPort   int
	TransitClockKey    string
	TransitClockAgency string
	logger             logger.Logger
}

// NewTransitClockOuttake /** Constructor for this Outtake. **/
func NewTransitClockOuttake(host string, port int, key string, agency string, l logger.Logger) TransitClockOuttake {
	return TransitClockOuttake{
		TransitClockHost:   host,
		TransitClockPort:   port,
		TransitClockKey:    key,
		TransitClockAgency: agency,
		logger:             l,
	}
}

// GenerateTransitClockRequest /** Simple HTTP request generation function. We take our outtake struct and generate the relevant query. **/
func (tco TransitClockOuttake) GenerateTransitClockRequest() (*http.Request, error) {
	s := []string{"https://", tco.TransitClockHost, "/api/v1/key/", tco.TransitClockKey, "/agency/", tco.TransitClockAgency, "/command/pushAvl"}
	compositeUrl := strings.Join(s, "")
	req, err := http.NewRequest("GET", compositeUrl, nil)
	if err != nil {
		tco.logger.Zap.Error("GenerateTransitClockRequest ", zap.Any("Error", err.Error()))
		return nil, err
	}
	tco.logger.Zap.Info("Request Dump", zap.Any("URL Host", req.URL.Host), zap.Any("URL Path", req.URL.Path))
	return req, err
}

// GenerateURLParams /** This functon takes a TransitClockEvent struct and adds to a TransitClock HTTP Request with relevant param to pass through to TransitClock. **/
func (tco TransitClockOuttake) GenerateURLParams(req *http.Request, tce outtakeRequest.TransitClockEvent) {
	q := req.URL.Query()
	q.Add("v", tce.VehicleId)
	q.Add("t", fmt.Sprintf("%d", tce.Time))
	q.Add("lat", fmt.Sprintf("%f", tce.Lat))
	q.Add("lon", fmt.Sprintf("%f", tce.Lon))
	q.Add("s", fmt.Sprintf("%f", tce.Speed))
	q.Add("h", fmt.Sprintf("%d", tce.Heading))
	q.Add("door", fmt.Sprintf("%d", tce.Door))
	q.Add("driverId", tce.DriverId)
	req.URL.RawQuery = q.Encode() //We need to encode the values back onto the request.
	tco.logger.Zap.Info("GenerateURLParams Raw Query", zap.Any("Query", req.URL.RawQuery))
}

// FlushDataToTransitClock /** Pushes our data using a client Transit Clock instance. **/
func (tco TransitClockOuttake) FlushDataToTransitClock(client *http.Client, req *http.Request) (*http.Response, error) {
	res, err := client.Do(req)
	if err != nil {
		tco.logger.Zap.Error("FlushDataToTransitClock", zap.Any("Error", err.Error()))
		return nil, err
	}
	tco.logger.Zap.Info("Response", zap.Any("Response Status Code", res.StatusCode), zap.Any("Response Header", res.Header), zap.Any("Response Body", res.Body))
	return res, err
}
