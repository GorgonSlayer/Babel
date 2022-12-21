package mock

import (
	"io"
	"net/http"
	"radiola.co.nz/babel/src/model/outtakeRequest"
	"strings"
	"time"
)

type MockWorker struct {
	refreshRate int64
}

// NewMockWorker /** Our Mock worker to test Queue **/
func NewMockWorker() MockWorker {
	return MockWorker{
		refreshRate: 30,
	}
}

func (m MockWorker) IntakeRequest(*http.Client) (*http.Response, error) {
	return &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader("Blahblahblah")), //This lets us return a string.
	}, nil
}

func (m MockWorker) ProcessData(response *http.Response) ([]outtakeRequest.TransitClockEvent, error) {
	var array []outtakeRequest.TransitClockEvent
	i := outtakeRequest.TransitClockEvent{
		VehicleId: "Da Bus",
		Speed:     50,
		Heading:   180,
		DriverId:  "Otto",
		Lat:       -178,
		Lon:       -36,
		Time:      time.Now().Unix(),
	}
	array = append(array, i)
	return array, nil
}

func (m MockWorker) OuttakeRequest(client *http.Client, tce []outtakeRequest.TransitClockEvent) (bool, error) {
	return true, nil
}

func (m MockWorker) RefreshRate() int64 {
	return m.refreshRate
}
