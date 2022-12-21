package service

import (
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"net/http"
	intake "radiola.co.nz/babel/src/model/intakeResponse"
	outtake "radiola.co.nz/babel/src/model/outtakeRequest"
	"radiola.co.nz/babel/src/util/logger"
)

// IFleetPinService /** FleetPin Service interface to allow others to access **/
type IFleetPinService interface {
	FleetPinAssetConstructor(res *http.Response) ([]intake.FleetPinAsset, error)
	FleetPinAssetToTransitClockFormatConverter(fpa []intake.FleetPinAsset) outtake.TransitClockEvent
}

// FleetPinService /** This struct should store any state needed to process these FleetPin JSONs into useable structs. **/
type FleetPinService struct {
	logger logger.Logger
}

// NewFleetPinService /** Constructor we use for importing this service. **/
func NewFleetPinService(l logger.Logger) FleetPinService {
	return FleetPinService{logger: l}
}

// FleetPinAssetConstructor /** This function should process a single JSON representing a single asset into an appropriate struct. **/
func (fps FleetPinService) FleetPinAssetConstructor(res *http.Response) ([]intake.FleetPinAsset, error) {
	var assets []intake.FleetPinAsset
	if res.StatusCode == http.StatusOK {
		decoder := json.NewDecoder(res.Body)
		for decoder.More() { //Iterate over the decoding json.
			var asset intake.FleetPinAsset
			err := decoder.Decode(&asset)
			if err != nil {
				fps.logger.Zap.Error("An Error occurred during FleetPinAssetConstructor JSON decoding", zap.Error(err))
			}
			fps.logger.Zap.Info("FleetPinAssetConstructor", zap.Any("FleetPinAsset", &asset)) //The individual asset.
			assets = append(assets, asset)
		}
		return assets, nil
	}
	return assets, errors.New("intake response received has status code other than 200")
}

// FleetPinAssetToTransitClockFormatConverter /** This function converts the FleetPinAsset struct into a transit clock friendly format. **/
func (fps FleetPinService) FleetPinAssetToTransitClockFormatConverter(fpa []intake.FleetPinAsset) []outtake.TransitClockEvent {
	var tceArray []outtake.TransitClockEvent
	for _, v := range fpa { //We loop over however many we have.
		var tce outtake.TransitClockEvent
		tce.VehicleId = v.Name
		tce.Time = v.Position.DeviceTime.Unix() //We take Fleetpin's time as a time.Time object, but Transit Clock takes it as a Unix Epoch Int.
		tce.Lat = v.Position.Lat
		tce.Lon = v.Position.Long
		tce.Heading = v.Position.Heading
		tce.Speed = float64(v.Position.Speed) //Fleetpin gives us an integer for speed, but Transit Clock takes a float.
		tce.Door = 0                          //Fleetpin doesn't provide us with anything like that.
		tce.DriverId = ""                     //Fleetpin seems to have no matching for Driver ID via the API at the moment.
		fps.logger.Zap.Info("FleetPinAssetToTransitClockFormatConverter", zap.Any("TransitClockEvent", &tce))
		tceArray = append(tceArray, tce)
	}
	return tceArray
}
