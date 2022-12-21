package test

import (
	"io"
	"net/http"
	"radiola.co.nz/babel/src/model/intakeResponse"
	"radiola.co.nz/babel/src/service"
	"radiola.co.nz/babel/src/util/logger"
	"reflect"
	"strings"
	"testing"
	"time"
)

/**
	FleetPin Service Unit Testing. This exists to transform our Intake Response, extract the JSON into a struct;
		We then transform it into a Transit Clock compatible struct for inserting into TransitClock.
**/

/** Testing with a bad status code. **/
func TestFleetPinAssetConstructorInvalidStatusCode(t *testing.T) {
	sampleResponse := &http.Response{
		StatusCode: 500,
		Body:       io.NopCloser(strings.NewReader(`{"id":"5a1dc7d15b4cf60a00b1bbd7","name":"66 MAN 11-190","description":"","icon":"school-bus","acc":true,"attributes":[{"key":"fleet","value":"66"},{"key":"","value":""}],"attributeKey":"fleet","odometer":27866.690390711694,"hourmeter":10544712386,"hubometer":181287.68689825718,"hubometerEnabled":true,"hubometerForService":true,"outputEnabled":false,"alerts":{"power-off":{"enabled":true,"count":14,"_event":{"_id":"638f9afcb4e3c9000d817634","deviceTime":"2022-12-06T19:41:46.724Z","lat":-37.739045999999995,"lng":176.145115}},"power-on":{"enabled":false,"count":207,"_event":"638fea5e70f5ba000c5eee5f"},"sos":{"enabled":true,"count":0},"service":{"enabled":false},"wof":{"enabled":true,"due":"2023-01-21T00:00:00.000Z"},"rego":{"enabled":true,"due":"2023-01-20T00:00:00.000Z"},"ruc":{"enabled":true}},"usageMetric":"kilometers","service":{"initial":0,"due":0},"ruc":{"initial":180403.037467346,"due":183175},"tier":{"key":"tier-3","label":"Tier 3 - Real-time Tracking"},"meters":{"0":{"count":0},"1":{"count":6689},"2":{"count":0},"3":{"count":0},"4":{}},"paused":null,"hidden":null,"key":{"key":"fleet","value":"66"},"active_alerts":["power-off"],"active_warnings":[],"position":{"_location":"5720329e9fbf66016f83fc6e","id":"638febce10fc3e000c70d282","lat":-37.739070999999996,"lng":176.14505499999999,"heading":177,"speed":0,"deviceTime":"2022-12-07T01:26:35.395Z","location":null,"alert":null,"isAccOn":true,"other":{"hdop":0.9,"gps_satellites":10,"g_sensor_max":{"x":22,"y":-79,"z":-957}}}}`)),
	}
	logger := logger.NewLogger(false, "test.log")
	converter := service.NewFleetPinService(logger)
	fpa, err := converter.FleetPinAssetConstructor(sampleResponse)
	if err == nil {
		t.Logf(" \n This should have thrown an error for an invalid status code. \n ")
	}
	if !reflect.ValueOf(fpa).IsZero() {
		t.Errorf("We got a zero value for this fleetpin asset.")
	}
}

/** Testing how we handle Valid FleetPin Objects. **/
func TestFleetPinAssetConstructorValidJSON(t *testing.T) {
	sampleResponse := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"id":"5a1dc7d15b4cf60a00b1bbd7","name":"66 MAN 11-190","description":"","icon":"school-bus","acc":true,"attributes":[{"key":"fleet","value":"66"},{"key":"","value":""}],"attributeKey":"fleet","odometer":27866.690390711694,"hourmeter":10544712386,"hubometer":181287.68689825718,"hubometerEnabled":true,"hubometerForService":true,"outputEnabled":false,"alerts":{"power-off":{"enabled":true,"count":14,"_event":{"_id":"638f9afcb4e3c9000d817634","deviceTime":"2022-12-06T19:41:46.724Z","lat":-37.739045999999995,"lng":176.145115}},"power-on":{"enabled":false,"count":207,"_event":"638fea5e70f5ba000c5eee5f"},"sos":{"enabled":true,"count":0},"service":{"enabled":false},"wof":{"enabled":true,"due":"2023-01-21T00:00:00.000Z"},"rego":{"enabled":true,"due":"2023-01-20T00:00:00.000Z"},"ruc":{"enabled":true}},"usageMetric":"kilometers","service":{"initial":0,"due":0},"ruc":{"initial":180403.037467346,"due":183175},"tier":{"key":"tier-3","label":"Tier 3 - Real-time Tracking"},"meters":{"0":{"count":0},"1":{"count":6689},"2":{"count":0},"3":{"count":0},"4":{}},"paused":null,"hidden":null,"key":{"key":"fleet","value":"66"},"active_alerts":["power-off"],"active_warnings":[],"position":{"_location":"5720329e9fbf66016f83fc6e","id":"638febce10fc3e000c70d282","lat":-37.739070999999996,"lng":176.14505499999999,"heading":177,"speed":0,"deviceTime":"2022-12-07T01:26:35.395Z","location":null,"alert":null,"isAccOn":true,"other":{"hdop":0.9,"gps_satellites":10,"g_sensor_max":{"x":22,"y":-79,"z":-957}}}}`)),
	}
	logger := logger.NewLogger(false, "test.log")
	converter := service.NewFleetPinService(logger)
	fpa, err := converter.FleetPinAssetConstructor(sampleResponse)
	if err != nil { // We should be producing an error.
		t.Logf(" \n Error: \n %v ", err.Error())
		t.Errorf("HTTP request to FleetPin failed.\n Response: \n %+v \n", sampleResponse)
	}
	if fpa[0].Id != "5a1dc7d15b4cf60a00b1bbd7" { //Valid the data exists .
		t.Logf("\n Error: \n ")
		t.Errorf("Something has gone wrong UnMarshalling")
	}
	if fpa[0].Position.Lat != -37.739070999999996 { //Validate positional data.
		t.Logf("\n Error: \n ")
		t.Errorf("Something has gone wrong UnMarshalling Position")
	}
	if fpa[0].Position.Other.Hdop != 0.9 {
		t.Logf("\n Error: \n ")
		t.Errorf("Something has gone wrong UnMarshalling Satelite data under Other Category.")
	}
}

/** Testing how we handle minimal FleetPin Objects, so it lacks the appropriate data. **/
func TestFleetPinAssetConstructorIdOnlyJSON(t *testing.T) {
	sampleResponse := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"id":"asdfsdgsfg"}`)),
	}
	logger := logger.NewLogger(false, "test.log")
	converter := service.NewFleetPinService(logger)
	fpa, err := converter.FleetPinAssetConstructor(sampleResponse)
	if err != nil {
		t.Logf(" \n This should have thrown tons of errors. \n ")
	}
	if fpa[0].Id != "asdfsdgsfg" { //Valid the data exists .
		t.Logf("\n Error: \n ")
		t.Errorf("Something has gone wrong UnMarshalling")
	}
	if fpa[0].Hubometer != 0 {
		t.Logf("\n Error: \n ")
		t.Errorf("We should have no data in this catalog.")
	} else if fpa[0].Hourmeter != 0 {
		t.Logf("\n Error: \n ")
		t.Errorf("We should have no data in this catalog.")
	}
}

/** Testing an invalid/empty FleetPinAsset object. It shouldn't fail to return anything, the objective here is to ensure we return something invalid. **/
func TestFleetPinAssetToTransitClockFormatConverterFail(t *testing.T) {
	f := intakeResponse.FleetPinAsset{
		Name: "blahblahblah",
	}
	var fpa []intakeResponse.FleetPinAsset
	fpa = append(fpa, f)
	logger := logger.NewLogger(false, "test.log")
	converter := service.NewFleetPinService(logger)
	tce := converter.FleetPinAssetToTransitClockFormatConverter(fpa)
	if tce[0].VehicleId != fpa[0].Name {
		t.Errorf("Something is broken in the Format Converter.")
	}
}

/** Testing a Valid FleetPin Object passed to our TransitClock formatter. **/
func TestFleetPinAssetToTransitClockFormatConverterSuccess(t *testing.T) {
	pos := intakeResponse.FleetPinAssetPosition{
		DeviceTime: time.Now(),
		Lat:        0,
		Long:       0,
		Heading:    180,
		Speed:      50,
	}
	f := intakeResponse.FleetPinAsset{
		Id:       "blahblahblah",
		Name:     "Da Bus",
		Position: pos,
	}
	var fpa []intakeResponse.FleetPinAsset
	fpa = append(fpa, f)
	logger := logger.NewLogger(false, "test.log")
	converter := service.NewFleetPinService(logger)
	tce := converter.FleetPinAssetToTransitClockFormatConverter(fpa)
	if tce[0].VehicleId != fpa[0].Name {
		t.Errorf("FleetPinAsset to TransitClock converter VehicleId is broken.")
	}
	if tce[0].Speed != float64(fpa[0].Position.Speed) {
		t.Errorf("FleetPinAsset to TransitClock converter Speed is broken.")
	}
	if tce[0].Heading != fpa[0].Position.Heading {
		t.Errorf("FleetPinAsset to TransitClock converter Heading is broken.")
	}
	if tce[0].Lat != fpa[0].Position.Lat {
		t.Errorf("FleetPinAsset to TransitClock converter Lat is broken.")
	}
	if tce[0].Lon != fpa[0].Position.Long {
		t.Errorf("FleetPinAsset to TransitClock converter Long is broken.")
	}
	if tce[0].Time != fpa[0].Position.DeviceTime.Unix() {
		t.Errorf("FleetPinAsset to TransitClock converter DeviceTime is off.")
	}
}
