package test

import (
	"io"
	"net/http"
	"radiola.co.nz/babel/src/intake"
	"strings"
	"testing"
)

/** This is "RoundTripper" in the HTTP client. We have it implement RoundTrip to replace the RoundTrip inside http.Client. **/
type RoundTripperFunc func(*http.Request) (*http.Response, error)

/** This lets us implement the round trip function with each test. In turn, we can test the response to different HTTP calls. **/
func (fn RoundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return fn(r)
}

/**
	FleetPin Intake Unit Testing. We test the intake functions for various possible responses.
**/

/** Testing for invalid token request. **/
func TestGetAssetsHttpRequestFailure(t *testing.T) {
	client := &http.Client{
		Transport: RoundTripperFunc(func(r *http.Request) (*http.Response, error) {
			// Assert on request attributes
			// Return a response or error you want
			return &http.Response{
				StatusCode: 401,
				Body:       io.NopCloser(strings.NewReader("UnAuthorised")), //This lets us return a string.
			}, nil
		}),
	}
	fpw := intake.NewFleetPinAPIWorker("blahblahblah") //Using the real HTTP request here. This could fail for unrelated reasons.
	res, err := fpw.GetAssetsHttpRequest(client)
	if err != nil {
		t.Logf(" \n Error: \n")
		t.Logf(err.Error())
		t.Fatalf("HTTP request to FleetPin failed.\n Response: \n %+v \n", res)
	}
	if res.StatusCode != http.StatusUnauthorized {
		t.Logf("This test is supposed to produce an HTTP code of 401")
		t.Fatalf("\n Response: \n %+v", res)
	}
}

/** Testing for a valid token request. **/
func TestGetAssetsHttpRequestSuccess(t *testing.T) {
	client := &http.Client{
		Transport: RoundTripperFunc(func(r *http.Request) (*http.Response, error) {
			// Assert on request attributes
			// Return a response or error you want
			return &http.Response{
				StatusCode: 200,
			}, nil
		}),
	}
	fpw := intake.NewFleetPinAPIWorker("blahblahblah") //Don't put the Key in here. Mock this.
	res, err := fpw.GetAssetsHttpRequest(client)
	if err != nil {
		t.Fatalf("Something went wrong during the HTTP request for successful GetAssets Request \n %v \n %v", err, err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Fatalf("This test should produce a statusCode of 200.")
	}
}

/** Testing for an invalid token request and an invalid machine id. **/
func TestGetAssetByIdHttpRequestUnAuthorisedFailure(t *testing.T) {
	client := &http.Client{
		Transport: RoundTripperFunc(func(r *http.Request) (*http.Response, error) {
			// Assert on request attributes
			// Return a response or error you want
			return &http.Response{
				StatusCode: 401,
				Body:       io.NopCloser(strings.NewReader("UnAuthorised")), //This lets us return a string.
			}, nil
		}),
	}
	fpw := intake.NewFleetPinAPIWorker("blahblahblah") //Using the real HTTP request here. This could fail for unrelated reasons.
	res, err := fpw.GetAssetByIdHttpRequest(client, "")
	if err != nil {
		t.Logf(" \n Error: \n")
		t.Logf(err.Error())
		t.Fatalf("HTTP request to FleetPin failed.\n Response: \n %+v \n", res)
	}
	if res.StatusCode != http.StatusUnauthorized {
		t.Logf("This test is supposed to produce an HTTP code of 401")
		t.Fatalf("\n Response: \n %+v", res)
	}
}

/** This simulates the incorrect Machine Id response **/
func TestGetAssetByIdHttpRequestIncorrectMachineIdFailure(t *testing.T) {
	client := &http.Client{
		Transport: RoundTripperFunc(func(r *http.Request) (*http.Response, error) {
			// Assert on request attributes
			// Return a response or error you want
			return &http.Response{
				StatusCode: 500,
				Body:       io.NopCloser(strings.NewReader(`{"message":"Cast to ObjectId failed for value \"5a1dc7d15b4cf60a00b1bbd7s\" at path \"_id\" for model \"Asset\"","name":"CastError","stringValue":"\"5a1dc7d15b4cf60a00b1bbd7s\"","kind":"ObjectId","value":"5a1dc7d15b4cf60a00b1bbd7s","path":"_id"}`)), //This lets us return a string.
			}, nil
		}),
	}
	fpw := intake.NewFleetPinAPIWorker("blahblahblah") //Using the real HTTP request here. This could fail for unrelated reasons.
	res, err := fpw.GetAssetByIdHttpRequest(client, "123456789")
	if err != nil {
		t.Logf(" \n Error: \n")
		t.Logf(err.Error())
		t.Fatalf("HTTP request to FleetPin failed.\n Response: \n %+v \n", res)
	}
	if res.StatusCode != http.StatusInternalServerError {
		t.Logf("This test is supposed to produce an HTTP code of 500")
		t.Fatalf("\n Response: \n %+v", res)
	}
}

func TestGetAssetByIdHttpRequestSuccess(t *testing.T) {
	client := &http.Client{
		Transport: RoundTripperFunc(func(r *http.Request) (*http.Response, error) {
			// Assert on request attributes
			// Return a response or error you want
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`{"id":"5a1dc7d15b4cf60a00b1bbd7","name":"66 MAN 11-190","description":"","icon":"school-bus","acc":false,"attributes":[{"key":"fleet","value":"66"},{"key":"","value":""}],"attributeKey":"fleet","odometer":27759.367009387293,"hourmeter":10534055336,"hubometer":181180.36351693276,"hubometerEnabled":true,"hubometerForService":true,"outputEnabled":false,"alerts":{"power-off":{"enabled":true,"count":12,"_event":{"_id":"638e4a115d68b1000ec49cce","deviceTime":"2022-12-05T19:44:15.122Z","lat":-37.739276,"lng":176.14509099999998}},"power-on":{"enabled":false,"count":205,"_event":"638e9a630e2569000cf319f1"},"sos":{"enabled":true,"count":0},"service":{"enabled":false},"wof":{"enabled":true,"due":"2023-01-21T00:00:00.000Z"},"rego":{"enabled":true,"due":"2023-01-20T00:00:00.000Z"},"ruc":{"enabled":true}},"usageMetric":"kilometers","service":{"initial":0,"due":0},"ruc":{"initial":180403.037467346,"due":183175},"tier":{"key":"tier-3","label":"Tier 3 - Real-time Tracking"},"meters":{"0":{"count":0},"1":{"count":6685},"2":{"count":0},"3":{"count":0},"4":{}},"paused":null,"hidden":null,"key":{"key":"fleet","value":"66"},"active_alerts":["power-off"],"active_warnings":[],"position":{"_location":"5720377e124f5b252f35e0af","id":"638e9fd0031d65000d7fbc8f","lat":-37.706156,"lng":176.110965,"heading":167,"speed":0,"deviceTime":"2022-12-06T01:50:06.533Z","location":null,"alert":"acc-off","isAccOn":false,"other":{"hdop":0.9,"gps_satellites":10,"g_sensor_max":{"x":-30,"y":-51,"z":-1025}}}}`)), //This lets us return a string.
			}, nil
		}),
	}
	fpw := intake.NewFleetPinAPIWorker("blahblahblah") //Using the real HTTP request here. This could fail for unrelated reasons.
	res, err := fpw.GetAssetByIdHttpRequest(client, "123456789")
	if err != nil {
		t.Logf(" \n Error: \n")
		t.Logf(err.Error())
		t.Fatalf("HTTP request to FleetPin failed.\n Response: \n %+v \n", res)
	}
	if res.StatusCode != http.StatusOK {
		t.Logf("This test is supposed to produce an HTTP code of 500")
		t.Fatalf("\n Response: \n %+v", res)
	}
}
