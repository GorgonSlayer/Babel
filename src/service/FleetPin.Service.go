package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	intake "radiola.co.nz/babel/src/model/intakeResponse"
	outtake "radiola.co.nz/babel/src/model/outtakeRequest"
)

// IFleetPinService /** FleetPin Service interface to allow others to access **/
type IFleetPinService interface {
	FleetPinAssetConstructor(res *http.Response) (intake.FleetPinAsset, error)
	FleetPinAssetToTransitClockFormatConverter(fpa intake.FleetPinAsset) outtake.TransitClockEvent
}

// FleetPinService /** This struct should store any state needed to process these FleetPin JSONs into useable structs. **/
type FleetPinService struct {
}

// NewFleetPinService /** Constructor we use for importing this service. **/
func NewFleetPinService() IFleetPinService {
	return FleetPinService{}
}

// FleetPinAssetConstructor /** This function should process a single JSON representing a single asset into an appropriate struct. **/
func (fps FleetPinService) FleetPinAssetConstructor(res *http.Response) (intake.FleetPinAsset, error) {
	asset := intake.FleetPinAsset{}
	if res.StatusCode == http.StatusOK {
		decoder := json.NewDecoder(res.Body)
		err := decoder.Decode(&asset)
		if err != nil {
			fmt.Println("Error during Decode")
			fmt.Println(err.Error())
		}
		fmt.Printf("\n%+#v\n", asset)
		return asset, nil
	}
	return asset, errors.New("intakeResponse struct received has status code other than 200")
}

// FleetPinAssetToTransitClockFormatConverter /** This function converts the FleetPinAsset struct into a transit clock friendly format. **/
func (fps FleetPinService) FleetPinAssetToTransitClockFormatConverter(fpa intake.FleetPinAsset) outtake.TransitClockEvent {
	var tce outtake.TransitClockEvent
	tce.VehicleId = fpa.Name
	tce.Time = fpa.Position.DeviceTime.Unix() //We take Fleetpin's time as a time.Time object, but Transit Clock takes it as a Unix Epoch Int.
	tce.Lat = fpa.Position.Lat
	tce.Lon = fpa.Position.Long
	tce.Heading = fpa.Position.Heading
	tce.Speed = float64(fpa.Position.Speed) //Fleetpin gives us an integer for speed, but Transit Clock takes a float.
	tce.Door = 0                            //Fleetpin doesn't provide us with anything like that.
	tce.DriverId = ""                       //Fleetpin seems to have no matching for Driver ID via the API at the moment.
	return tce
}
