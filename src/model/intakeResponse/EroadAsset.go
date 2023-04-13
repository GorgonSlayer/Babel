package intakeResponse

import "time"

/** ERoad Asset Introspection object. **/
type ERoadIntrospectAsset struct {
	Entities       []ERoadIntrospectEntity     `json:"entities"`
	OrganisationId string                      `json:"organisationId"`
	Permissions    []ERoadIntrospectPermission `json:"permissions"`
}

/** ERoad Introspect Entity **/
type ERoadIntrospectEntity struct {
	EntityId   string `json:"entityId"`
	EntityType string `json:"entityType"`
}

/** ERoad Introspect Driver Entity **/
type ERoadIntrospectDriver struct {
	Alias     string `json:"alias"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Name      string `json:"name"`
}

/** ERoad Introspect Machine Entity **/
type ERoadIntrospectMachine struct {
	DisplayName       string `json:"displayName"`
	MachineName       string `json:"name"`
	RegistrationPlate string `json:"registrationPlate"`
	Vin               string `json:"vin"`
}

/** Handles the case of an Introspect Entity being a fleet. **/
type ERoadIntrospectFleet struct {
	Name string `json:"name"`
}

/** ERoad Introspect Permissions **/
type ERoadIntrospectPermission struct {
	Actions     []string `json:"actions"`
	EntityId    string   `json:"entityId"`
	EntityType  string   `json:"entityType"`
	ProductId   string   `json:"productId"`
	ProductName string   `json:"productName"`
}

/*
*

	ERoad Machine State.

*
*/
type ERoadMachineState struct {
	Status              string                 `json:"status"`
	PrivateModeActive   bool                   `json:"privateModeActive"`
	MachineId           string                 `json:"machineId"`
	OrganisationId      string                 `json:"organisationId"`
	EventType           string                 `json:"eventType"`
	EventTime           time.Time              `json:"eventTime"`
	EventSequenceNumber int                    `json:"eventSequenceNumber"`
	EngineHours         float32                `json:"engineHours"`
	EventSource         string                 `json:"eventSource"`
	Driver              ERoadLastMachineDriver `json:"driver"`
	Device              ERoadMachineDevice     `json:"device"`
	Location            ERoadMachineLocation   `json:"location"`
}

/** Last Driver of this Machine. **/
type ERoadLastMachineDriver struct {
	LastDriverId string    `json:""`
	LoggedIn     bool      `json:""`
	LoginTime    time.Time `json:""`
}

/** Cellphone related data from E-Road **/
type ERoadMachineDevice struct {
	Id                 string   `json:""`
	CellSignalStrength int      `json:""`
	Warnings           []string `json:""`
	Odometer           float32  `json:""`
}

/** **/
type ERoadMachineLocation struct {
	GpsAccuracy      int     `json:""`
	Bearing          int     `json:""`
	Latitude         float32 `json:""`
	Longitude        float32 `json:""`
	ReadableLocation string  `json:""`
	Speed            int     `json:""`
	OffRoad          bool    `json:""`
	NearCity         bool    `json:""`
}
