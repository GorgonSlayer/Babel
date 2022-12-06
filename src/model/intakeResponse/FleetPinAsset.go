package intakeResponse

import "time"

// FleetPinAsset /** Root level FleetPin Asset **/
type FleetPinAsset struct {
	Id                  string                  `json:"id"`
	Name                string                  `json:"name"`
	Description         string                  `json:"description"`
	Icon                string                  `json:"icon"`
	Acc                 bool                    `json:"acc"`
	Attributes          []FleetPinAssetKeyValue `json:"attributes"`
	AttributeKey        string                  `json:"attributeKey"`
	ActiveTrip          string                  `json:"activeTrip"`
	Odometer            float64                 `json:"odometer"`
	Hourmeter           int64                   `json:"hourmeter"`
	Hubometer           float64                 `json:"hubometer"`
	HubometerEnabled    bool                    `json:"hubometerEnabled"`
	HubometerForService bool                    `json:"hubometerForService"`
	OutputEnabled       bool                    `json:"outputEnabled"`
	Alerts              FleetPinAssetAlerts     `json:"alerts"`
	UsageMetric         string                  `json:"usageMetric"`
	Service             FleetPinAssetService    `json:"service"`
	Ruc                 FleetPinAssetRuc        `json:"ruc"`
	Tier                FleetPinAssetTier       `json:"tier"`
	Meters              FleetPinAssetMeter      `json:"meters"`
	Key                 FleetPinAssetKeyValue   `json:"key"`
	Active_alerts       []string                `json:"active_alerts"`
	Active_warnings     []string                `json:"active_warnings"`
	Position            FleetPinAssetPosition   `json:"position"`
}

// FleetPinAssetAlerts /** This seems to be for Fleetpin software specific alerting. **/
type FleetPinAssetAlerts struct {
	PowerOff FleetPinAssetAlertPowerOff `json:"power-off"`
	PowerOn  FleetPinAssetAlertPowerOn  `json:"power-on"`
	Sos      FleetPinAssetAlertSOS      `json:"sos"`
	Service  FleetPinAssetAlertService  `json:"service"`
	Wof      FleetPinAssetAlertWof      `json:"wof"`
	Rego     FleetPinAssetAlertRego     `json:"rego"`
	Ruc      FleetPinAssetAlertRUC      `json:"ruc"`
}

// FleetPinAssetAlertPowerOff /** Fleetpin flag for whether this vehicle is current off and the event it is associated with. **/
type FleetPinAssetAlertPowerOff struct {
	Enabled         bool                                 `json:"enabled"`
	Count           int64                                `json:"count"`
	UnderScoreEvent FleetPinAssetAlertPowerLocationEvent `json:"_event"`
}

// FleetPinAssetAlertPowerOn /** Fleetpin flag for whether this vehicle is currently on and the event it is associated with. **/
type FleetPinAssetAlertPowerOn struct {
	Enabled         bool   `json:"enabled"`
	Count           int64  `json:"count"`
	UnderScoreEvent string `json:"_event"`
}

// FleetPinAssetAlertPowerLocationEvent /** This is used for the power off and power on lat long location events **/
type FleetPinAssetAlertPowerLocationEvent struct {
	UnderScoreId string    `json:"_id"`
	DeviceTime   time.Time `json:"deviceTime"`
	Lat          float64   `json:"lat"`
	Long         float64   `json:"lng"`
	Location     string    `json:"location"`
}

// FleetPinAssetAlertSOS /** Presumably this is a flag for whether this vehicle is in need of assistance or emergency assistance. **/
type FleetPinAssetAlertSOS struct {
	Enabled bool  `json:"enabled"`
	Count   int64 `json:"count"`
}

// FleetPinAssetAlertService /**  Presumably this is a flag for whether this vehicle is in service. **/
type FleetPinAssetAlertService struct {
	Enabled bool `json:"enabled"`
}

// FleetPinAssetAlertWof /** WOF due date tracking for the Alert system. **/
type FleetPinAssetAlertWof struct {
	Enabled bool      `json:"enabled"`
	Due     time.Time `json:"due"`
}

// FleetPinAssetAlertRego /** Registration Alerting for this vehicle. **/
type FleetPinAssetAlertRego struct {
	Enabled bool      `json:"enabled"`
	Due     time.Time `json:"due"`
}

// FleetPinAssetAlertRUC /** RUC alert inside Alert **/
type FleetPinAssetAlertRUC struct {
	Enabled bool `json:"enabled"`
}

type FleetPinAssetService struct {
	Initial float64 `json:"initial"`
	Due     int64   `json:"due"`
}

type FleetPinAssetRuc struct {
	Initial float64 `json:"initial"`
	Due     int64   `json:"due"`
}

// FleetPinAssetTier /** Tracking tier, a FleetPin specific code. Devices can be Tier 3 (real time tracking), Tier 2 (Machine Monitoring) or Tier 1 (Asset Tracking) **/
type FleetPinAssetTier struct {
	Key   string `json:"key"`
	Label string `json:"label"`
}

// FleetPinAssetMeter /** Unknown Meter measurements. **/
type FleetPinAssetMeter struct {
	Zero  FleetPinAssetMeterReading `json:"0"`
	One   FleetPinAssetMeterReading `json:"1"`
	Two   FleetPinAssetMeterReading `json:"2"`
	Three FleetPinAssetMeterReading `json:"3"`
	Four  FleetPinAssetMeterReading `json:"4"`
}

// FleetPinAssetMeterReading /** Some sort of reading of a Meter, it is not clear what this does.  **/
type FleetPinAssetMeterReading struct {
	Count int64 `json:"count"`
}

// FleetPinAssetKeyValue /** Vehicle Key **/
type FleetPinAssetKeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// FleetPinAssetPosition /** Vehicle Position **/
type FleetPinAssetPosition struct {
	UnderScoreLocation FleetPinAssetLocation        `json:"_location"`
	PositionId         string                       `json:"id"`
	Lat                float64                      `json:"lat"`
	Long               float64                      `json:"lng"`
	Heading            int64                        `json:"heading"`
	Speed              int64                        `json:"speed"`
	DeviceTime         time.Time                    `json:"deviceTime"`
	Location           string                       `json:"location"`
	Alert              string                       `json:"alert"`
	IsAccOn            bool                         `json:"isAccOn"`
	Other              FleetPinAssetPositionOmnibus `json:"other"`
}

// FleetPinAssetLocation /** FleetPin Position Location. This seems to be used for geofencing. I.e. Bus location is "Depot". **/
type FleetPinAssetLocation struct {
	UnderScoreId string `json:"_id"`
	Name         string `json:"name"`
}

// FleetPinAssetPositionOmnibus /** GPS related data. HDOP (Horizontal Dilution of Precision) exists **/
type FleetPinAssetPositionOmnibus struct {
	Hdop          float64                     `json:"hdop"`
	State         FleetPinAssetPositionState  `json:"state"`
	BackupVoltage float64                     `json:"backup_voltage"`
	MainVoltage   float64                     `json:"main_voltage"`
	GsmSignal     int64                       `json:"gsm_signal"`
	GpsSatellite  int64                       `json:"gps_satellites"`
	GSensorMax    FleetPinAssetPositionSensor `json:"g_sensor_max"`
}

// FleetPinAssetPositionState /** GPS matrix data which is provided by these FleetPin trackers. **/
type FleetPinAssetPositionState struct {
	Inputs  []int64 `json:"inputs"`
	Outputs []int64 `json:"outputs"`
}

// FleetPinAssetPositionSensor /** G Sensor Max data. Presumably it has some meaning for the physics involved with GPS calculations.**/
type FleetPinAssetPositionSensor struct {
	X int64 `json:"x"`
	Y int64 `json:"y"`
	Z int64 `json:"z"`
}
