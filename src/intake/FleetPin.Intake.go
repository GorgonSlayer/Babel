package intake

import (
	"net/http"
	"strings"

	"go.uber.org/zap"
	"radiola.co.nz/babel/src/util/logger"
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
type FleetPinAPIKey struct {
	JwtToken  string //Environmental variable API key
	UrlString string //"https://app.fleetpin.co.nz/api/assets"
	logger    logger.Logger
}

// NewFleetPinAPIWorker /** Generates our JWT token, needs to be formatted as "JWT /apikey/" **/
func NewFleetPinAPIWorker(key string, url string, l logger.Logger) IFleetPinIntakeAPI {
	s := []string{"JWT ", key}
	jwt := strings.Join(s, "")
	return FleetPinAPIKey{JwtToken: jwt, UrlString: url, logger: l}
}

// GetAssetsHttpRequest /** Makes an HTTP intakeRequest for FleetPin's assets. **/
func (fpak FleetPinAPIKey) GetAssetsHttpRequest(client *http.Client) (*http.Response, error) {
	req, err := http.NewRequest("GET", fpak.UrlString, nil)
	if err != nil {
		fpak.logger.Zap.Error("GetAssetHttpRequest New Request failed.", zap.Error(err))
		return nil, err
	}
	req.Header.Add("Authorization", fpak.JwtToken)
	res, err := client.Do(req)
	if err != nil {
		fpak.logger.Zap.Error("GetAssetHttpRequest Do Request failed.", zap.Error(err))
		return nil, err
	}
	fpak.logger.Zap.Debug("GetAssetsHttpRequests Response", zap.Any("Response Header", res.Header), zap.Any("Response Body", res.Body))
	return res, err
}

// GetAssetByIdHttpRequest /** Makes an HTTP intakeRequest for a specific FleetPin asset. **/
func (fpak FleetPinAPIKey) GetAssetByIdHttpRequest(client *http.Client, machineId string) (*http.Response, error) {
	urlParts := []string{fpak.UrlString, "/", machineId}
	urlString := strings.Join(urlParts, "")
	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		fpak.logger.Zap.Error("GetAssetHttpRequest New Request failed. Error: %v", zap.Any("Error", err.Error()))
		return nil, err
	}
	req.Header.Add("Authorization", fpak.JwtToken)
	res, err := client.Do(req)
	if err != nil {
		fpak.logger.Zap.Error("GetAssetHttpRequest Do Request failed.", zap.Any("Error", err.Error()))
		return nil, err
	}
	fpak.logger.Zap.Debug("GetAssetsHttpRequests Response", zap.Any("Response Header", res.Header), zap.Any("Response Body", res.Body))
	return res, err
}
