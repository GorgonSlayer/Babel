package worker

import (
	"net/http"
	"radiola.co.nz/babel/src/intake"
	"radiola.co.nz/babel/src/model/outtakeRequest"
	"radiola.co.nz/babel/src/outtake"
	"radiola.co.nz/babel/src/service"
	"radiola.co.nz/babel/src/util/logger"
)

type FleetPinWorker struct {
	fpk         intake.FleetPinAPIKey
	fps         service.FleetPinService
	tco         outtake.TransitClockOuttake
	refreshRate int64
}

// NewFleetPinWorker /** Constructor for a FleetPin Worker. **/
func NewFleetPinWorker(jwt string, fleetPinUrl string, host string, port int, key string, agency string, logger logger.Logger) Worker {
	return FleetPinWorker{
		fpk:         intake.NewFleetPinAPIWorker(jwt, fleetPinUrl, logger),
		fps:         service.NewFleetPinService(logger),
		tco:         outtake.NewTransitClockOuttake(host, port, key, agency, logger),
		refreshRate: 30, //Seconds. Should be 30 seconds per interval.
	}
}

// IntakeRequest /** Intake Function. It should only take in an HTTP client **/
func (fpw FleetPinWorker) IntakeRequest(client *http.Client) (*http.Response, error) {
	res, err := fpw.fpk.GetAssetsHttpRequest(client)
	if err != nil { //If any errors emerge in the GetAssets method.
		return nil, err
	}
	return res, err
}

// ProcessData /** Processing Function.**/
func (fpw FleetPinWorker) ProcessData(response *http.Response) ([]outtakeRequest.TransitClockEvent, error) {
	var tce []outtakeRequest.TransitClockEvent
	fpa, err := fpw.fps.FleetPinAssetConstructor(response)
	if err != nil {
		return tce, err
	}
	tce = fpw.fps.FleetPinAssetToTransitClockFormatConverter(fpa)
	return tce, err
}

// OuttakeRequest /** Pushes to TransitClock with the fresh data. **/
func (fpw FleetPinWorker) OuttakeRequest(client *http.Client, tce []outtakeRequest.TransitClockEvent) (bool, error) {
	for i, _ := range tce { //Loop over our array of TransitClockEvents to send the data out.
		req, err := fpw.tco.GenerateTransitClockRequest()
		if err != nil {
			return false, err
		}
		fpw.tco.GenerateURLParams(req, tce[i])
		res, err := fpw.tco.FlushDataToTransitClock(client, req)
		if err != nil {
			return false, err
		}
		if res.StatusCode != http.StatusOK { //Status Code is anything other than success.
			return false, err
		}
	}
	return true, nil
}

// RefreshRate /** Returns our refresh rate. **/
func (fpw FleetPinWorker) RefreshRate() int64 {
	return fpw.refreshRate
}
