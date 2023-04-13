package intake

import (
	"net/http"
	"net/http/httputil"
	"strings"

	"go.uber.org/zap"
	"radiola.co.nz/babel/src/util/logger"
)

/**
	E-Road Job
	This file should contain all the HTTP intakeRequest and intakeResponse logic for interacting with the E-Road API.
**/
/**
	E-Road Jobs Interface is our interface for importing e-road.
	The objective is to provide a single interface for importing to intakeRequest E-road Data.
**/
type IERoadIntakeAPI interface {
	RetrieveIntrospect(client *http.Client) (*http.Response, error)
	RetrieveMachineState(client *http.Client, orgId string, machineId string) (*http.Response, error)
}

/*
*

	E-Road Keys for this particular account are the only state we want to keep here.

*
*/
type ERoadKey struct {
	jwtToken string
	urlHost  string
	logger   logger.Logger
}

/** Constructor for the Intake API object**/
func NewERoadAPI(j string, uh string, l logger.Logger) IERoadIntakeAPI {
	return ERoadKey{
		jwtToken: j,
		urlHost:  uh,
		logger:   l,
	}
}

/** This function requests the full inventory of machines and drivers in fleet road. **/
func (erk ERoadKey) RetrieveIntrospect(client *http.Client) (*http.Response, error) {
	s := []string{erk.urlHost, "/introspect"}
	urlString := strings.Join(s, "")
	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		erk.logger.Zap.Error("RetrieveIntrospect New Request failed.", zap.Error(err))
		return nil, err
	}
	req.Header.Add("eroad-api-key", erk.jwtToken)
	dump, _ := httputil.DumpRequest(req, false)
	erk.logger.Zap.Debug("RetrieveIntrospect Request", zap.Any("Request Header", string(dump)))
	res, err := client.Do(req)
	if err != nil {
		erk.logger.Zap.Error("RetrieveIntrospect Do Request failed.", zap.Error(err))
		return nil, err
	}
	erk.logger.Zap.Debug("RetrieveIntrospect Response", zap.Any("Response Header", res.Header), zap.Any("Response Body", res.Body))
	return res, err
}

/** This function performs the HTTP request for machine state of each individual vehicle in E-Road. **/
func (erk ERoadKey) RetrieveMachineState(client *http.Client, orgId string, machineId string) (*http.Response, error) {
	s := []string{erk.urlHost, "/machine/org/", orgId, "/machine/", machineId, "/state"}
	urlString := strings.Join(s, "")
	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		erk.logger.Zap.Error("RetrieveIntrospect New Request failed.", zap.Error(err))
		return nil, err
	}
	req.Header.Add("eroad-api-key", erk.jwtToken)
	res, err := client.Do(req)
	if err != nil {
		erk.logger.Zap.Error("RetrieveIntrospect Do Request failed.", zap.Error(err))
		return nil, err
	}
	erk.logger.Zap.Debug("RetrieveIntrospect Response", zap.Any("Response Header", res.Header), zap.Any("Response Body", res.Body))
	return res, err
}
