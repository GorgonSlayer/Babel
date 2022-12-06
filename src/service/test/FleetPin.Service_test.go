package test

import "testing"

/**
	FleetPin Service Unit Testing. This exists to transform our Intake Response, extract the JSON into a struct;
		We then transform it into a Transit Clock compatible struct for inserting into TransitClock.
**/

/** Testing how we handle invalid FleetPin Objects. **/
func TestFleetPinAssetConstructorInvalidJSON(t *testing.T) {

}

/** Testing how we handle minimal FleetPin Objects, containing only what is necessary to insert into TransitClock. **/
func TestFleetPinAssetConstructorMinimalJSON(t *testing.T) {

}

/** Testing how we handle complete FleetPin Objects, containing all possible data from FleetPin. **/
func TestFleetPinAssetConstructorMaximalJSON(t *testing.T) {

}

/** Testing Invalid FleetPin Objects passed to our TransitClock formatter. **/
func TestFleetPinAssetToTransitClockFormatConverterInvalid(t *testing.T) {

}

/** Testing a Valid FleetPin Object passed to our TransitClock formatter. **/
func TestFleetPinAssetToTransitClockFormatConverterSuccess(t *testing.T) {

}
