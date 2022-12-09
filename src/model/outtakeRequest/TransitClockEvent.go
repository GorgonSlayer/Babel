package outtakeRequest

// TransitClockEvent /** This is our outtake object we use to push data into Transit Clock. **/
type TransitClockEvent struct {
	VehicleId string  `json:"v"`        //This is usually an int but can be a string.
	Time      int64   `json:"t"`        //This is an integer
	Lat       float64 `json:"lat"`      //Latitude is always a float
	Lon       float64 `json:"lon"`      //Longitude is always a float
	Speed     float64 `json:"s"`        //Some providers give this as an int and others as a float.
	Heading   int64   `json:"h"`        //This is a 0 to 360 measure
	Door      int64   `json:"door"`     //Assumption is the door is either closed or open, we do not factor in whether it is partially open or fully closed etc.
	DriverId  string  `json:"driverId"` //This is usually an int but can be a string
}

type TransitClockResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
