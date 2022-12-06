package intake

import (
	"net/http"
	"strings"
)

/**
	FleetPin Intake API
	This Jobs class provides us with all of our means to make HTTP requests to Fleetpins API.
**/

// IFleetPinIntakeAPI /** This interface is exclusively for HTTP requests to FleetPin's API **/
type IFleetPinIntakeAPI interface {
	GetAssetsHttpRequest(client *http.Client) (*http.Response, error)
	GetAssetByIdHttpRequest(client *http.Client, machineId string) (*http.Response, error)
}

// FleetPinAPIKey /** FleetPin key in its original format. **/
type fleetPinAPIKey struct {
	jwtToken string //Environmental variable API key
}

// NewFleetPinAPIWorker /** Generates our JWT token, needs to be formatted as "JWT /apikey/" **/
func NewFleetPinAPIWorker(key string) IFleetPinIntakeAPI {
	s := []string{"JWT ", key}
	jwt := strings.Join(s, "")
	return fleetPinAPIKey{jwtToken: jwt}
}

// GetAssetsHttpRequest /** Makes an HTTP intakeRequest for FleetPin's assets. **/
func (fpak fleetPinAPIKey) GetAssetsHttpRequest(client *http.Client) (*http.Response, error) {
	req, err := http.NewRequest("GET", "https://app.fleetpin.co.nz/api/assets", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fpak.jwtToken)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, err
}

// GetAssetByIdHttpRequest /** Makes an HTTP intakeRequest for a specific FleetPin asset. **/
func (fpak fleetPinAPIKey) GetAssetByIdHttpRequest(client *http.Client, machineId string) (*http.Response, error) {
	urlParts := []string{"https://app.fleetpin.co.nz/api/assets/", machineId}
	urlString := strings.Join(urlParts, "")
	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fpak.jwtToken)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, err
}
