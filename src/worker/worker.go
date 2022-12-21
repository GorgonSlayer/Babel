package worker

import (
	"net/http"
	"radiola.co.nz/babel/src/model/outtakeRequest"
)

// Worker /** This worker interface is designed to enable a worker implementation for any GIS polling system to work with the priority queue. **/
type Worker interface {
	IntakeRequest(*http.Client) (*http.Response, error)
	ProcessData(response *http.Response) ([]outtakeRequest.TransitClockEvent, error)
	OuttakeRequest(client *http.Client, tce []outtakeRequest.TransitClockEvent) (bool, error)
	RefreshRate() int64
}
